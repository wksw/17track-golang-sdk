/*
	17track数据结构定义， 详细信息请参考https://console.17track.net/zh-cn/doc
*/

package track17

// Error 错误返回
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Errors []RejectError `json:"errors"`
	} `json:"data"`
}

// TrackAccept 成功返回
//go:generate 17track
type TrackAccept struct {
	Origin  int32   `json:"origin"`
	Number  string  `json:"number"`
	Carrier Carrier `json:"carrier"`
}

// TrackReject 失败返回
type TrackReject struct {
	Number string      `json:"number"`
	Error  RejectError `json:"error"`
}

// RejectError 失败错误信息
type RejectError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// TrackReq 运单请求
type TrackReq struct {
	Number        string  `json:"number"`
	Carrier       Carrier `json:"carrier"`
	FinalCarrier  Carrier `json:"final_carrier"`
	AutoDetection bool    `json:"auto_detection"`
	Tag           string  `json:"tag"`
}

// TrackResp 运单返回
type TrackResp struct {
	Code int32 `json:"code"`
	Data struct {
		Accepted []TrackAccept `json:"accepted"`
		Rejected []TrackReject `json:"rejected"`
	} `json:"data"`
}

// ChangeCarrierReq 修改运输商请求
type ChangeCarrierReq struct {
	Number          string  `json:"number"`
	OldCarrier      Carrier `json:"carrier_old"`
	NewCarrier      Carrier `json:"carrier_new"`
	OldFinalCarrier Carrier `json:"final_carrier_old"`
	NewFinalCarrier Carrier `json:"final_carrier_new"`
}

// TrackUpdateReq 运单信息更新请求
type TrackUpdateReq struct {
	Number  string  `json:"number"`
	Carrier Carrier `json:"carrier"`
	Items   struct {
		Tag string `json:"tag"`
	} `json:"items"`
}

// PackageState 包裹状态
type PackageState int32

const (
	// PackageStateNotFound 查询不到
	PackageStateNotFound PackageState = 0
	// PackageStateInTransit 运输中
	PackageStateInTransit PackageState = 10
	// PackageStateTooLong 运输过久
	PackageStateTooLong PackageState = 20
	// PackageStateArrived 到达待取
	PackageStateArrived PackageState = 30
	// PackageStateFail 投递失败
	PackageStateFail PackageState = 35
	// PackageStateSuccess 成功签收
	PackageStateSuccess PackageState = 40
	// PackageStateAbnormal 可能异常
	PackageStateAbnormal PackageState = 50
)

// TrackingState 追踪状态
type TrackingState int32

const (
	// TrackingStateRunning 追踪中
	TrackingStateRunning TrackingState = 1
	// TrackingStateStop 停止追踪
	TrackingStateStop TrackingState = 0
)

// TrackListReq 运单列表请求
type TrackListReq struct {
	Number           string        `json:"number"`
	Carrier          Carrier       `json:"carrier"`
	RegisteTimeFrom  string        `json:"registe_time_from"`
	RegisteTimeTo    string        `json:"registe_time_to"`
	TrackingTimeFrom string        `json:"tracking_time_from"`
	TrackingTimeTo   string        `json:"tracking_time_to"`
	PushTimeFrom     string        `json:"push_time_from"`
	PushTimeTo       string        `json:"push_time_to"`
	PushState        int32         `json:"push_state"`
	StopTimeFrom     string        `json:"stop_time_from"`
	StopTimeTo       string        `json:"stop_time_to"`
	PackageState     PackageState  `json:"package_state"`
	TrackingState    TrackingState `json:"tracking_state"`
	PageNo           int32         `json:"page_no"`
}

// StopTrackReason 停止追踪原因
type StopTrackReason int32

const (
	// StopTrackReasonExpired 过期自动停止
	StopTrackReasonExpired StopTrackReason = 1
	// StopTrackReasonAPI 接口请求停止
	StopTrackReasonAPI StopTrackReason = 2
	// StopTrackReasonArtificial 手工操作
	StopTrackReasonArtificial StopTrackReason = 3
	// StopTrackReasonInvalidCarrier 无效的运输商
	StopTrackReasonInvalidCarrier StopTrackReason = 4
)

// TrackListResp 运单列表返回
type TrackListResp struct {
	Code int32 `json:"code"`
	Data struct {
		Accepted []struct {
			// 跟踪单号
			Number string          `json:"number"`
			W1     int32           `json:"w1"`
			W2     int32           `json:"w2"`
			B      int32           `json:"b"`
			C      int32           `json:"c"`
			E      int32           `json:"e"`
			RT     string          `json:"rt"`
			TT     string          `json:"tt"`
			PT     string          `json:"pt"`
			PS     int32           `json:"ps"`
			ST     string          `json:"st"`
			SR     StopTrackReason `json:"sr"`
			IR     bool            `json:"ir"`
			TS     bool            `json:"ts"`
			MC     int32           `json:"mc"`
			Tag    string          `json:"tag"`
		} `json:"accepted"`
		Rejected []TrackReject `json:"rejected"`
	} `json:"data"`
}

// CarrierState 运输商状态
type CarrierState int32

const (
	// CarrierStateUnrecognized 无法识别
	CarrierStateUnrecognized CarrierState = 0
	// CarrierStateSuccess 正常查有信息
	CarrierStateSuccess CarrierState = 1
	// CarrierStateNotFound 尚无信息
	CarrierStateNotFound CarrierState = 2
	// CarrierStateSiteErr 网站出错
	CarrierStateSiteErr CarrierState = 10
	// CarrierStateProcessErr 处理出错
	CarrierStateProcessErr CarrierState = 11
	// CarrierStateFetchErr 查询出错
	CarrierStateFetchErr CarrierState = 12
	// CarrierStateSiteErrUseCache 网站错误,使用缓存
	CarrierStateSiteErrUseCache CarrierState = 20
	// CarrierStateProcessErrUseCache 处理错误,使用缓存
	CarrierStateProcessErrUseCache CarrierState = 21
	// CarrierStateFetchErrUseCache 查询出错,使用缓存
	CarrierStateFetchErrUseCache CarrierState = 22
)

// TrackInfo 物流详情
type TrackInfo struct {
	Code int32 `json:"code"`
	Data struct {
		Accepted []struct {
			Number string `json:"number"`
			Tag    string `json:"tag"`
			Track  struct {
				W1   int32        `json:"w1"`
				W2   int32        `json:"w2"`
				B    int32        `json:"b"`
				C    int32        `json:"c"`
				E    int32        `json:"e"`
				F    int32        `json:"f"`
				Z0   []TrackEvent `json:"z0"`
				Z1   []TrackEvent `json:"z1"`
				Z2   []TrackEvent `json:"z2"`
				Z9   []TrackEvent `json:"z9"`
				Ygt1 int32        `json:"ygt1"`
				Ygt2 int32        `json:"ygt2"`
				Ygt9 int32        `json:"ygt9"`
				Ylt1 int32        `json:"ylt1"`
				Ylt2 int32        `json:"ylt2"`
				Ylt9 int32        `json:"ylt9"`
				IS1  CarrierState `json:"is1"`
				IS2  CarrierState `json:"is2"`
				Ln1  string       `json:"ln1"`
				Ln2  string       `json:"ln2"`
				Ln9  string       `json:"ln9"`
				Hs   string       `json:"hs"`
				Yt   string       `json:"yt"`
				Zex  TrackExpand  `json:"zex"`
			} `json:"track"`
		} `json:"accepted"`
		Rejected []TrackReject `json:"rejected"`
	} `json:"data"`
}

// TrackExpand 追踪扩展
type TrackExpand struct {
	TRC  string `json:"trC"`
	TRN  string `json:"trN"`
	DTS  string `json:"dtS"`
	DTP  string `json:"dtP"`
	DTD  string `json:"dtD"`
	DTL  string `json:"dtL"`
	DT   string `json:"dt"`
	PSex string `json:"psex"`
}

// TrackEvent 跟踪事件
type TrackEvent struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D string `json:"d"`
	Z string `json:"z"`
}
