package creative

import (
	"strconv"
	"strings"

	"github.com/bububa/biz-weibo/model"
)

// BatchRequest 批量获取创意详情 API Request
type BatchRequest struct {
	// AdvertiserID  您管理的广告主id，为空则为token对应账户uid
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDs 创意id
	IDs []uint64 `json:"ids,omitempty"`
}

// URL implement Request interface
func (r BatchRequest) URL() string {
	return "v4/creatives/batch"
}

// Payload implement Request interface
func (r BatchRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	ids := make([]string, 0, len(r.IDs))
	for _, id := range r.IDs {
		ids = append(ids, strconv.FormatUint(id, 10))
	}
	p.AddValue("ids", strings.Join(ids, ","))
	return p
}

// BatchResponse 批量获取创意详情 API Response
type BatchResponse struct {
	model.BaseResponse
	Data []CreativeDetail `json:"data,omitempty"`
}
