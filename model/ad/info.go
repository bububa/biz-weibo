package ad

import (
	"strconv"

	"github.com/bububa/biz-weibo/model"
	"github.com/bububa/biz-weibo/util"
)

// InfoRequest 广告计划详情 API Request
type InfoRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// AdID 广告ID
	AdID uint64 `json:"ad_id,omitempty"`
}

// URL implement Request interface
func (r InfoRequest) URL() string {
	return util.StringsJoin("v3/ads/", strconv.FormatUint(r.AdID, 10))
}

// Payload implement Request interface
func (r InfoRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	}
	return p
}

// InfoResponse 广告计划详情 API Response
type InfoResponse struct {
	model.BaseResponse
	Ad
}
