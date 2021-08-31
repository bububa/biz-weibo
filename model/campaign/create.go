package campaign

import (
	"encoding/json"

	"github.com/bububa/biz-weibo/model"
)

// CreateRequest 创建广告系列
type CreateRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// Name 系列名称 最长50个字符
	Name string `json:"name,omitempty"`
	// Objective 营销目标
	Objective model.Objective `json:"objective,omitempty"`
	// GuaranteedDelivery  是否定价保量，1为定价保量，0为非定价保量（默认）
	GuaranteedDelivery bool `json:"guaranteed_delivery,omitempty"`
	// LifetimeBudget 系列总预算，0表示不限，最低系列总预算为150元，最高系列总预算9,999,999.99元
	LifetimeBudget float64 `json:"lifetime_budget,omitempty"`
}

// URL implement Request interface
func (r CreateRequest) URL() string {
	return "v3/campaigns"
}

// Payload implement Request interface
func (r CreateRequest) Payload() *model.Payload {
	buf, _ := json.Marshal(r)
	return model.NewPostPayload(buf)
}

// CreateResponse 创建广告系列 API Response
type CreateResponse struct {
	model.BaseResponse
	Campaign
}
