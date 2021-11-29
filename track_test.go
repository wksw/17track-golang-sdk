package track17

import (
	"testing"

	pb "github.com/wksw/17track-golang-sdk/proto"
)

func TestRegiste(t *testing.T) {
	// apiKey is defined in file config_test.go, it will be ignored when push code into remote repo
	client, err := NewClient(apiKey)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, rerr := client.Registe([]*pb.TrackReq{
		{Number: "RR123456789CN", Carrier: 3011},
	})
	if rerr != nil {
		t.Errorf("%+v", rerr)
		t.FailNow()
	}
	t.Logf("%+v", resp)
	if len(resp.Data.Accepted) != 1 {
		t.Error("registe no response")
		t.FailNow()
	}
}

func TestChangeCarrier(t *testing.T) {
	client, err := NewClient(apiKey)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, rerr := client.ChangeCarrier([]*pb.ChangeCarrierReq{
		{Number: "RR123456789CN", CarrierOld: 3011, CarrierNew: 3012},
	})
	if rerr != nil {
		t.Errorf("%+v", rerr)
		t.FailNow()
	}
	t.Logf("%+v", resp)
}
