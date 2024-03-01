package account

import (
	"encoding/json"

	"github.com/bububa/biz-weibo/model"
)

// BudgetRequest 修改账户日限额 API Request
type BudgetRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// SpendCap 日限额
	SpendCap float64 `json:"spend_cap,omitempty"`
}

// URL implement Request interface
func (r BudgetRequest) URL() string {
	return "v3/account/budget"
}

// Payload implement Request interface
func (r BudgetRequest) Payload() *model.Payload {
	buf, _ := json.Marshal(r)
	return model.NewPostPayload(buf)
}

// BudgetResponse 修改账户日限额 API Response
type BudgetResponse struct {
	model.BaseResponse
	Account
}
