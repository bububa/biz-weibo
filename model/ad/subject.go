package ad

import (
	"github.com/bububa/biz-weibo/enum"
)

// Subject 标的对象
type Subject struct {
	// ID 计划中标的记录唯一索引
	ID uint64 `json:"id,omitempty"`
	// ID 计划Id
	AdID uint64 `json:"ad_id,omitempty"`
	// CustomerID 广告主Id
	CustomerID uint64 `json:"customer_id,omitempty"`
	// ObjType 投放标类型，详见：【附录-投放标的类型】
	ObjType enum.SubjectObjType `json:"obj_type,omitempty"`
	// ObjURL 标的对象地址
	ObjURL string `json:"obj_url,omitempty"`
	// AppType app类型,默认-1，0 Android,1 ios, APP营销目标时必填 ，详见：【附录-APP类型】
	AppType enum.AppType `json:"app_type,omitempty"`
	// AppID app id
	AppID uint64 `json:"app_id,omitempty"`
	// IsAppStoreDirect 是否开启应用商店直达，false关闭，true开启
	IsAppStoreDirect int `json:"is_app_store_direct,omitempty"`
	// AppWM 渠道号
	AppWM string `json:"app_wm,omitempty"`
	// DirectURL 直达链接地址
	DirectURL string `json:"direct_url,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
	// IsDirect 是否开启直达链接， false关闭，true开启
	IsDirect int `json:"is_direct,omitempty"`
	// DirectType 直达类型 0 未使用 1、为手动填写，2、自动生成 ，详见：【附录-直达类型】
	DirectType enum.DirectType `json:"direct_type,omitempty"`
	// ConvertMonitorType 转化监测类型 0默认无， 1. 应用api、2. js、 3. 线索api、 4微信转化api 5 SDK转化 6 小程序监测 ，详见：【附录-转化监测类型】
	ConvertMonitorType enum.ConvertMonitorType `json:"convert_monitor_type,omitempty"`
	// CollectType 归因方式， 详见：【附录-归因方式】
	CollectType enum.CollectType `json:"collect_type,omitempty"`
	// MonitorAPI 监测api地址，
	MonitorAPI string `json:"monitor_api,omitempty"`
	// MiniProgram 小程序
	MiniProgram *MiniProgram `json:"mini_program,omitempty"`
	// CreatedAt 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	UpdatedAt string `json:"updated_at,omitempty"`
	// TopicName 话题名称，详见：【工具-计划定向-话题列表】
	TopicName string `json:"topic_name,omitempty"`
	// ShopProductID 商品对象Id
	ShopProductID string `json:"shop_product_id,omitempty"`
	// URLH5 h5 地址
	URLH5 string `json:"url_h5,omitempty"`
	// UniverseURL ios直达链接
	UniverseURL string `json:"universe_url,omitempty"`
	// MonitorMiniProgramID 小程序监测Id 【示例】 "gh_348f8fb4632a"
	MonitorMiniProgramID string `json:"monitor_mini_program_id,omitempty"`
	// MonitorMiniProgramMacro 小程序监测宏列表 【示例】["oaid","ipv4"]
	MonitorMiniProgramMacro []string `json:"monitor_mini_program_macro,omitempty"`
}

// MiniProgram 小程序
type MiniProgram struct {
	// MiniProgramID 小程序原始Id 【示例】 "gh_348f8fb4632a"
	MiniProgramID string `json:"mini_program_id,omitempty"`
	// MiniProgramAddress 小程序地址 【示例】 "path/test"
	MiniProgramAddress string `json:"mini_program_address,omitempty"`
}
