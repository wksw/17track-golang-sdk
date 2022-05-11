package main

import (
	"encoding/json"
	"log"

	modelsv2 "github.com/wksw/17track-golang-sdk/models/v2"
	v2 "github.com/wksw/17track-golang-sdk/v2"
)

func main() {
	client, err := v2.NewClient("xxx")
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, rerr := client.GetInfo([]*modelsv2.TrackReq{{Number: "MZ147501922"}})
	if err != nil {
		log.Printf("getInfo fail[%s]", rerr.Message)
		return
	}
	// log.Printf("resp: %+v Data: %+v", resp, resp.Data)
	body, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Printf("marshal resp fail[%s]", err.Error())
		return
	}
	log.Println(string(body))
}
