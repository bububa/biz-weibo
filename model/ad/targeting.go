package ad

// Targeting 定向对象
type Targeting struct {
	// ExplicitContentTargeting 内容定向
	ExplicitContentTargeting int `json:"explicit_content_targeting,omitempty"`
	// AgeMin 最小年龄大于8
	AgeMin int `json:"age_min,omitempty"`
	// AgeMax 最大年龄小于80
	AgeMax int `json:"age_max,omitempty"`
	// Genders 性别，新用户
	Genders int `json:"genders,omitempty"`
	// GeoLocations 省市
	GeoLocations []int64 `json:"geo_locations,omitempty"`
	// Rbd 商圈
	Rbd []int64 `json:"rbd,omitempty"`
	// District 区县
	District []int64 `json:"district,omitempty"`
	// AccurateInterests 兴趣关键词
	AccurateInterests string `json:"accurate_interests,omitempty"`
	// UserOS 系统版本
	UserOS []int64 `json:"user_os,omitempty"`
	// UserDevice 用户设备
	UserDevice []int64 `json:"user_device,omitempty"`
	// UserOsVersion 操作系统版本，某些版本以上
	UserOsVersion int64 `json:"user_os_version,omitempty"`
	// UserNetwork 用户网络
	UserNetwork []int64 `json:"user_network,omitempty"`
	// AppIDList 应用列表
	AppIDList string `json:"app_id_list,omitempty"`
	// AppCategoryList 应用分类列表
	AppCategoryList []int64 `json:"app_category_list,omitempty"`
	// LoginFrequency 登陆频次
	LoginFrequency []int64 `json:"login_frequency,omitempty"`
	// Frequency 曝光频次
	Frequency string `json:"frequency,omitempty"`
	// FilterInstallAccount 是否过滤已安装人群 0 不过滤 1 过滤
	FilterInstallAccount int `json:"filter_install_account,omitempty"`
	// RbdExtend 地图数据
	RbdExtend string `json:"rbd_extend,omitempty"`
	// ExceptEnlargeTarget 自动扩量二期排除
	ExceptEnlargetTarget []int64 `json:"except_enlarget_target,omitempty"`
	// IsEnlarge 是否是自动扩量0否 1 开启
	IsEnlarge int `json:"is_enlarge,omitempty"`
	// TagID 标签列表
	TagID []int64 `json:"tag_id,omitempty"`
	// NcTags 热词的id列表
	NcTags []int64 `json:"nc_tags,omitempty"`
	// DesignatedAccount 指定账户
	DesignatedAccount string `json:"designated_account,omitempty"`
	// AgeGroup 年龄组
	AgeGroup []int64 `json:"age_group,omitempty"`
	// ResidentLocal 常驻地 0不设为常驻地 1 设为常驻地
	ResidentLocal int `json:"resident_local,omitempty"`
	// LatelyLocal 最近到的 0不设最近到过 1 设为最近到过
	LatelyLocal int `json:"lately_local,omitempty"`
	// AccountClassify 指定账户分类
	AccountClassify []int64 `json:"account_classify,omitempty"`
	// MobilePrice 手机价格
	MobilePrice []int64 `json:"mobile_price,omitempty"`
	// ExpConvertAccount 过滤已转化人群 0 不过滤 1 计划 2 系列 3 账户 4 APP
	ExpConvertAccount int `json:"exp_convert_account,omitempty"`
	// ExpLowQualityFans 是否过滤低质量粉丝 0 不过滤 1 过滤
	ExpLowQualityFans int `json:"exp_low_quality_fans,omitempty"`
	// IncludeAudience 定向人群
	IncludeAudience string `json:"include_audience,omitempty"`
	// ExcludeAudience 排除人群
	ExcludeAudience string `json:"exclude_audience,omitempty"`
	// PromoteFans 是否投放给粉丝
	PromoteFans int64 `json:"promote_fans,omitempty"`
	// MultiInterestCategory 兴趣分类
	MultiInterestCategory []int64 `json:"multi_interest_category,omitempty"`
	// NewAccountType 新用户
	NewAccountType []int64 `json:"new_account_type,omitempty"`
	// TemplateID 定项模板id，此字段会覆盖target, 参考 定项模板records[target_template_id]字段
	TemplateID string `json:"template_id,omitempty"`
}

// Target 定向码表
type Target struct {
	ID   int64  `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}
