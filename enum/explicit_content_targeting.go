package enum

// ExplicitContentTargeting 内容定向，【示例】0 受众定向 1 内容定向（只支持cpm出价方式) 详见：【附录-内容定向类型】
type ExplicitContentTargeting int

const (
	// ExplicitContentTargeting_AUDIENCE 受众定向
	ExplicitContentTargeting_AUDIENCE ExplicitContentTargeting = 0
	// ExplicitContentTargeting_CONTENT 内容定向
	ExplicitContentTargeting_CONTENT ExplicitContentTargeting = 1
)
