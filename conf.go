package track17

import (
	"time"
)

// Config paasport configuration
type Config struct {
	// gateway endpoint
	endpoint     string
	secret       string
	apiVersion   string
	userAgent    string
	proxyHost    string
	proxyUser    string
	proxyPwd     string
	HTTPTimeout  HTTPTimeout
	HTTPMaxConns HTTPMaxConns
	sslVerify    bool
}

// HTTPTimeout http timeout define
type HTTPTimeout struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	HeaderTimeout    time.Duration
	LongTimeout      time.Duration
	IdleConnTimeout  time.Duration
}

// HTTPMaxConns max idle connections
type HTTPMaxConns struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
}

type configurer func(conf *Config)

func (c *Config) withDefault() {
	if c.apiVersion == "" {
		c.apiVersion = DEFAULT_API_VERSION
	}
	if c.HTTPTimeout.ReadWriteTimeout == 0 {
		c.HTTPTimeout.ReadWriteTimeout = 5 * time.Second
	}
}

// WithSslVerify set sslVerify
func WithSslVerify(sslVerify bool) configurer {
	return func(conf *Config) {
		conf.sslVerify = sslVerify
	}
}

// WithTimeout set timeout
func WithTimeout(httpTimeout HTTPTimeout) configurer {
	return func(conf *Config) {
		conf.HTTPTimeout = httpTimeout
	}
}

// WithProxy set proxy
func WithProxy(proxyHost, proxyUser, proxyPwd string) configurer {
	return func(conf *Config) {
		conf.proxyHost = proxyHost
		conf.proxyUser = proxyUser
		conf.proxyPwd = proxyPwd
	}
}

// WithMaxConnections set max connections
func WithMaxConnections(httpMaxConns HTTPMaxConns) configurer {
	return func(conf *Config) {
		conf.HTTPMaxConns = httpMaxConns
	}
}

// WithApiVersion set api version
func WithApiVersion(apiVersion string) configurer {
	return func(conf *Config) {
		conf.apiVersion = apiVersion
	}
}
