package enum

// BillingType 出价方式，详见：【附录-出价方式】
type BillingType int

const (
	// BillingType_CPM CPM
	BillingType_CPM BillingType = 4
	// BillingType_CPC CPC
	BillingType_CPC BillingType = 12
	// BillingType_OCPM OCPM
	BillingType_OCPM BillingType = 1
)
