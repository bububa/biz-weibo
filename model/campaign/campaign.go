package campaign

import (
	"fmt"
	"strconv"

	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
)

// Campaign 广告系列
type Campaign struct {
	Id                 uint64                       `json:"id,omitempty"`
	CustomerId         uint64                       `json:"customer_id,omitempty"`
	Name               string                       `json:"name,omitempty"`
	Objective          int                          `json:"objective,omitempty"`
	GuaranteedDelivery int                          `json:"guaranteed_delivery,omitempty"`
	ConfiguredStatus   enum.ConfiguredStatus        `json:"configured_status,omitempty"`
	EffectiveStatus    enum.CampaignEffectiveStatus `json:"effective_status,omitempty"`
	BudgetType         enum.BudgetType              `json:"budget_type,omitempty"`
	LifetimeBudget     model.Float64                `json:"lifetime_budget,omitempty"`
	DailyBudget        model.Float64                `json:"daily_budget,omitempty"`
	IsRecycled         int                          `json:"is_recycled,omitempty"`
}

// CampaignRequest 系列详情
type CampaignRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ID 系列ID
	ID uint64 `json:"id,omitempty"`
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
	p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return p
}

// CampaignResponse 系列详情API Response
type CampaignResponse struct {
	model.BaseResponse
	Campaign
}
