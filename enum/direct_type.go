package enum

// DirectType 直达类型 0 未使用 1、为手动填写，2、自动生成 ，详见：【附录-直达类型】
type DirectType int

const (
	// DirectType_UNUSED 未使用
	DirectType_UNUSED DirectType = 0
	// DirectType_MANUAL 手动填写
	DirectType_MANUAL DirectType = 1
	// DirectType_AUTO 自动生成
	DirectType_AUTO DirectType = 2
)
