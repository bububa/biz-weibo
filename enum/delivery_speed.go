package enum

// DeliverySpeed 投放速度， 详见：【附录-投放速度】 【示例】1
type DeliverySpeed int

const (
	// DeliverySpeed_FAST 优先跑量
	DeliverySpeed_FAST DeliverySpeed = 1
	// DeliverySpeed_CONSTANT 匀速投放
	DeliverySpeed_CONSTANT DeliverySpeed = 2
)
