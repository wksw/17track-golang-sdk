package models

// TrackReq 物流追踪请求
type TrackReq struct {
	Number        string `json:"number"`
	Param         string `json:"param"`
	Carrier       int32  `json:"carrier"`
	FinalCarrier  int32  `json:"final_carrier"`
	AutoDetection bool   `json:"auto_detection"`
	Tag           string `json:"tag"`

	Lang       string `json:"lang,omitempty"`
	Email      string `json:"email,omitempty"`
	OrderNo    string `json:"order_no,omitempty"`
	OrderTime string `json:"order_time,omitempty"`
	Remark     string `json:"remark,omitempty"`
}

// TrackResp 物流追踪返回
type TrackResp struct {
	Code int32          `json:"code"`
	Data *TrackRespData `json:"data"`
}

// TrackRespData 物流追踪返回
type TrackRespData struct {
	Accepted []*TrackRespAccept `json:"accepted"`
	Rejected []*TrackRespReject `json:"rejected"`
}

// TrackRespAccept 物流追踪成功返回
type TrackRespAccept struct {
	Origin    int32      `json:"origin"`
	Number    string     `json:"number"`
	Tag       string     `json:"tag"`
	Carrier   int32      `json:"carrier"`
	Param     string     `json:"param"`
	TrackInfo *TrackInfo `json:"track_info"`
}

// TrackRespReject 物流追踪错误返回
type TrackRespReject struct {
	Number string          `json:"number"`
	Tag    string          `json:"tag"`
	Error  *TrackRespError `json:"error"`
}

// TrackRespError 错误返回详情
type TrackRespError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// TrackInfo 物流追踪详情
type TrackInfo struct {
	ShippingInfo struct {
		ShipperAddress   TrackLocation `json:"shipper_address"`
		RecipientAddress TrackLocation `json:"recipient_address"`
	} `json:"shipping_info"`
	LatestStatus struct {
		Status         string `json:"status"`
		SubStatus      string `json:"sub_status"`
		SubStatusDescr string `json:"sub_status_descr"`
	} `json:"latest_status"`
	LatestEvent *TrackEvent `json:"latest_event"`
	TimeMetrics struct {
		DaysAfterOrder        int32 `json:"days_after_order"`
		DaysOfTransit         int32 `json:"days_of_transit"`
		DaysOfTransitDone     int32 `json:"days_of_transit_done"`
		DaysAfterLastUpdate   int32 `json:"days_after_last_update"`
		EstimatedDeliveryDate struct {
			Source string `json:"source"`
			From   string `json:"from"`
			To     string `json:"to"`
		} `json:"estimated_delivery_date"`
	} `json:"time_metrics"`
	Milestone []struct {
		KeyStage string `json:"key_stage"`
		TimeISO  string `json:"time_iso"`
		TimeUTC  string `json:"time_utc"`
	} `json:"milestone"`
	MiscInfo struct {
		RiskFactor      int32  `json:"risk_factor"`
		ServiceType     string `json:"service_type"`
		WeightRaw       string `json:"weight_raw"`
		WeightKg        string `json:"weight_kg"`
		Pieces          string `json:"pieces"`
		Dimensions      string `json:"dimensions"`
		CustomerNumber  string `json:"customer_number"`
		ReferenceNumber string `json:"reference_number"`
		LocalNumber     string `json:"local_number"`
		LocalProvider   string `json:"local_provider"`
		LocalKey        int32  `json:"local_key"`
	} `json:"misc_info"`
	Tracking struct {
		ProvidersHash int64 `json:"providers_hash"`
		Providers     []struct {
			Provider         TrackProvider `json:"provider"`
			ServiceType      string        `json:"service_type"`
			LatestSyncStatus string        `json:"latest_sync_status"`
			LatestSyncTime   string        `json:"latest_sync_time"`
			EventsHash       int64         `json:"events_hash"`
			Events           []*TrackEvent `json:"events"`
		} `json:"providers"`
	} `json:"tracking"`
}

// TrackLocation 物流追踪地区信息
type TrackLocation struct {
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
	Street      string `json:"street"`
	PostalCode  string `json:"postal_code"`
	Coordinates struct {
		Long float64 `json:"longitude"`
		Lat  float64 `json:"latitude"`
	} `json:"coordinates"`
}

// TrackProvider 物流追踪运输商
type TrackProvider struct {
	Key      int32  `json:"key"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Tel      string `json:"tel"`
	Homepage string `json:"homepage"`
	Country  string `json:"country"`
}

// TrackEvent 物流追踪事件
type TrackEvent struct {
	TimeISO     string        `json:"time_iso"`
	TimeUTC     string        `json:"time_utc"`
	Description string        `json:"description"`
	Location    string        `json:"location"`
	Stage       string        `json:"stage"`
	Address     TrackLocation `json:"address"`
}
