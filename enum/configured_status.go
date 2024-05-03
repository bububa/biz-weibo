package enum

// 用户设置状态，详见：【附录-计划用户设置状态】
type ConfiguredStatus int

const (
	// ConfiguredStatus_PAUSE 暂停
	ConfiguredStatus_PAUSE ConfiguredStatus = 0
	// ConfiguredStatus_ACTIVE 开启
	ConfiguredStatus_ACTIVE ConfiguredStatus = 1
	// ConfiguredStatus_DELETE 删除
	ConfiguredStatus_DELETE ConfiguredStatus = 9
)
