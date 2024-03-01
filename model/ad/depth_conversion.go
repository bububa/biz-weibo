package ad

// DepthConversion 深度转化对象
type DepthConversion struct {
	ID                        int     `json:"id"`                           // 计划中深度转化唯一索引
	AdID                      int     `json:"ad_id"`                        // 计划Id
	IsDepthConversion         int     `json:"is_depth_conversion"`          // 0 关闭 1 开启
	DepthConversionStage      int     `json:"depth_conversion_stage"`       // 深度转换阶段：1 已选深度优化，门槛积累中,2 已选深度优化，门槛已达
	DepthConversionGoalEncode int     `json:"depth_conversion_goal_encode"` // 深度转换目标， 详见：【附录-深度转化目标】
	DepthConversionBidAmount  int     `json:"depth_conversion_bid_amount"`  // 深度转换出价
	DepthConversionRate       float64 `json:"depth_conversion_rate"`        // 深度转化比率 (0,100]
	CreatedAt                 string  `json:"created_at"`                   // 创建时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	UpdatedAt                 string  `json:"updated_at"`                   // 更新时间，时间字符串 YYYY-MM-DD HH:mm:ss 【示例】2023-11-09 00:08:33
	DataVersion               int     `json:"data_version"`                 // 数据版本
}
