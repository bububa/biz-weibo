package ad

// Ad 计划对象
type Ad struct {
	ID                       int             `json:"id"`                           // 计划id
	Name                     string          `json:"name"`                         // 计划名称
	CampaignID               int             `json:"campaign_id"`                  // 所属广告系列id
	Objective                int             `json:"objective"`                    // 营销目标，详见：【附录-营销目标】
	DailyBudget              string          `json:"daily_budget"`                 // 计划日预算
	ReverseDailyBudget       string          `json:"reverse_daily_budget"`         // 计划次日预算
	ReverseOCPXBidAmountDate string          `json:"reverse_ocpx_bid_amount_date"` // 待生效OCPX出价生效日期，时间字符串，格式为YYYY-MM-DD 【示例】2023-10-10
	DeliveryType             int             `json:"delivery_type"`                // 投放类型， 详见：【附录-投放类型】 【示例】1
	DeliverySpeed            int             `json:"delivery_speed"`               // 投放速度， 详见：【附录-投放速度】 【示例】1
	BidAmount                string          `json:"bid_amount"`                   // 出价 【示例】200.01
	OCPXBidAmount            string          `json:"ocpx_bid_amount"`              // 出价 【示例】100.01
	BillingType              int             `json:"billing_type"`                 // 出价方式，详见：【附录-出价方式】
	OptimizationGoal         int             `json:"optimization_goal"`            // 投放目标，详见：【附录-投放目标】
	TransformGoal            string          `json:"transform_goal"`               // 转化目标，详见：【附录-转化目标】
	ConfiguredStatus         int             `json:"configured_status"`            // 用户设置状态，详见：【附录-计划用户设置状态】
	EffectiveStatus          int             `json:"effective_status"`             // 计划状态（超粉系统)，详见：【附录-计划状态】
	DeliverScope             int             `json:"deliver_scope"`                // 投放范围（历史原因保留字段，1：微博站内、2：优选流量、3：新浪站内），建议使用详情接口获取投放范围，详见：获取广告计划详情，【附录-投放范围】
	IsRecycled               int             `json:"is_recycled"`                  // 是否删除 1：已删除，0：未删除
	RecycleDatetime          string          `json:"recycle_datetime"`             // 删除时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-11 00:08:33
	OAuth2ID                 int             `json:"oauth2_id"`                    // 计划来源，0来源超粉系统
	CreatedAt                string          `json:"created_at"`                   // 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	UpdatedAt                string          `json:"updated_at"`                   // 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-10 00:08:33
	Version                  int             `json:"version"`                      // 计划版本
	IsGesture                int             `json:"is_gesture"`                   // 是否是手势互动计划 0 未标记 1 手势互动计划 2 无手势互动计划
	BrandPartnerSwitch       int             `json:"brand_partner_switch"`         // 星品联动标识 0:未标记 1:开启2:关闭
	CanAddCreative           bool            `json:"can_add_creative"`             // 是否可添加创意 true:可以添加 false:不可以添加【示例】true
	IsEdit                   bool            `json:"is_edit"`                      // 计划是否可以被编辑，true：可以被编辑 false:不可以被编辑【示例】true
	AutoMessage              []AutoMessage   `json:"auto_message"`                 // 自动私信，【示例】[{"code":14000005,"id":13}]
	DepthConversion          DepthConversion `json:"depth_conversion"`             // 深度转化对象，当未开启深度转化时为null
	Subject                  Subject         // 标的对象
	Targeting                Targeting       `json:"targeting"` // 定向对象
	Brand                    Brand           `json:"brand"`     // 品牌对象
}

type AutoMessage struct {
	Code int `json:"code"` // 自动私信类型 ，详见：【附录-自动私信类型】
	ID   int `json:"id"`   // 私信模板id
}

type Brand struct {
	ID              int    `json:"id"`                // 计划中品牌对象唯一索引
	AdID            int    `json:"ad_id"`             // 计划Id
	IsBan           int    `json:"is_ban"`            // 限制投放设置 0：不限制 1：限制
	ExcludeBannedID string `json:"exclude_banned_id"` // 排除品牌 id 字符串，逗号隔开，排除所有品牌，为空字符串，有特定品牌，逗号隔开 【示例】 "1,2"
	ExcludeCoopID   string `json:"exclude_coop_id"`   // 排除合作商 id 字符串，逗号隔开，排除所有合作商，为空字符串，有特定合作商，逗号隔开 【示例】 "1,2"
	CreatedAt       string `json:"created_at"`        // 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	UpdatedAt       string `json:"updated_at"`        // 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
}
