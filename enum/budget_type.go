package enum

// BudgetType 系列预算类型
type BudgetType int

const (
	// BudgetType_TOTAL 系列总预算
	BudgetType_TOTAL BudgetType = 1
	// BudgetType_DAILY 系列日预算
	BudgetType_DAILY BudgetType = 2
)
