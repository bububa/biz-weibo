package campaign

import (
	"encoding/json"

	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
)

// UpdateStatusRequest 更新系列状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDs 系列id，数组型，支持多个
	IDs []uint64 `json:"ids,omitempty"`
	// ConfiguredStatus 目标状态
	ConfiguredStatus enum.ConfiguredStatus `json:"configured_status"`
}

// URL implement Request interface
func (r UpdateStatusRequest) URL() string {
	return "v3/campaigns/update-status"
}

// Payload implement Request interface
func (r UpdateStatusRequest) Payload() *model.Payload {
	buf, _ := json.Marshal(r)
	return model.NewPostPayload(buf)
}

// UpdateStatusResponse 更新系列状态 API Response
type UpdateStatusResponse []UpdateStatusResult

// IsError implement Response interface
func (r UpdateStatusResponse) IsError() bool {
	return false
}

// Error implement Response interface
func (r UpdateStatusResponse) Error() string {
	return ""
}

// UpdateStatus 系列更新状态结果
type UpdateStatusResult struct {
	CampaignID   uint64 `json:"campaign_id,omitempty"`
	Status       bool   `json:"status,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
