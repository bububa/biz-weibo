package account

import (
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// InfoRequest 获取广告账户信息 API Request
type InfoRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// URL implement Request interface
func (r InfoRequest) URL() string {
	return "v3/account"
}

// Payload implement Request interface
func (r InfoRequest) Payload() *model.Payload {
	if r.AdvertiserID <= 0 {
		return nil
	}
	p := new(model.Payload)
	p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return p
}

// InfoResponse 获取广告账户信息 API Response
type InfoResponse struct {
	model.BaseResponse
	Account
}
