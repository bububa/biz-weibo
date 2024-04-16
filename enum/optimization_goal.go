package enum

// OptimizationGoal 投放目标，详见：【附录-投放目标】
type OptimizationGoal int

const (
	// OptimizationGoal_CONVERSION 转化
	OptimizationGoal_CONVERSION OptimizationGoal = 86006001
	// OptimizationGoal_IMPRESSION 曝光
	OptimizationGoal_IMPRESSION OptimizationGoal = 86001002
	// OptimizationGoal_CLICK 点击
	OptimizationGoal_CLICK OptimizationGoal = 86002002
	// OptimizationGoal_INTERACT 社交互动
	OptimizationGoal_INTERACT OptimizationGoal = 86002001
)
