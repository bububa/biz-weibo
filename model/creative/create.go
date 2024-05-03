package creative

import (
	"encoding/json"

	"github.com/bububa/biz-weibo/model"
)

// CreateRequest 创建创意 API Request
type CreateRequest struct {
	CreativeDetail
	// AdvertiserID  您管理的广告主id，为空则为token对应账户uid
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// URL implement Request interface
func (r CreateRequest) URL() string {
	return "v4/creatives"
}

// Payload implement Request interface
func (r CreateRequest) Payload() *model.Payload {
	bs, _ := json.Marshal(r)
	return model.NewPostPayload(bs)
}

// CreateResponse 创建广告创意 API Response
type CreateResponse struct {
	model.BaseResponse
	Data uint64 `json:"data,omitempty"`
}
