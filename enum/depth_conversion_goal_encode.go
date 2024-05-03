package enum

// DepthConversionGoalEncode 深度转换目标， 详见：【附录-深度转化目标】
type DepthConversionGoalEncode int

const (
	// DepthConversionGoalEncode_PAY 付费
	DepthConversionGoalEncode_PAY DepthConversionGoalEncode = 21001
	// DepthConversionGoalEncode_RETENTATION 次留
	DepthConversionGoalEncode_RETENTATION DepthConversionGoalEncode = 21002
	// DepthConversionGoalEncode_ORDER 用户下单
	DepthConversionGoalEncode_ORDER DepthConversionGoalEncode = 21003
	// DepthConversionGoalEncode_EFFECTIVE_CONSULT 有效咨询
	DepthConversionGoalEncode_EFFECTIVE_CONSULT DepthConversionGoalEncode = 21006
	// DepthConversionGoalEncode_EFFECVIVE_CLUE 有效线索
	DepthConversionGoalEncode_EFFECVIVE_CLUE DepthConversionGoalEncode = 21007
)
