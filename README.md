> 17track golang版本SDK
> 官方文档请参考[api文档](https://console.17track.net/zh-cn/doc)

### Install

```bash
go get github.com/wksw/17track-golang-sdk
```

### Install 17track command

```bash
go get github.com/wksw/17track-golang-sdk/cmd/17track
# Starting in Go 1.17
go install github.com/wksw/17track-golang-sdk/cmd/17track
```

### Get Latest carriers and countries

```bash
# it's actually execute 17track command
go generate

# Or execute 17track command
# it will generage two files, 17track_carrier.go and 17track_country.go
# at current dir
17track
```

### Examples

#### V1版本

##### Registe

```golang
import (
   track17 "github.com/wksw/17track-golang-sdk"
   pb "github.com/wksw/17track-golang-sdk/proto"
)
func main() {
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
}
```

#### V2版本

##### Registe

```golang
import (
   track17V2 "github.com/wksw/17track-golang-sdk/v2"
   pb "github.com/wksw/17track-golang-sdk/models/v2"
)
func main() {
    client, err := track17V2.NewClient("api-key")
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
```

##### v1&v2同时使用

```golang

import (
    // v2版本
   v2 "github.com/wksw/17track-golang-sdk/v2"
   modelsV2 "github.com/wksw/17track-golang-sdk/models/v2"
   // v1版本
   v1 "github.com/wksw/17track-golang-sdk"
   modelsV2 "github.com/wksw/17track-golang-sdk/proto"
)
func main() {
    // v1版本客户端
    clientV1, err := v1.NewClient("api-key")
    if err != nil {
        log.Fatal(err.Error())
    }
    // v1版本注册
    respV1, rerrV1 := clientV1.Registe([]*modelsV1.TrackReq{
        {Number: "RR123456789CN", Carrier: 3011},
    })
    if rerrV1 != nil {
        log.Fatalf("%v", rerrV1)
    }
    log.Printf("%v\n", respV1)

    // v2版本客户端
    clientV2 := v2.Client{C: clientV1}
    // v2版本注册
    respV2, rerrV2 := clientV2.Registe([]*modelsV2.TrackReq{
        {Number: "RR123456789CN", Carrier: 3011},
    })
    if rerrV2 != nil {
        log.Fatalf("%v", rerrV2)
    }
    log.Printf("%v\n", respV2)
}
```
