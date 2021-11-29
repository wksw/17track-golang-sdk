package track17

// Carrier 运输商
type Carrier int32

const (
	// AfghanPost AfghanPost
	AfghanPost Carrier = 01021
)

// Carriers 供应商列表
var Carriers = map[Carrier]struct {
	Key           int32
	Country       int32
	Url           string
	Name          string
	Code          string
	ExpiressGroup string
	IconBgColor   string
}{}
