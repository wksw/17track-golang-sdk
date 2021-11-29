package track17

import (
	"testing"
)

func TestRegiste(t *testing.T) {
	client, err := NewClient("shit api define", "https://api.17track.net/")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, rerr := client.Registe([]*TrackReq{
		{Number: "RR123456789CN", Carrier: 3011},
	})
	if rerr != nil {
		t.Errorf("%+v", rerr)
		t.FailNow()
	}
	t.Logf("%+v", resp)
}

func TestChangeCarrier(t *testing.T) {
	client, err := NewClient("shit api define", "https://api.17track.net/")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	resp, rerr := client.ChangeCarrier([]*ChangeCarrierReq{
		{Number: "RR123456789CN", OldCarrier: 3011, NewCarrier: 3012},
	})
	if rerr != nil {
		t.Errorf("%+v", rerr)
		t.FailNow()
	}
	t.Logf("%+v", resp)
}
