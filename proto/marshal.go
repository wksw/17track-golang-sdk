package track17

import (
	"encoding/json"
)

// Int64OrStringToInt64 int64类型或者string类型转为int64类型
type Int64OrStringToInt64 struct {
	Val int64
}

// UnmarshalJSON 反序列化
func (intstr *Int64OrStringToInt64) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		return json.Unmarshal(value[1:len(value)-1], &intstr.Val)
	}
	return json.Unmarshal(value, &intstr.Val)
}

type trackExpandTmp struct {
	TrC  int32                 ` json:"trC"`
	TrN  string                ` json:"trN" `
	DtS  *Int64OrStringToInt64 ` json:"dtS"`
	DtP  *Int64OrStringToInt64 `json:"dtP"`
	DtD  *Int64OrStringToInt64 `json:"dtD"`
	DtL  *Int64OrStringToInt64 ` json:"dtL" `
	Dt   *Int64OrStringToInt64 `json:"dt" `
	Psex int32
}

// UnmarshalJSON 反序列化,将dt,dtD,dtL,dtP,dtS反序列化为int64整形
func (t *TrackExpand) UnmarshalJSON(value []byte) error {
	var tmp trackExpandTmp
	if err := json.Unmarshal(value, &tmp); err != nil {
		return err
	}
	t.TrC = tmp.TrC
	t.TrN = tmp.TrN
	t.DtS = tmp.DtS.Val
	t.DtP = tmp.DtP.Val
	t.DtD = tmp.DtD.Val
	t.DtL = tmp.DtL.Val
	t.Dt = tmp.Dt.Val
	t.Psex = tmp.Psex
	return nil
}
