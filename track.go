package track17

import "net/http"

// Registe 运单注册
func (c Client) Registe(in []*TrackReq) (*TrackResp, *Error) {
	var resp TrackResp
	err := c.Do(http.MethodPost, "/register", in, &resp)
	return &resp, err
}

// ChangeCarrier 修改运输商
func (c Client) ChangeCarrier(in []*ChangeCarrierReq) (*TrackResp, *Error) {
	var resp TrackResp
	err := c.Do(http.MethodPost, "/changecarrier", in, &resp)
	return &resp, err
}

// Update 运单修改
func (c Client) Update(in []*TrackUpdateReq) (*TrackResp, *Error) {
	var resp TrackResp
	err := c.Do(http.MethodPost, "/changeinfo", in, &resp)
	return &resp, err
}

// StopTrack 停止追踪
func (c Client) StopTrack(in []*TrackReq) (*TrackResp, *Error) {
	var resp TrackResp
	err := c.Do(http.MethodPost, "/stoptrack", in, &resp)
	return &resp, err
}

// ReTrack 重启追踪
func (c Client) ReTrack(in []*TrackReq) (*TrackResp, *Error) {
	var resp TrackResp
	err := c.Do(http.MethodPost, "/retrack", in, &resp)
	return &resp, err
}

// GetTrackList 获取注册单号列表信息
func (c Client) GetTrackList(in *TrackListReq) (*TrackListResp, *Error) {
	var resp TrackListResp
	err := c.Do(http.MethodPost, "/gettracklist", in, &resp)
	return &resp, err
}

// GetTrackInfo 获取详情
func (c Client) GetTrackInfo(in []*TrackReq) (*TrackInfo, *Error) {
	var resp TrackInfo
	err := c.Do(http.MethodPost, "/gettrackinfo", in, &resp)
	return &resp, err
}
