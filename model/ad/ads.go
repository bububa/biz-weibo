package ad

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// AdssRequest 广告计划列表 API Request
type AdsRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// Name 按照广告系列名称关键字过滤，若无需求请勿传值
	Name string `json:"name,omitempty"`
	// EffectiveStatus 按照广告系列状态过滤，若无需求请勿传值
	EffectiveStatus []int `json:"effective_status,omitempty"`
	// Objectives 按照营销目标过滤，若无需求请勿传值
	Objectives []model.Objective `json:"objectives,omitempty"`
	// IDs id筛选
	IDs []int64 `json:"ids,omitempty"`
	// BillingTypes 按照计费模式过滤，若无需求请勿传值
	BillingTypes []BillingType `json:"billing_type,omitempty"`
	// MidSources 广告类型
	MidSources []int `json:"mid_source,omitempty"`
	// GuaranteedDelivery 是否定价保量
	GuaranteedDelivery bool `json:"guaranteed_delivery,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页大小
	PageSize int `json:"page_size,omitempty"`
}

// URL implement Request interface
func (r AdsRequest) URL() string {
	return "v3/ads"
}

// Payload implement Request interface
func (r AdsRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
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
	if len(r.BillingTypes) > 0 {
		buf, _ := json.Marshal(r.BillingTypes)
		p.AddValue("billing_type", string(buf))
	}
	if len(r.MidSources) > 0 {
		buf, _ := json.Marshal(r.MidSources)
		p.AddValue("mid_source", string(buf))
	}
	if r.GuaranteedDelivery {
		p.AddValue("guaranteed_delivery", "true")
	}
	if r.Page > 0 {
		p.AddValue("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		p.AddValue("page_size", strconv.Itoa(r.PageSize))
	}
	return p
}

// AdsResponse 广告计划列表 API Response
type AdsResponse struct {
	model.BaseResponse
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 分页数
	PageSize int `json:"page_size,omitempty"`
	// PageTotal 总页数
	PageTotal int `json:"page_total,omitempty"`
	// Total 总数
	Total int `json:"total,omitempty"`
	// List 广告计划列表
	List []Ad `json:"list,omitempty"`
}
