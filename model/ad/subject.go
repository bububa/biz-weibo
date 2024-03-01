package ad

// Subject 标的对象
type Subject struct {
	ID                 int    `json:"id"`                   // 计划中标的记录唯一索引
	AdID               int    `json:"ad_id"`                // 计划Id
	CustomerID         int    `json:"customer_id"`          // 广告主Id
	ObjType            int    `json:"obj_type"`             // 投放标类型，详见：【附录-投放标的类型】
	ObjURL             string `json:"obj_url"`              // 标的对象地址
	AppType            int    `json:"app_type"`             // app类型,默认-1，0 Android,1 ios, APP营销目标时必填 ，详见：【附录-APP类型】
	AppID              int    `json:"app_id"`               // app id
	IsAppStoreDirect   int    `json:"is_app_store_direct"`  // 是否开启应用商店直达，false关闭，true开启
	AppWM              string `json:"app_wm"`               // 渠道号
	DirectURL          string `json:"direct_url"`           // 直达链接地址
	DownloadURL        string `json:"download_url"`         // 下载链接
	IsDirect           int    `json:"is_direct"`            // 是否开启直达链接， false关闭，true开启
	DirectType         int    `json:"direct_type"`          // 直达类型 0 未使用 1、为手动填写，2、自动生成 ，详见：【附录-直达类型】
	ConvertMonitorType int    `json:"convert_monitor_type"` // 转化监测类型 0默认无， 1. 应用api、2. js、 3. 线索api、 4微信转化api 5 SDK转化 6 小程序监测 ，详见：【附录-转化监测类型】
	CollectType        int    `json:"collect_type"`         // 归因方式， 详见：【附录-归因方式】
	MonitorAPI         string `json:"monitor_api"`          // 监测api地址，
	MiniProgram        struct {
		MiniProgramID      string `json:"mini_program_id"`      // 小程序原始Id 【示例】 "gh_348f8fb4632a"
		MiniProgramAddress string `json:"mini_program_address"` // 小程序地址 【示例】 "path/test"
	} `json:"mini_program"` // 小程序
	CreatedAt               string   `json:"created_at"`                 // 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	UpdatedAt               string   `json:"updated_at"`                 // 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	TopicName               string   `json:"topic_name"`                 // 话题名称，详见：【工具-计划定向-话题列表】
	ShopProductID           string   `json:"shop_product_id"`            // 商品对象Id
	URLH5                   string   `json:"url_h5"`                     // h5 地址
	UniverseURL             string   `json:"universe_url"`               // ios直达链接
	MonitorMiniProgramID    string   `json:"monitor_mini_program_id"`    // 小程序监测Id 【示例】 "gh_348f8fb4632a"
	MonitorMiniProgramMacro []string `json:"monitor_mini_program_macro"` // 小程序监测宏列表 【示例】["oaid","ipv4"]
}
