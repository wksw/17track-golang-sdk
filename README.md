> 17track golang版本SDK
> 官方文档请参考[api文档](https://console.17track.net/zh-cn/doc)

### Install

```bash
go get github.com/wksw/17track-golang-sdk
```

### Install 17track command

```bash
go get github.com/wksw/17track-golang-sdk/cmd/17track
```

### Get Latest carriers and countries

```bash
go generate
```

### Registe

```golang
    client, err := track17.NewClient("api-key")
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
```