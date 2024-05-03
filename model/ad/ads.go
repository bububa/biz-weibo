package ad

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
)

// AdssRequest 广告计划列表 API Request
type AdsRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 按照广告系列名称关键字过滤，若无需求请勿传值
	Name string `json:"name,omitempty"`
	// EffectiveStatus 按照广告系列状态过滤，若无需求请勿传值
	EffectiveStatus []enum.EffectiveStatus `json:"effective_status,omitempty"`
	// Objectives 按照营销目标过滤，若无需求请勿传值
	Objectives []enum.Objective `json:"objectives,omitempty"`
	// IDs id筛选
	IDs []uint64 `json:"ids,omitempty"`
	// BillingTypes 按照计费模式过滤，若无需求请勿传值
	BillingTypes []enum.BillingType `json:"billing_type,omitempty"`
	// IsRecycled 标记是否删除,当查询已删除列表时，is_recycled=1
	IsRecycled int `json:"is_recycled,omitempty"`
	// CampaignID 所属的系列id 【示例】 1234
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// AdStartTime 根据计划创建时间筛,开始时间，格式：yyyy-MM-dd，【示例】 "2022-10-01"
	AdStartTime string `json:"ad_start_time,omitempty"`
	// AdEndTime 根据计划创建时间筛，结束时间，格式：yyyy-MM-dd，【示例】 "2022-10-01"
	AdEndTime string `json:"ad_end_time,omitempty"`
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
	if len(r.BillingTypes) > 0 {
		buf, _ := json.Marshal(r.BillingTypes)
		p.AddValue("billing_type", string(buf))
	}
	if r.IsRecycled > 0 {
		p.AddValue("is_recycled", strconv.Itoa(r.IsRecycled))
	}
	if r.CampaignID > 0 {
		p.AddValue("campaing_id", strconv.FormatUint(r.CampaignID, 10))
	}
	if r.AdStartTime != "" {
		p.AddValue("ad_start_time", r.AdStartTime)
	}
	if r.AdEndTime != "" {
		p.AddValue("ad_end_time", r.AdEndTime)
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
	model.PageInfo
	// List 广告计划列表
	List []Ad `json:"list,omitempty"`
}
