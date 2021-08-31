package account

import (
	"encoding/json"

	"github.com/bububa/biz-weibo/model"
)

// DeleteReverseBudgetRequest 删除账户次日日限额 API Request
type DeleteReverseBudgetRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
}

// URL implement Request interface
func (r DeleteReverseBudgetRequest) URL() string {
	return "v3/account/delete_reverse_budget"
}

// Payload implement Request interface
func (r DeleteReverseBudgetRequest) Payload() *model.Payload {
	buf, _ := json.Marshal(r)
	return model.NewPostPayload(buf)
}

// DeleteReverseBudgetResponse 删除账户次日日限额 API Response
type DeleteReverseBudgetResponse struct {
	model.BaseResponse
	Account
}
