package account

import (
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// FlowRequest 获取广告账户流水信息 API Request
type FlowRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// Page 页数
	Page int `json:"page,omitempty"`
	// PageSize 每页条数
	PageSize int `json:"page_size,omitempty"`
	// StartDate 开始时间，如：2020-08-01
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，如：2020-08-02
	EndDate string `json:"end_date,omitempty"`
}

// URL implement Request interface
func (r FlowRequest) URL() string {
	return "v3/account/flow"
}

// Payload implement Request interface
func (r FlowRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	}
	if r.Page > 0 {
		p.AddValue("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		p.AddValue("page_size", strconv.Itoa(r.PageSize))
	}
	if r.StartDate != "" {
		p.AddValue("start_date", r.StartDate)
	}
	if r.EndDate != "" {
		p.AddValue("end_date", r.EndDate)
	}
	return p
}

// FlowResponse 获取广告账户流水信息 API Response
type FlowResponse struct {
	model.BaseResponse
	// Count 条数
	Count int `json:"count,omitempty"`
	// SumExpenditue 总支出
	SumExpenditue float64 `json:"sum_expenditue,omitempty"`
	// SumIncom 总收入
	SumIncom float64 `json:"sum_incom,omitempty"`
	// List
	List []Flow `json:"list,omitempty"`
}

// Flow
type Flow struct {
	// Income 收入
	Income float64 `json:"income,omitempty"`
	// Expenditue 支出
	Expenditue float64 `json:"expenditue,omitempty"`
	// Balance 余额
	Balance float64 `json:"balance,omitempty"`
	// Remark 备注，json格式
	Remark string `json:"remark,omitempty"`
	// Impression 曝光量
	Impression int64 `json:"impression,omitempty"`
	// CustID 广告主uid
	CustID int64 `json:"cust_id,omitempty"`
	// AdType 产品线类型
	AdType int `json:"ad_type,omitempty"`
}
