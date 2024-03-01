package ad

// Targeting 定向对象
type Targeting struct {
	ID                       int    `json:"id"`                         // 计划中定向记录唯一索引
	Name                     string `json:"name"`                       // 冗余字段，暂时留空
	AdID                     int    `json:"ad_id"`                      // 计划Id
	CustomerID               int    `json:"customer_id"`                // 广告主Id
	ExplicitContentTargeting int    `json:"explicit_content_targeting"` // 内容定向，【示例】0 受众定向 1 内容定向（只支持cpm出价方式) 详见：【附录-内容定向类型】
	AgeMin                   int    `json:"age_min"`                    // 最小年龄大于8
	AgeMax                   int    `json:"age_max"`                    // 最大年龄小于80
	Genders                  int    `json:"genders"`                    // 性别，详见：【工具-计划定向-查询计划定向码表】 【示例】 [80,32701,33201]
	GeoLocations             []int  `json:"geo_locations"`              // 省市，详见：【工具-计划定向-查询计划定向码表】 【示例】 [32501,32701,33201]
	RBD                      []int  `json:"rbd"`                        // 商圈，详见：【工具-计划定向-查询计划定向码表】 【示例】[302010101001,302010101002]
	District                 []int  `json:"district"`                   // 区县，详见：【工具-计划定向-查询计划定向码表】 【示例】[3020101,3020102]
	AccurateInterests        []struct {
		Suggestion string `json:"suggestion"` // 兴趣关键词
		Count      int    `json:"count"`      // 覆盖月活数
	} `json:"accurate_interests"` // 兴趣，详见: 【工具-计划定向-兴趣关键词】【示例】 [{"suggestion":"教育","count":9822550}]
	UserOS        []int `json:"user_os"`         // 投放平台，详见：【工具-计划定向-查询计划定向码表】 【示例】 [90110202,90110200]
	UserDevice    []int `json:"user_device"`     // 手机品牌，详见：【工具-计划定向-查询计划定向码表】 【示例】[90101000,90102000]
	UserOSVersion int   `json:"user_os_version"` // 操作系统版本，某些版本以上，详见：【工具-计划定向-查询计划定向码表】
	UserNetwork   []int `json:"user_network"`    // 用户网络，详见：【工具-计划定向-查询计划定向码表】 【示例】[60110200,60110304,60110305,60110100]
	AppIDList     []struct {
		AppName string `json:"app_name"` // app名称
		Count   int    `json:"count"`    // 覆盖月活数
		AppIDs  []int  `json:"app_ids"`  // app_id列表
	} `json:"app_id_list"` // 应用列表，详见：【工具-计划定向-APP定向列表】 【示例】 [{"app_name":"微信","count":89185427,"app_ids":[1217603,93134]}]
	AppCategoryList []int `json:"app_category_list"` // 应用分类列表，详见：【工具-计划定向-查询计划定向码表】 【示例】[31001004,31001007]
	LoginFrequency  []int `json:"login_frequency"`   // 登陆频次 ，详见：【工具-计划定向-查询计划定向码表】 【示例】[210001,210003]
	HighPotential   int   `json:"high_potential"`    // 游戏高潜力人群 0 关闭 1 开启
	Redirect        int   `json:"redirect"`          // 重定向开关 0 关闭 1 开启
	Frequency       struct {
		TimeInterval int `json:"time_interval"` // 曝光频次中天数设置
		Limit        int `json:"limit"`         // 曝光频次中，次数限制
	} `json:"frequency"` // 曝光频次 【示例】{"time_interval":8,"limit":2} 只支持品牌营销目标
	IsTemplate             int    `json:"is_template"`              // 是否使用定向模版，1：是，0：否
	TemplateID             int    `json:"template_id"`              // 定项模板id
	ConfiguredStatus       int    `json:"configured_status"`        // 用户设置状态（同步计划状态）
	EffectiveStatus        int    `json:"effective_status"`         // 系统状态（同步计划系统状态）
	CreatedAt              string `json:"created_at"`               // 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】 2020-08-11 12:00:00
	UpdatedAt              string `json:"updated_at"`               // 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】 2020-08-11 12:00:00
	TopicOrientationSwitch int    `json:"topic_orientation_switch"` // 话题定向开关，0 关闭，1 开启
	TopicWords             []struct {
		TopicID   string `json:"topic_id"`   // 话题词id
		TopicName string `json:"topic_name"` // 话题词
	} `json:"topic_words"` // 话题词信息，详见：【工具-计划定向-话题列表】【示例】 [{"topic_id":"1022:231522fd6f3402158f42cb713678772964f222","topic_name":"实时测试"}]
	ExceptEnlargeTarget []string `json:"except_enlarge_target"` // 自动扩量排除支持， network_type只在APP营销目标下可用 详见【附录-自动扩量排除定向码表】
	IsEnlarge           int      `json:"is_enlarge"`            // 是否是自动扩量0否 1 开启
	FlappingWing        int      `json:"flapping_wing"`         // 扶翼标识: 0 不支持扶翼 1 支持扶翼 2 扶翼直投
	EnlargeTarget       []string `json:"enlarge_target"`        // 自动扩量
	TagID               []int    `json:"tag_id"`                // 标签列表，详见：【工具-计划定向-获取行业词包微博热词】 【示例】 [12,34,5]
	NcTagsContent       []struct {
		Name string `json:"name"` // 热词对应的名称
		ID   string `json:"id"`   // 热词对象对应的id
	} `json:"nc_tags_content"` // 热词列表,集合nc_tags使用，可以获取热词对应的名称
	NcTags            []string `json:"nc_tags"` // 热词的id列表 ，详见：【工具-计划定向-获取行业词包微博热词 【示例】["a","b"]
	DesignatedAccount []struct {
		ID             int `json:"id"`              // 用户uid
		FollowersCount int `json:"followers_count"` // 可以默认为0
	} `json:"designated_account"` // 指定账户 ，详见：【工具-计划定向-粉丝定向指定账户】 【示例】 [{"id":123,"followers_count":0}]，followers_count可以默认传0
	SeedPack []struct {
		PackID       int `json:"pack_id"`        // 包id
		OriginPackID int `json:"origin_pack_id"` // 原始包id
		PackType     int `json:"pack_type"`      // 包类型
		IsDivide     int `json:"is_divide"`      // 是否分成 0 否 1 是
	} `json:"seed_pack"` // 扩量种子人群，详见：【DMP人群管理-获取人群包列表】 其中 pack_attr_list为：[{"pack_attr_type":"pack_type","pack_attr_value":[1,29]}]
	AgeGroup          []int `json:"age_group"`            // 年龄组，详见：【工具-计划定向-查询计划定向码表】 【示例】[1101, 1102]
	ResidentLocal     int   `json:"resident_local"`       // 常驻地 0不设为常驻地 1 设为常驻地
	LatelyLocal       int   `json:"lately_local"`         // 最近到的 0不设最近到过 1 设为最近到过
	AccountClassify   []int `json:"account_classify"`     // 指定账户分类，详见：【工具-计划定向-账户分类】 【示例】[31001004,31001018]
	MobilePrice       []int `json:"mobile_price"`         // 手机价格
	IsFilterCustomTag int   `json:"is_filter_custom_tag"` // 0不过滤 1 过滤
	CustomTag         []int `json:"custom_tag"`           // 自定义人群标签，详见：【工具-计划定向-查询自定义人群】 【示例】[2,3,4]
	CustomTagRelation int   `json:"custom_tag_relation"`  // 自定义人群关系 1 交集 2 并集
	IsOutsideAccount  int   `json:"is_outside_account"`   // 0 不包含合作账户 1 包含合作账户
	Price             struct {
		StartPrice   float64 `json:"start_price"`   // 起始价格
		EndPrice     float64 `json:"end_price"`     // 结束价格
		Period       int     `json:"period"`        // 周期，详见：【附录-定向周期】
		PeriodDetail int     `json:"period_detail"` // 周期详细，详见：【附录-定向周期】
	} `json:"price"` // 价格定向 【示例】{"start_price": 1,"end_price": 5,"period": 1, "period_detail": 1}
}

// Target 定向码表
type Target struct {
	ID   int64  `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}
