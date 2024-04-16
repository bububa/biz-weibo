package ad

import (
	"github.com/bububa/biz-weibo/enum"
)

// DepthConversion 深度转化对象
type DepthConversion struct {
	// ID 计划中深度转化唯一索引
	ID uint64 `json:"id,omitempty"`
	// AdID  计划Id
	AdID uint64 `json:"ad_id,omitempty"`
	// IsDepthConversion 0 关闭 1 开启
	IsDepthConversion int `json:"is_depth_conversion"`
	// DepthConversionStage 深度转换阶段：1 已选深度优化，门槛积累中,2 已选深度优化，门槛已达
	DepthConversionStage int `json:"depth_conversion_stage"`
	// DepthConversionGoalEncode 深度转换目标， 详见：【附录-深度转化目标】
	DepthConversionGoalEncode enum.DepthConversionGoalEncode `json:"depth_conversion_goal_encode"`
	// DepthConversionBidAmount 度转换出价
	DepthConversionBidAmount float64 `json:"depth_conversion_bid_amount"`
	// DepthConversionRate 深度转化比率 (0,100]
	DepthConversionRate float64 `json:"depth_conversion_rate"`
	// CreatedAt 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	UpdatedAt string `json:"updated_at"`
	// DataVersion 数据版本
	DataVersion int `json:"data_version"`
}
