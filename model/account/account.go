package account

// Account 账号
type Account struct {
	// ID 广告账号id
	ID uint64 `json:"id,omitempty"`
	// CustomerID 广告客户微博id
	CustomerID uint64 `json:"customer_id,omitempty"`
	// Name 广告客户名称
	Name string `json:"name,omitempty"`
	// ProfileImage 广告账户头像
	ProfileImage string `json:"profile_image,omitempty"`
	// DisableReason 禁用原因
	DisableReason string `json:"disable_reason,omitempty"`
	// SpendCap 日限额
	SpendCap string `json:"spend_cap,omitempty"`
	// ReverseSpendCap 限额调低后第二天才能生效的日限额 为0时忽略
	ReverseSpendCap string `json:"reverse_spend_cap,omitempty"`
	// Version 当前超粉版本号
	Version int `json:"version,omitempty"`
	// EnrichType 0，幽灵账户；1，自助客户；2，SME客户；3，KA广告主
	EnrichType int `json:"enrich_type,omitempty"`
	// Block 阻止登录
	Block int `json:"block,omitempty"`
	// CreatedAt 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// MigratedAt
	MigratedAt string `json:"migrated_at,omitempty"`
	// ConfiguredStatus
	ConfiguredStatus int `json:"configured_status,omitempty"`
	// EffectiveStatus 系统状态
	EffectiveStatus int `json:"effective_status,omitempty"`
}
