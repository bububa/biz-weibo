package ad

import "github.com/bububa/biz-weibo/enum"

// DeliverySchedule 排期对象
type DeliverySchedule struct {
	// ID 计划中投放排期记录唯一索引
	ID uint64 `json:"id,omitempty"`
	// AdID 计划Id
	AdID uint64 `json:"ad_id,omitempty"`
	// StartTime 排期开始时间,普通投放时必填
	StartTime string `json:"start_time,omitempty"`
	// StopTime 排期结束时间,普通投放时必填
	StopTime string `json:"stop_time,omitempty"`
	// DailyDeliveryType 每日投放类型
	DailyDeliveryType enum.DailyDeliveryType `json:"daily_delivery_type,omitempty"`
	// StartTime 每日开始小时 全时段时是0
	DailyStartTime int `json:"daily_start_time,omitempty"`
	// DailyStopTime 每日结束小时 全时段时是24
	DailyStopTime int `json:"daily_stop_time,omitempty"`
	// HourlyPerWeek 每日投放时段，长度为7的数组，每个数组存储投放时段
	HourlyPerWeek string `json:"hourly_per_week,omitempty"`
	// CreatedAt 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】 2020-08-11 12:00:00
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss
	UpdatedAt string `json:"updated_at,omitempty"`
	// ConfiguredStatus 用户设置状态（同步计划状态）
	ConfiguredStatus enum.ConfiguredStatus `json:"configured_status,omitempty"`
	// EffectiveStatus 系统状态（同步计划系统状态）
	EffectiveStatus enum.EffectiveStatus `json:"effective_status,omitempty"`
	// DataVersion 数据版本
	DataVersion int `json:"data_version,omitempty"`
	// CustomerID 广告主Id
	CustomerID uint64 `json:"customer_id,omitempty"`
}
