package enum

// CollectType 归因方式， 详见：【附录-归因方式】
type CollectType int

const (
	// CollectType_DEFAULT 默认值 无
	CollectType_DEFAULT CollectType = 0
	// CollectType_CLICK 点击
	CollectType_CLICK CollectType = 10000
	// CollectType_INTERACT 互动
	CollectType_INTERACT CollectType = 11000
)
