package enum

// AppType app类型,默认-1，0 Android,1 ios, APP营销目标时必填 ，详见：【附录-APP类型】
type AppType int

const (
	// AppType_DEFAULT 默认值
	AppType_DEFAULT AppType = -1
	// AppType_ANDROID Android
	AppType_ANDROID AppType = 0
	// AppType_IOS iOS
	AppType_IOS AppType = 1
)
