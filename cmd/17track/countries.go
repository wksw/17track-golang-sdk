package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Country country define
type Country struct {
	Key  int32  `json:"key,string"`
	Name string `json:"_name"`
}

func genCountries(g *Generator) {
	// Print the header and package clause.
	g.Printf("// Code generated by \"17track %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
	g.Printf("\n")
	g.Printf("package %s", g.pkg.name)
	g.Printf("\n")
	// g.Printf("import \"strconv\"\n") // Used by all methods.

	// Run generate for each type.
	// g.Printf("// Country the country define \n")
	// g.Printf("type Country int32")
	// g.Printf("\n")

	g.Printf("// Countries 17track defined country list\n")
	g.Printf(`var Countries = map[int32]struct{
		Key int32
		Name string}{`)
	g.Printf("\n")
	resp, err := http.Get("https://www.17track.net/zh-cn/apicountry")
	if err != nil {
		log.Fatalf("get carriers api fail[%s]", err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read response body fail[%s]", err.Error())
	}
	defer resp.Body.Close()

	var countryies []Country
	if err := json.Unmarshal(body, &countryies); err != nil {
		log.Fatalf("unmarshal response body '%s' fail[%s]", string(body), err.Error())
	}

	for _, country := range countryies {

		g.Printf("%d: {Key: %d, Name: \"%s\"},",
			country.Key, country.Key,
			country.Name)
		g.Printf("\n")

	}

	g.Printf("}")
}
