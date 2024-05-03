package enum

// ConvertMonitorType 转化监测类型 0默认无， 1. 应用api、2. js、 3. 线索api、 4微信转化api 5 SDK转化 6 小程序监测 ，详见：【附录-转化监测类型】
type ConvertMonitorType int

const (
	// ConvertMonitorType_DEFAULT 默认值 无
	ConvertMonitorType_DEFAULT ConvertMonitorType = 0
	// ConvertMonitorType_APP_API 应用api
	ConvertMonitorType_APP_API ConvertMonitorType = 1
	// ConvertMonitorType_JS js
	ConvertMonitorType_JS ConvertMonitorType = 2
	// ConvertMonitorType_CLUE_API 线索api
	ConvertMonitorType_CLUE_API ConvertMonitorType = 3
	// ConvertMonitorType_WECHAT_CONVERSION_API 微信转化api
	ConvertMonitorType_WECHAT_CONVERSION_API ConvertMonitorType = 4
	// ConvertMonitorType_SDK_CONVERSION sdk转化
	ConvertMonitorType_SDK_CONVERSION ConvertMonitorType = 5
	// ConvertMonitorType_MINIAPP_MONITOR 小程序监测
	ConvertMonitorType_MINIAPP_MONITOR ConvertMonitorType = 6
)
