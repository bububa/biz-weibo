package account

import (
	"encoding/json"

	"github.com/bububa/biz-weibo/model"
)

// ReverseSpendCapRequest 修改账户次日日限额 API Request
type ReverseSpendCapRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// ReverseSpendCap 日限额账户次日预算值与当日预算值差额不小于100
	ReverseSpendCap float64 `json:"reverse_spend_cap,omitempty"`
}

// URL implement Request interface
func (r ReverseSpendCapRequest) URL() string {
	return "v3/account/reverse_spend_cap"
}

// Payload implement Request interface
func (r ReverseSpendCapRequest) Payload() *model.Payload {
	buf, _ := json.Marshal(r)
	return model.NewPostPayload(buf)
}

// ReverseSpendCapResponse 修改账户次日日限额 API Response
type ReverseSpendCapResponse struct {
	model.BaseResponse
	Account
}
