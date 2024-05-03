package creative

import (
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// AdRecRequest 智能创意文案 API Request
type AdRecRequest struct {
	// AdvertiserID  您管理的广告主id，为空则为token对应账户uid
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SearchKey 查询关键词
	SearchKey string `json:"search_key,omitempty"`
}

// URL implement Request interface
func (r AdRecRequest) URL() string {
	return "v4/creatives/tool/adrec"
}

// Payload implement Request interface
func (r AdRecRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	p.AddValue("search_key", r.SearchKey)
	return p
}

// AdRecResponse 智能创意文案 API Response
type AdRecResponse struct {
	model.BaseResponse
	Data []AdRec `json:"data,omitempty"`
}

type AdRec struct {
	// ActualContent 实际正文内容投放使用
	ActualContent string `json:"actual_content,omitempty"`
	// Content 正文内容
	Content string `json:"content,omitempty"`
}
