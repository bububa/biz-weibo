package creative

import (
	"strconv"

	"github.com/bububa/biz-weibo/model"
	"github.com/bububa/biz-weibo/util"
)

// InfoRequest 获取创意详情 API Request
type InfoRequest struct {
	// AdvertiserID  您管理的广告主id，为空则为token对应账户uid
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ID 创意id
	ID uint64 `json:"id,omitempty"`
}

// URL implement Request interface
func (r InfoRequest) URL() string {
	return util.StringsJoin("v4/creatives/%d", strconv.FormatUint(r.ID, 10))
}

// Payload implement Request interface
func (r InfoRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	return p
}

// InfoResponse 广告创意详情 API Response
type InfoResponse struct {
	model.BaseResponse
	Creative
}
