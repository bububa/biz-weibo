package creative

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/model"
	"github.com/bububa/biz-weibo/util"
)

// UpdateRequest 编辑创意 API Request
type UpdateRequest struct {
	CreativeDetail
	// AdvertiserID  您管理的广告主id，为空则为token对应账户uid
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// URL implement Request interface
func (r UpdateRequest) URL() string {
	return util.StringsJoin("v4/creatives/", strconv.FormatUint(r.CreativeID, 10))
}

// Payload implement Request interface
func (r UpdateRequest) Payload() *model.Payload {
	bs, _ := json.Marshal(r)
	return model.NewPostPayload(bs)
}

// UpdateResponse 编辑广告创意 API Response
type UpdateResponse struct {
	model.BaseResponse
	Data uint64 `json:"data,omitempty"`
}
