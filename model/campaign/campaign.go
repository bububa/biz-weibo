package campaign

import (
	"fmt"
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// Campaign 广告系列
type Campaign struct {
	Id                 int    `json:"id"`
	CustomerId         int64  `json:"customer_id"`
	Name               string `json:"name"`
	Objective          int    `json:"objective"`
	LifetimeBudget     string `json:"lifetime_budget"`
	GuaranteedDelivery int    `json:"guaranteed_delivery"`
	ConfiguredStatus   int    `json:"configured_status"`
	EffectiveStatus    int    `json:"effective_status"`
	IsRecycled         int    `json:"is_recycled"`
	RecycleDatetime    string `json:"recycle_datetime"`
	AdBuilder          int    `json:"ad_builder"`
	Oauth2Id           int    `json:"oauth2_id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	Version            int    `json:"version"`
	FundType           int    `json:"fund_type"`
	AdDeliveryType     int    `json:"ad_delivery_type"`
	BudgetType         int    `json:"budget_type"`
	DailyBudget        string `json:"daily_budget"`
	ObjectiveName      string `json:"objective_name"`
}

// CampaignRequest 系列详情
type CampaignRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
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
	p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return p
}

// CampaignResponse 系列详情API Response
type CampaignResponse struct {
	model.BaseResponse
	Campaign
}
