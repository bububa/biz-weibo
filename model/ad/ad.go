package ad

import "github.com/bububa/biz-weibo/model"

// Ad 广告计划
type Ad struct {
	// ID 计划id
	ID int64 `json:"id,omitempty"`
	// Name 计划名称（唯一）
	Name string `json:"name,omitempty"`
	// CampaignID 所属广告系列id
	CampaignID int64 `json:"campaign_id,omitempty"`
	// Objective 营销目标
	Objective model.Objective `json:"objective,omitempty"`
	// DailyBudget 日限额
	DailyBudget float64 `json:"daily_budget,omitempty"`
	// BidAmount 出价
	BidAmount float64 `json:"bid_amount,omitempty"`
	// OcpxBidAmount 出价
	OcpxBidAmount float64 `json:"ocpx_bid_amount,omitempty"`
	// BillingType 出价类型
	BillingType BillingType `json:"billing_type,omitempty"`
	// OptimizationGoal 优化目标
	OptimizationGoal int `json:"optimization_goal,omitempty"`
	// ConvertGoal 转化目标
	ConvertGoal []int `json:"convert_goal,omitempty"`
	// AutoMessage 私信模板
	AutoMessage []int `json:"auto_message,omitempty"`
	// DeliveryType 投放类型
	DeliveryType int `json:"delivery_type,omitempty"`
	// DeliverySpeed 投放速度
	DeliverySpeed int `json:"delivery_speed,omitempty"`
	// EffectiveStatus 计划状态
	EffectiveStatus int `json:"effective_status,omitempty"`
	// MidSource 广告类型0常规投放，1已有博文，2聚宝盆代投
	MidSource int `json:"mid_source,omitempty"`
	// Subject 标的对象
	Subject *Subject `json:"subject,omitempty"`
	// Targeting 定向对象
	Targeting *Targeting `json:"targeting,omitempty"`
	// DeliverySchedule 排期对象
	DeliverySchedule []DeliverySchedule `json:"delivery_schedule,omitempty"`
	// DepthConversion 深度转化对象
	DepthConversion *DepthConversion `json:"depth_conversion,omitempty"`
}
