package account

import (
	"github.com/bububa/biz-weibo/model"
	"strconv"
)

type ManageListRequest struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"` // 您管理的广告主id，为空则为token对应账户 （不使用请勿下发）
	Page         int    `json:"page,omitempty"`          // 当前页码
	PageSize     int    `json:"page_size,omitempty"`     // 每页条数，默认为20
	Status       int    `json:"status,omitempty"`        // 状态
}

func (r ManageListRequest) URL() string {
	return "v3/account/manage-list"
}

func (r ManageListRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	if r.Page > 0 {
		p.AddValue("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		p.AddValue("page_size", strconv.Itoa(r.PageSize))
	}
	p.AddValue("status", strconv.Itoa(r.Status))
	return p
}

type ManageListResponse struct {
	model.BaseResponse
	PageTotal int      `json:"page_total"`
	PageSize  int      `json:"page_size"`
	Page      int      `json:"page"`
	List      []Manage `json:"list"`
}

type Manage struct {
	EffectiveStatus  int    `json:"effective_status"`
	Role             int    `json:"role"`
	ChildId          int64  `json:"child_id"`
	UpdatedAt        string `json:"updated_at"`
	MainName         string `json:"main_name"`
	MainId           uint64 `json:"main_id"`
	ConfiguredStatus int    `json:"configured_status"`
	CreatedAt        string `json:"created_at"`
	Id               int    `json:"id"`
}
