package ad

// DepthConversion 深度转化对象
type DepthConversion struct {
	// IsDepthConversion 是否开启深度转化，0不开启，1开启
	IsDepthConversion int `json:"is_depth_conversion,omitempty"`
	// DepthConversionStage 深度转换阶段：1 已选深度优化，门槛积累中,2 已选深度优化，门槛已达
	DepthConversionStage int `json:"depth_conversion_stage,omitempty"`
	// DepthConversionGoal 深度转换目标编码：付费 21001、注册 21002、次留21003
	DepthConversionGoal int `json:"depth_conversion_goal,omitempty"`
	// DepthConversionBidAmount 深度转换出价
	DepthConversionBidAmount float64 `json:"depth_conversion_bid_amount,omitempty"`
}
