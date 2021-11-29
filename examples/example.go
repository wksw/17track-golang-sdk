package main

import (
	"log"

	track17 "github.com/wksw/17track-golang-sdk"
	pb "github.com/wksw/17track-golang-sdk/proto"
)

func main() {
	client, err := track17.NewClient("api-key", "https://api.17track.net/")
	if err != nil {
		log.Fatal(err.Error())
	}
	resp, rerr := client.Registe([]*pb.TrackReq{
		{Number: "RR123456789CN", Carrier: 3011},
	})
	if rerr != nil {
		log.Fatalf("%v", rerr)
	}
	log.Printf("%v\n", resp)
}
