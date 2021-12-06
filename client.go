package track17

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Client paasport client define
type Client struct {
	conf   *Config
	client *http.Client
}

// NewClient create a new paasport client
func NewClient(secret string, configures ...configurer) (*Client, error) {
	conf := &Config{
		endpoint: ENDPOINT,
		secret:   secret,
	}
	for _, configure := range configures {
		configure(conf)
	}
	conf.withDefault()
	client := Client{
		conf: conf,
	}
	return client.init()
}

func (c *Client) init() (*Client, error) {
	transport, err := c.newTransport()
	if err != nil {
		return c, err
	}
	c.client = &http.Client{
		Transport: transport,
	}
	return c, nil
}

func (c *Client) newTransport() (*http.Transport, error) {
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			conn, err := net.DialTimeout(network, addr, c.conf.HTTPTimeout.ConnectTimeout)
			if err != nil {
				return nil, err
			}
			return newConn(conn, c.conf.HTTPTimeout.ReadWriteTimeout,
				c.conf.HTTPTimeout.LongTimeout), nil

		},
		MaxIdleConns:          c.conf.HTTPMaxConns.MaxIdleConns,
		MaxIdleConnsPerHost:   c.conf.HTTPMaxConns.MaxIdleConnsPerHost,
		IdleConnTimeout:       c.conf.HTTPTimeout.IdleConnTimeout,
		ResponseHeaderTimeout: c.conf.HTTPTimeout.HeaderTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: !c.conf.sslVerify},
	}
	if c.conf.proxyHost != "" {
		proxyUrl, err := url.Parse(c.conf.proxyHost)
		if err != nil {
			return nil, err
		}
		if c.conf.proxyUser != "" {
			if c.conf.proxyPwd != "" {
				proxyUrl.User = url.UserPassword(c.conf.proxyUser, c.conf.proxyPwd)
			} else {
				proxyUrl.User = url.User(c.conf.proxyUser)
			}
		}
		transport.Proxy = http.ProxyURL(proxyUrl)
	}
	return transport, nil
}

// WithApiVersion set api version
func (c *Client) WithApiVersion(apiVersion string) {
	if apiVersion != "" {
		c.conf.apiVersion = apiVersion
	}
}

// WithEndpoint set api endpoint
func (c *Client) WithEndpoint(endpoint string) {
	if endpoint != "" {
		c.conf.endpoint = endpoint
	}
}

// Do send http request
func (c Client) Do(method, path string, in interface{}, out interface{}) *Error {
	var requestBody []byte
	requestParams := make(url.Values)
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		body, err := json.Marshal(in)
		if err != nil {
			return &Error{
				Code:    -1,
				Message: err.Error(),
			}
		}
		requestBody = body
	default:
		query, err := queryMap(in)
		if err != nil {
			return &Error{
				Code:    -1,
				Message: err.Error(),
			}
		}
		requestParams = query
	}
	requestPath, err := c.requestPath(path, requestParams)
	if err != nil {
		return &Error{
			Code:    -1,
			Message: err.Error(),
		}
	}
	req, err := http.NewRequest(method, requestPath, bytes.NewBuffer(requestBody))
	if err != nil {
		return &Error{
			Code:    -1,
			Message: err.Error(),
		}
	}
	c.setDefaultHeader(req)
	return c.do(req, requestBody, out)
}

// send request
func (c Client) do(req *http.Request, requestBody []byte, out interface{}) *Error {
	resp, err := c.client.Do(req)
	if err != nil {
		return &Error{
			Code:    -1,
			Message: err.Error(),
		}
	}
	// parse error
	if req.Method != http.MethodHead {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &Error{
				Code:    -1,
				Message: err.Error(),
			}
		}
		defer resp.Body.Close()
		c.debug(requestBody, body, req, resp)
		var respError Error
		if err := json.Unmarshal(body, &respError); err != nil {
			return &Error{
				Code:    -1,
				Message: err.Error(),
			}
		}
		if respError.Code != 0 {
			return &respError
		}
		if out != nil {
			if err := json.Unmarshal(body, out); err != nil {
				return &Error{
					Code:    -1,
					Message: err.Error(),
				}
			}
		}
		return nil
	}
	return nil
}

func (c Client) debug(requestBody, responseBody []byte, req *http.Request, resp *http.Response) {
	if os.Getenv("DEBUG") != "" {
		fmt.Printf("> %s %s %s\n", req.Method, req.URL.RequestURI(), req.Proto)
		fmt.Printf("> Host: %s\n", req.Host)
		for key, header := range req.Header {
			for _, value := range header {
				fmt.Printf("> %s: %s\n", key, value)
			}
		}
		fmt.Println(">")
		fmt.Println(string(requestBody))
		fmt.Println(">")
		fmt.Printf("< %s %s\n", resp.Proto, resp.Status)
		for key, header := range resp.Header {
			for _, value := range header {
				fmt.Printf("< %s: %s\n", key, value)
			}
		}

		fmt.Println("< ")
		fmt.Println(string(responseBody))
		fmt.Println("< ")
	}
}

// set default headers
func (c Client) setDefaultHeader(req *http.Request) {
	req.Header.Add(HTTPHeaderUserAgent, userAgent())
	req.Header.Add(HTTPHeaderContentType, "application/json")
	req.Header.Add(HTTPHeaderToken, c.conf.secret)
}

// get request path
// if is onebox request add onebox param in query
// if ignore http code then add ihc param in query
// requestPath = apiVerison + requestPath + queryParams
func (c Client) requestPath(path string, params url.Values) (string, error) {
	requestPath := fmt.Sprintf("%s/%s/%s",
		strings.TrimRight(c.conf.endpoint, "/"),
		c.conf.apiVersion,
		strings.TrimLeft(path, "/"))
	requestUrl, err := url.Parse(requestPath)
	if err != nil {
		return "", err
	}
	values := requestUrl.Query()
	for k, v := range params {
		for _, v1 := range v {
			values.Add(k, v1)
		}
	}
	requestUrl.RawQuery = values.Encode()
	return requestUrl.String(), nil
}
