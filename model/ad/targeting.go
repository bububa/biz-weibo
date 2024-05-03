package ad

import "github.com/bububa/biz-weibo/enum"

// Targeting 定向对象
type Targeting struct {
	// ID 计划中定向记录唯一索引
	ID uint64 `json:"id,omitempty"`
	// Name 冗余字段，暂时留空
	Name string `json:"name,omitempty"`
	// AdID 计划Id
	AdID uint64 `json:"ad_id,omitempty"`
	// CustomerID 广告主Id
	CustomerID uint64 `json:"customer_id,omitempty"`
	// ExplicitContentTargeting 内容定向，【示例】0 受众定向 1 内容定向（只支持cpm出价方式) 详见：【附录-内容定向类型】
	ExplicitContentTargeting enum.ExplicitContentTargeting `json:"explicit_content_targeting,omitempty"`
	// AgeMin 最小年龄大于8
	AgeMin int `json:"age_min,omitempty"`
	// AgeMax 最大年龄小于80
	AgeMax int `json:"age_max,omitempty"`
	// Genders 性别，详见：【工具-计划定向-查询计划定向码表】 【示例】 [80,32701,33201]
	Genders int `json:"genders,omitempty"`
	// GeoLocations 省市，详见：【工具-计划定向-查询计划定向码表】 【示例】 [32501,32701,33201]
	GeoLocations []uint64 `json:"geo_locations,omitempty"`
	// RDB 商圈，详见：【工具-计划定向-查询计划定向码表】 【示例】[302010101001,302010101002]
	RBD []uint64 `json:"rbd,omitempty"`
	// District 区县，详见：【工具-计划定向-查询计划定向码表】 【示例】[3020101,3020102]
	District []uint64 `json:"district,omitempty"`
	// AccurateInterests 兴趣，详见: 【工具-计划定向-兴趣关键词】【示例】 [{"suggestion":"教育","count":9822550}]
	AccurateInterests []AccurateInterests `json:"accurate_interests,omitempty"`
	// UserOS 投放平台，详见：【工具-计划定向-查询计划定向码表】 【示例】 [90110202,90110200]
	UserOS []uint64 `json:"user_os,omitempty"`
	// UserDevice 手机品牌，详见：【工具-计划定向-查询计划定向码表】 【示例】[90101000,90102000]
	UserDevice []uint64 `json:"user_device,omitempty"`
	// UserOSVersion 操作系统版本，某些版本以上，详见：【工具-计划定向-查询计划定向码表】
	UserOSVersion uint64 `json:"user_os_version,omitempty"`
	// UserNetwork 用户网络，详见：【工具-计划定向-查询计划定向码表】 【示例】[60110200,60110304,60110305,60110100]
	UserNetwork []uint64 `json:"user_network,omitempty"`
	// AppIDList 应用列表，详见：【工具-计划定向-APP定向列表】 【示例】 [{"app_name":"微信","count":89185427,"app_ids":[1217603,93134]}]
	AppIDList []AppIDList `json:"app_id_list,omitempty"`
	// AppCategoryList 应用分类列表，详见：【工具-计划定向-查询计划定向码表】 【示例】[31001004,31001007]
	AppCategoryList []uint64 `json:"app_category_list"`
	// LoginFrequency 登陆频次 ，详见：【工具-计划定向-查询计划定向码表】 【示例】[210001,210003]
	LoginFrequency []uint64 `json:"login_frequency"`
	// HighPotential 游戏高潜力人群 0 关闭 1 开启
	HighPotential *int `json:"high_potential"`
	// Redirect 重定向开关 0 关闭 1 开启
	Redirect *int `json:"redirect"`
	// Frequency 曝光频次 【示例】{"time_interval":8,"limit":2} 只支持品牌营销目标
	Frequency []TargetFrequency `json:"frequency"`
	// IsTemplate 是否使用定向模版，1：是，0：否
	IsTemplate *int `json:"is_template,omitempty"`
	// TemplateID 定项模板id
	TemplateID uint64 `json:"template_id,omitempty"`
	// ConfiguredStatus 用户设置状态（同步计划状态）
	ConfiguredStatus enum.ConfiguredStatus `json:"configured_status,omitempty"`
	// EffectiveStatus 系统状态（同步计划系统状态）
	EffectiveStatus enum.EffectiveStatus `json:"effective_status,omitempty"`
	// CreatedAt 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】 2020-08-11 12:00:00
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】 2020-08-11 12:00:00
	UpdatedAt string `json:"updated_at,omitempty"`
	// TopicOrientationSwitch 话题定向开关，0 关闭，1 开启
	TopicOrientationSwitch *int `json:"topic_orientation_switch,omitempty"`
	// TopicWords 话题词信息，详见：【工具-计划定向-话题列表】【示例】 [{"topic_id":"1022:231522fd6f3402158f42cb713678772964f222","topic_name":"实时测试"}]
	TopicWords []TopicWord `json:"topic_words,omitempty"`
	// ExceptEnlargeTarget 自动扩量排除支持， network_type只在APP营销目标下可用 详见【附录-自动扩量排除定向码表】
	ExceptEnlargeTarget []enum.EnlargeTarget `json:"except_enlarge_target,omitempty"`
	// IsEnlarge 是否是自动扩量0否 1 开启
	IsEnlarge *int `json:"is_enlarge,omitempty"`
	// FlappingWing 扶翼标识: 0 不支持扶翼 1 支持扶翼 2 扶翼直投
	FlappingWing *int `json:"flapping_wing,omitempty"`
	// EnlargeTarget 自动扩量
	EnlargeTarget []enum.EnlargeTarget `json:"enlarge_target,omitempty"`
	// TagID 标签列表，详见：【工具-计划定向-获取行业词包微博热词】 【示例】 [12,34,5]
	TagID []uint64 `json:"tag_id,omitempty"`
	// NcTagsContent 热词列表,集合nc_tags使用，可以获取热词对应的名称
	NcTagsContent []NcTag `json:"nc_tags_content,omitempty"`
	// IsRta 是否是Rta计划
	IsRta *int `json:"is_rta,omitempty"`
	// IsQualifier 高质量人群，白名单用户使用 0:不是，1:是
	IsQualifier *int `json:"is_qualifier,omitempty"`
	// PromoteFans 是否投放给粉丝 【示例】 0不投放给粉丝1 不限 2投放给我的粉丝
	PromoteFans *int `json:"promote_fans,omitempty"`
	// MultiInterestCategory 兴趣分类 ，详见：【工具-计划定向-兴趣分类】 【示例】[1010201,1010202,1010203]
	MultiInterestCategory []uint64 `json:"multi_interest_category,omitempty"`
	// NcTags 热词的id列表 ，详见：【工具-计划定向-获取行业词包微博热词 【示例】["a","b"]
	NcTags []string `json:"nc_tags,omitempty"`
	// DesignatedAccount 指定账户 ，详见：【工具-计划定向-粉丝定向指定账户】 【示例】 [{"id":123,"followers_count":0}]，followers_count可以默认传0
	DesignatedAccount []DesignatedAccount `json:"designated_account,omitempty"`
	// SeedPack 扩量种子人群，详见：【DMP人群管理-获取人群包列表】 其中 pack_attr_list为：[{"pack_attr_type":"pack_type","pack_attr_value":[1,29]}]
	SeedPack []Audience `json:"seed_pack,omitempty"`
	// AgeGroup 年龄组，详见：【工具-计划定向-查询计划定向码表】 【示例】[1101, 1102]
	AgeGroup []int `json:"age_group,omitempty"`
	// ResidentLocal 常驻地 0不设为常驻地 1 设为常驻地
	ResidentLocal *int `json:"resident_local,omitempty"`
	// LatelyLocal 最近到的 0不设最近到过 1 设为最近到过
	LatelyLocal *int `json:"lately_local,omitempty"`
	// AccountClassify 指定账户分类，详见：【工具-计划定向-账户分类】 【示例】[31001004,31001018]
	AccountClassify []uint64 `json:"account_classify,omitempty"`
	// MobilePrice 手机价格
	MobilePrice []int `json:"mobile_price,omitempty"`
	// ExpConvertAccount 过滤已转化人群，详见：【附录-过滤已转化人群】 【示例】4
	ExpConvertAccount uint64 `json:"exp_convert_account,omitempty"`
	// ExpLowQualityFans 是否过滤低质量粉丝 0 不过滤 1 过滤
	ExpLowQualityFans *int `json:"exp_low_quality_fans,omitempty"`
	// IncludeAucience 定向人群，详见：【DMP人群管理-获取人群包列表】 【示例】[{"pack_id":"13312", "pack_type":"1","origin_pack_id":"1024"}]
	IncludeAucience []Audience `json:"include_aucience,omitempty"`
	// ExcludeAucience 排除人群，详见：【DMP人群管理-获取人群包列表】 【示例】[{"pack_id":"13312", "pack_type":"1","origin_pack_id":"1024"}]
	ExcludeAucience []Audience `json:"exclude_aucience,omitempty"`
	// FilterInstallAccount 安装应用人群 0 不限 1 已安装 2 未安装 ，详见：【附录-安装应用人群】
	FilterInstallAccount *int `json:"filter_install_account,omitempty"`
	// NewAccountType 新用户，详见：【工具-计划定向-查询计划定向码表】【示例】[220001， 220002]
	NewAccountType []uint64 `json:"new_account_type,omitempty"`
	// IsFilterCustomTag 0不过滤 1 过滤
	IsFilterCustomTag *int `json:"is_filter_custom_tag,omitempty"`
	// CustomTag 自定义人群标签，详见：【工具-计划定向-查询自定义人群】 【示例】[2,3,4]
	CustomTag []uint64 `json:"custom_tag,omitempty"`
	// CustomTagRelation 自定义人群关系 1 交集 2 并集
	CustomTagRelation *int `json:"custom_tag_relation,omitempty"`
	// IsOutsideAccount 0 不包含合作账户 1 包含合作账户
	IsOutsideAccount *int `json:"is_outside_account,omitempty"`
	// Price 价格定向 【示例】{"start_price": 1,"end_price": 5,"period": 1, "period_detail": 1}
	Price *TargetingPrice `json:"price,omitempty"`
}

// Target 定向码表
type Target struct {
	ID   int64  `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

// AccurateInterests 兴趣，详见: 【工具-计划定向-兴趣关键词】【示例】 [{"suggestion":"教育","count":9822550}]
type AccurateInterests struct {
	// Suggestion 兴趣关键词
	Suggestion string `json:"suggestion,omitempty"`
	// Count 覆盖月活数
	Count int64 `json:"count,omitempty"`
}

// AppIDList 应用列表，详见：【工具-计划定向-APP定向列表】 【示例】 [{"app_name":"微信","count":89185427,"app_ids":[1217603,93134]}]
type AppIDList struct {
	// AppName app名称
	AppName string `json:"app_name,omitempty"`
	// Count 覆盖月活数
	Count int64 `json:"count,omitempty"`
	// AppIDs app_id列表
	AppIDs []uint64 `json:"app_ids,omitempty"`
}

// TargetFrequency 曝光频次 【示例】{"time_interval":8,"limit":2} 只支持品牌营销目标
type TargetFrequency struct {
	TimeInterval int `json:"time_interval,omitempty"` // 曝光频次中天数设置
	Limit        int `json:"limit,omitempty"`         // 曝光频次中，次数限制
}

// TopicWord 话题词信息，详见：【工具-计划定向-话题列表】【示例】 [{"topic_id":"1022:231522fd6f3402158f42cb713678772964f222","topic_name":"实时测试"}]
type TopicWord struct {
	// TopicID 话题词id
	TopicID string `json:"topic_id,omitempty"`
	// TopicName 话题词
	TopicName string `json:"topic_name,omitempty"`
}

// 热词,集合nc_tags使用，可以获取热词对应的名称
type NcTag struct {
	Name string `json:"name,omitempty"` // 热词对应的名称
	ID   string `json:"id,omitempty"`   // 热词对象对应的id
}

// DesignatedAccount 定账户 ，详见：【工具-计划定向-粉丝定向指定账户】 【示例】 [{"id":123,"followers_count":0}]，followers_count可以默认传0
type DesignatedAccount struct {
	// ID 用户uid
	ID uint64 `json:"id,omitempty"`
	// FollowersCount 可以默认为0
	FollowersCount int64 `json:"followers_count,omitempty"`
}

// Audience 人群
type Audience struct {
	// PackID 包id
	PackID uint64 `json:"pack_id,omitempty"`
	// OriginPackID 原始包id
	OriginPackID uint64 `json:"origin_pack_id,omitempty"`
	// PackType 包类型
	PackType int `json:"pack_type,omitempty"`
	// PackRatio 分成比例
	PackRatio float64 `json:"pack_ratio,omitempty"`
	// IsDivide 是否分成 0 否 1 是
	IsDivide *int `json:"is_divide,omitempty"`
}

// TargetingPrice 价格定向 【示例】{"start_price": 1,"end_price": 5,"period": 1, "period_detail": 1}
type TargetingPrice struct {
	// StartPrice 起始价格
	StartPrice float64 `json:"start_price,omitempty"`
	// EndPrice 结束价格
	EndPrice float64 `json:"end_price,omitempty"`
	// Period 周期，详见：【附录-定向周期】
	Period int `json:"period,omitempty"`
	// PeriodDetail 周期详细，详见：【附录-定向周期】
	PeriodDetail int `json:"period_detail,omitempty"`
}
