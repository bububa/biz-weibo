package ad

// DeliverySchedule 排期对象
type DeliverySchedule struct {
	// StartTime 排期开始时间,普通投放时必填
	StartTime string `json:'start_time,omitempty'`
	// StopTime 排期结束时间,普通投放时必填
	StopTime string `json:"stop_time,omitempty"`
	// DailyDeliveryType 每日投放类型
	DailyDeliveryType int `json:"daily_delivery_type,omitempty"`
	// DailyStartTime 每日开始小时 全时段时是0
	DailyStartTime int `json:"daily_start_time,omitempty"`
	// DailyStopTime 每日结束小时 全时段时是24
	DailyStopTime int `json:"daily_stop_time,omitempty"`
	// HourlyPerWeek 每日投放时段，长度为7的数组，每个数组存储投放时段
	HourlyPerWeek string `json:"hourly_per_week,omitempty"`
}
