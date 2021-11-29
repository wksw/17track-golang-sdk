package track17

import (
	"net/http"

	pb "github.com/wksw/17track-golang-sdk/proto"
)

// Registe 运单注册
func (c Client) Registe(in []*pb.TrackReq) (*pb.TrackResp, *Error) {
	var resp pb.TrackResp
	err := c.Do(http.MethodPost, "/register", in, &resp)
	return &resp, err
}

// ChangeCarrier 修改运输商
func (c Client) ChangeCarrier(in []*pb.ChangeCarrierReq) (*pb.TrackResp, *Error) {
	var resp pb.TrackResp
	err := c.Do(http.MethodPost, "/changecarrier", in, &resp)
	return &resp, err
}

// Update 运单修改
func (c Client) Update(in []*pb.TrackUpdateReq) (*pb.TrackResp, *Error) {
	var resp pb.TrackResp
	err := c.Do(http.MethodPost, "/changeinfo", in, &resp)
	return &resp, err
}

// StopTrack 停止追踪
func (c Client) StopTrack(in []*pb.TrackReq) (*pb.TrackResp, *Error) {
	var resp pb.TrackResp
	err := c.Do(http.MethodPost, "/stoptrack", in, &resp)
	return &resp, err
}

// ReTrack 重启追踪
func (c Client) ReTrack(in []*pb.TrackReq) (*pb.TrackResp, *Error) {
	var resp pb.TrackResp
	err := c.Do(http.MethodPost, "/retrack", in, &resp)
	return &resp, err
}

// GetTrackList 获取注册单号列表信息
func (c Client) GetTrackList(in *pb.TrackListReq) (*pb.TrackListResp, *Error) {
	var resp pb.TrackListResp
	err := c.Do(http.MethodPost, "/gettracklist", in, &resp)
	return &resp, err
}

// GetTrackInfo 获取详情
func (c Client) GetTrackInfo(in []*pb.TrackReq) (*pb.TrackInfo, *Error) {
	var resp pb.TrackInfo
	err := c.Do(http.MethodPost, "/gettrackinfo", in, &resp)
	return &resp, err
}
