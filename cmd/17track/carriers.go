package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Carrier 运输商
type Carrier struct {
	Key          int32       `json:"key,string"`
	CanTrack     json.Number `json:"_canTrack"`
	Country      int32       `json:"_country,string"`
	Url          string      `json:"_url"`
	Name         string      `json:"_name"`
	Code         string      `json:"_code"`
	ExpressGroup string      `json:"_expressGroup"`
	IconBgColor  string      `json:"_iconBgColor"`
}

func genCarriers(g *Generator) {
	// Print the header and package clause.
	g.Printf("// Code generated by \"17track %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
	g.Printf("\n")
	g.Printf("package %s", g.pkg.name)
	g.Printf("\n")
	// g.Printf("import \"strconv\"\n") // Used by all methods.

	// Run generate for each type.
	g.Printf("// Carrier the carrier define \n")
	g.Printf("type Carrier int32")
	g.Printf("\n")

	g.Printf("// Carriers 17track defined carrier list\n")
	g.Printf(`var Carriers = map[Carrier]struct{
		Key int32
		CanTrack bool
		Country int32
		Url string
		Name string 
		Code string
		ExpressGroup string
		IconBgColor string}{`)
	g.Printf("\n")
	resp, err := http.Get("https://www.17track.net/zh-cn/apicarrier")
	if err != nil {
		log.Fatalf("get carriers api fail[%s]", err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read response body fail[%s]", err.Error())
	}
	defer resp.Body.Close()

	var carriers []Carrier
	if err := json.Unmarshal(body, &carriers); err != nil {
		log.Fatalf("unmarshal response body '%s' fail[%s]", string(body), err.Error())
	}

	for _, carrier := range carriers {
		canTrack := false
		if carrier.CanTrack.String() == "1" {
			canTrack = true
		}
		g.Printf("%d: {Key: %d, CanTrack: %v, Country: %d, Url: \"%s\", Name: \"%s\", Code: \"%s\", ExpressGroup: \"%s\", IconBgColor: \"%s\"},",
			carrier.Key, carrier.Key,
			canTrack,
			carrier.Country,
			carrier.Url,
			carrier.Name,
			carrier.Code,
			carrier.ExpressGroup,
			carrier.IconBgColor)
		g.Printf("\n")

	}

	g.Printf("}")
}