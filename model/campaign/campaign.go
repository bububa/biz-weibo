package campaign

import (
	"fmt"
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// Campaign 广告系列
type Campaign struct {
	// ID 系列id
	ID int64 `json:"id,omitempty"`
	// Name 系列名称
	Name string `json:"name,omitempty"`
	// Objective 营销目标
	Objective model.Objective `json:"objective,omitempty"`
	// CustomerID 广告客户微博id
	CustomerID int64 `json:"customer_id,omitempty"`
	// EffectiveStatus 状态
	EffectiveStatus int `json:"effective_status,omitempty"`
	// ConfiguredStatus 状态
	ConfiguredStatus int `json:"configured_status,omitempty"`
	// GuaranteedDelivery 是否是定价保量 1是 0否
	GuaranteedDelivery int `json:"guaranteed_delivery,omitempty"`
	// LifetimeBudget 系列限额
	LifetimeBudget string `json:"lifetime_budget,omitempty"`
	// IsRecycle 是否在回收站 1是 0否
	IsRecycle int `json:"is_recycle,omitempty"`
	// Oauth2ID
	Oauth2ID int64 `json:"oauth2_id,omitempty"`
	// AdBuilder
	AdBuilder int `json:"ad_builder,omitempty"`
	// CreateTime
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime
	UpdateTime string `json:"update_time,omitempty"`
}

// CampaignRequest 系列详情
type CampaignRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// ID 系列ID
	ID int64 `json:"id,omitempty"`
}

// URL implement Request interface
func (r CampaignRequest) URL() string {
	return fmt.Sprintf("v3/campaigns/%d", r.ID)
}

// Payload implement Request interface
func (r CampaignRequest) Payload() *model.Payload {
	if r.AdvertiserID <= 0 {
		return nil
	}
	p := new(model.Payload)
	p.AddValue("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	return p
}

// CampaignResponse 系列详情API Response
type CampaignResponse struct {
	model.BaseResponse
	Campaign
}
