package creative

import (
	"encoding/json"

	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
)

// UpdateStatusBatchRequest 创意状态更新 API Request
type UpdateStatusBatchRequest struct {
	// AdvertiserID  您管理的广告主id，为空则为token对应账户uid
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativeIDs 创意id数组
	// 格式：
	// creative_ids=111,222
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// ConfiguredStatus 0：暂停；1：启用；9：结束
	ConfiguredStatus enum.ConfiguredStatus `json:"configured_status,omitempty"`
}

// URL implement Request interface
func (r UpdateStatusBatchRequest) URL() string {
	return "v4/creatives/update_status_batch"
}

// Payload implement Request interface
func (r UpdateStatusBatchRequest) Payload() *model.Payload {
	bs, _ := json.Marshal(r)
	return model.NewPostPayload(bs)
}

// UpdateStatusBatchResponse 编辑广告创意 API Response
type UpdateStatusBatchResponse struct {
	model.BaseResponse
	Data []UpdateResult `json:"data,omitempty"`
}

type UpdateResult struct {
	// Status 	操作成功、失败
	Status bool `json:"status,omitempty"`
	// CreativeID 创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message,omitempty"`
}

func (r UpdateResult) IsError() bool {
	return !r.Status
}

func (r UpdateResult) Error() string {
	return r.ErrorMessage
}
