package enum

// DailyDeliveryType 投放时段
type DailyDeliveryType int

const (
	// DailyDeliveryType_ALL_TIME 全时段
	DailyDeliveryType_ALL_TIME DailyDeliveryType = 1
	// DailyDeliveryType_WEEKLY 每周投放时段
	DailyDeliveryType_WEEKLY DailyDeliveryType = 3
)
