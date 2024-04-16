package campaign

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
)

// CampaignsRequest 广告系列列表 API Request
type CampaignsRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 按照广告系列名称关键字过滤，若无需求请勿传值
	Name string `json:"name,omitempty"`
	// EffectiveStatus 按照广告系列状态过滤，若无需求请勿传值
	EffectiveStatus []enum.CampaignEffectiveStatus `json:"effective_status,omitempty"`
	// Objectives 按照营销目标过滤，若无需求请勿传值
	Objectives []enum.Objective `json:"objectives,omitempty"`
	// IDs id筛选
	IDs []uint64 `json:"ids,omitempty"`
	// IsRecycled 标记是否删除，当查询已删除列表时，is_recycled=1
	IsRecycled int `json:"is_recycled,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页大小
	PageSize int `json:"page_size,omitempty"`
}

// URL implement Request interface
func (r CampaignsRequest) URL() string {
	return "v3/campaigns"
}

// Payload implement Request interface
func (r CampaignsRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	if r.Name != "" {
		p.AddValue("name", r.Name)
	}
	if len(r.EffectiveStatus) > 0 {
		buf, _ := json.Marshal(r.EffectiveStatus)
		p.AddValue("effective_status", string(buf))
	}
	if len(r.Objectives) > 0 {
		buf, _ := json.Marshal(r.Objectives)
		p.AddValue("objectives", string(buf))
	}
	if len(r.IDs) > 0 {
		buf, _ := json.Marshal(r.IDs)
		p.AddValue("ids", string(buf))
	}
	if r.IsRecycled > 0 {
		p.AddValue("is_recycled", strconv.Itoa(r.IsRecycled))
	}
	if r.Page > 0 {
		p.AddValue("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		p.AddValue("page_size", strconv.Itoa(r.PageSize))
	}
	return p
}

// CampaignsResponse 广告系列列表 API Response
type CampaignsResponse struct {
	model.BaseResponse
	model.PageInfo
	// List 广告系列列表
	List []Campaign `json:"list,omitempty"`
}
