package campaign

import (
	"encoding/json"
	"fmt"

	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
)

// UpdateRequest 更新系列预算或名称
type UpdateRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ID 系列ID
	ID uint64 `json:"id,omitempty"`
	// Name 系列名称 最长50个字符
	Name string `json:"name,omitempty"`
	// BudgetType 预算类型
	BudgetType enum.BudgetType `json:"budget_type,omitempty"`
	// LifetimeBudget 系列总预算，0表示不限，系列总预算修改必须符合以下限制：最低系列总预算为150元，最高系列总预算9,999,999.99元，单次修改幅度不得低于100元，可调低，调低立即生效，调低限制：不得低于当前系列实时总消耗*120%，不得低于当前系列实时总消耗+150元
	LifetimeBudget float64 `json:"lifetime_budget,omitempty"`
	// DailyBudget 每日预算，单位：元。0表示不限
	DailyBudget float64 `json:"daily_budget,omitempty"`
}

// URL implement Request interface
func (r UpdateRequest) URL() string {
	return fmt.Sprintf("v3/campaigns/%d", r.ID)
}

// Payload implement Request interface
func (r UpdateRequest) Payload() *model.Payload {
	buf, _ := json.Marshal(r)
	return model.NewPostPayload(buf)
}

// UpdateResponse 更新系列预算或名称 API Response
type UpdateResponse struct {
	model.BaseResponse
	Campaign
}
