package enum

// DeliveryType  投放类型， 详见：【附录-投放类型】 【示例】1
type DeliveryType int

const (
	// DeliveryType_LONGTERM 长期投放
	DeliveryType_LONGTERM DeliveryType = 1
	// DeliveryType_NORMAL 普通投放
	DeliveryType_NORMAL DeliveryType = 2
)
