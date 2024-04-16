package ad

import (
	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
	"github.com/bububa/biz-weibo/model/campaign"
)

// Ad 计划对象
type Ad struct {
	// ID 计划id
	ID uint64 `json:"id,omitempty"`
	// Name 计划名称
	Name string `json:"name,omitempty"`
	// Campaign 所属广告系列id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CustomerID 广告主Id
	CustomerID uint64 `json:"customer_id,omitempty"`
	// Objective 营销目标，详见：【附录-营销目标】
	Objective enum.Objective `json:"objective,omitempty"`
	// DailyBudget 计划日预算
	DailyBudget model.Float64 `json:"daily_budget,omitempty"`
	// ReverseDailyBudget 计划次日预算
	ReverseDailyBudget model.Float64 `json:"reverse_daily_budget,omitempty"`
	// 待生效OCPX出价生效日期，时间字符串，格式为YYYY-MM-DD 【示例】2023-10-10
	ReverseOCPXBidAmountDate string `json:"reverse_ocpx_bid_amount_date,omitempty"`
	// DeliveryType  投放类型， 详见：【附录-投放类型】 【示例】1
	DeliveryType enum.DeliveryType `json:"delivery_type,omitempty"`
	// DeliverySpeed 投放速度， 详见：【附录-投放速度】 【示例】1
	DeliverySpeed enum.DeliverySpeed `json:"delivery_speed,omitempty"`
	// BidAmount 出价 【示例】200.01
	BidAmount model.Float64 `json:"bid_amount,omitempty"`
	// OcpxBidAmount 出价 【示例】100.01
	OcpxBidAmount model.Float64 `json:"ocpx_bid_amount,omitempty"`
	// BillingType 出价方式，详见：【附录-出价方式】
	BillingType enum.BillingType `json:"billing_type,omitempty"`
	// OptimizationGoal  投放目标，详见：【附录-投放目标】
	OptimizationGoal enum.OptimizationGoal `json:"optimization_goal,omitempty"`
	// TransformGoal 转化目标，详见：【附录-转化目标】
	TransformGoal enum.TransformGoal `json:"transform_goal,omitempty"`
	// 用户设置状态，详见：【附录-计划用户设置状态】
	ConfiguredStatus enum.ConfiguredStatus `json:"configured_status,omitempty"`
	// EffectiveStatus 计划状态（超粉系统)，详见：【附录-计划状态】
	EffectiveStatus enum.EffectiveStatus `json:"effective_status,omitempty"`
	// DeliverScope 投放范围（历史原因保留字段，1：微博站内、2：优选流量、3：新浪站内），建议使用详情接口获取投放范围，详见：获取广告计划详情，【附录-投放范围】
	DeliverScope enum.DeliverScope `json:"deliver_scope,omitempty"`
	// IsRecycled 是否删除 1：已删除，0：未删除
	IsRecycled int `json:"is_recycled,omitempty"`
	// RecycleDatetime 删除时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-11 00:08:33
	RecycleDatetime string `json:"recycle_datetime,omitempty"`
	// OAuth2ID 计划来源，0来源超粉系统
	OAuth2ID int `json:"oauth2_id,omitempty"`
	// CreatedAt 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-10 00:08:33
	UpdatedAt string `json:"updated_at,omitempty"`
	// Version 计划版本
	Version int `json:"version,omitempty"`
	// GuaranteeStatus 保障标记，-1无保障 1成本保障中 2不满足保障 3保障核实中 4 已赔付
	GuaranteeStatus int `json:"guarantee_status,omitempty"`
	// IsGesture 是否是手势互动计划 0 未标记 1 手势互动计划 2 无手势互动计划
	IsGesture int `json:"is_gesture,omitempty"`
	// BrandPartnerSwitch 星品联动标识 0:未标记 1:开启2:关闭
	BrandPartnerSwitch int `json:"brand_partner_switch,omitempty"`
	// CanAddCreative 是否可添加创意 true:可以添加 false:不可以添加【示例】true
	CanAddCreative bool `json:"can_add_creative,omitempty"`
	// IsEdit 计划是否可以被编辑，true：可以被编辑 false:不可以被编辑【示例】true
	IsEdit bool `json:"is_edit,omitempty"`
	// AutoMessage 自动私信，【示例】[{"code":14000005,"id":13}]
	AutoMessage []AutoMessage `json:"auto_message,omitempty"`
	// DepthConversion 深度转化对象，当未开启深度转化时为null
	DepthConversion *DepthConversion `json:"depth_conversion,omitempty"`
	// Subject 标的对象
	Subject *Subject `json:"subject,omitempty"`
	// Targeting 定向对象
	Targeting *Targeting `json:"targeting,omitempty"`
	// DeliverySchedule 排期对象
	DeliverySchedule *DeliverySchedule `json:"delivery_schedule,omitempty"`
	// AdExtendBid 一键起量对象
	AdExtendBid *AdExtendBid `json:"ad_extend_bid,omitempty"`
	// Campaign 系列对象
	Campaign *campaign.Campaign `json:"campaign,omitempty"`
}

type AutoMessage struct {
	// Code 自动私信类型 ，详见：【附录-自动私信类型】
	Code enum.AutoMessageCode `json:"code,omitempty"`
	ID   uint64               `json:"id,omitempty"` // 私信模板id
}

type Brand struct {
	ID              uint64 `json:"id,omitempty"`                // 计划中品牌对象唯一索引
	AdID            uint64 `json:"ad_id,omitempty"`             // 计划Id
	IsBan           int    `json:"is_ban,omitempty"`            // 限制投放设置 0：不限制 1：限制
	ExcludeBannedID string `json:"exclude_banned_id,omitempty"` // 排除品牌 id 字符串，逗号隔开，排除所有品牌，为空字符串，有特定品牌，逗号隔开 【示例】 "1,2"
	ExcludeCoopID   string `json:"exclude_coop_id,omitempty"`   // 排除合作商 id 字符串，逗号隔开，排除所有合作商，为空字符串，有特定合作商，逗号隔开 【示例】 "1,2"
	CreatedAt       string `json:"created_at,omitempty"`        // 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	UpdatedAt       string `json:"updated_at,omitempty"`        // 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
}
