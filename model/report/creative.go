package report

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
)

type CreativeDataRequest struct {
	AdvertiserID uint64              `json:"advertiser_id,omitempty"` // 您管理的广告主id，为空则为token对应账户 （不使用请勿下发）
	StartDate    string              `json:"start_date"`              // 起始日期, 格式YYYY-MM-DD
	EndDate      string              `json:"end_date"`                // 结束日期,格式YYYY-MM-DD
	GroupBy      string              `json:"group_by,omitempty"`      // 分组粒度：date、hour、sum，默认：date
	OrderField   string              `json:"order_field,omitempty"`   // 排序字段，如pv
	OrderType    string              `json:"order_type,omitempty"`    // 排序方式，默认desc。asc、desc
	Page         int                 `json:"page,omitempty"`          // 默认1
	PageSize     int                 `json:"page_size,omitempty"`     // 默认20，最大200
	Fields       []string            `json:"fields"`                  // 指定需要的指标名称，参考数据报表指标说明
	Filtering    *CreativeDataFilter `json:"filtering,omitempty"`     // 筛选条件
}

type CreativeDataFilter struct {
	CreativeIds []uint64         `json:"creative_ids"`
	Objective   []enum.Objective `json:"objective"`
	FundType    int              `json:"fund_type"`
}

// URL implement Request interface
func (r CreativeDataRequest) URL() string {
	return "v3/report/effect/get_creative_data"
}

// Payload implement Request interface
func (r CreativeDataRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	p.AddValue("start_date", r.StartDate)
	p.AddValue("end_date", r.EndDate)
	if r.GroupBy != "" {
		p.AddValue("group_by", r.GroupBy)
	}
	if r.OrderField != "" {
		p.AddValue("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		p.AddValue("order_type", r.OrderType)
	}
	if r.Page > 0 {
		p.AddValue("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		p.AddValue("page_size", strconv.Itoa(r.PageSize))
	}
	for _, field := range r.Fields {
		p.AddValue("fields", field)
	}
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		p.AddValue("filtering", string(filter))
	}
	return p
}

type CreativeDataResponse struct {
	model.BaseResponse
	Data CreativeData `json:"data"`
}

type CreativeData struct {
	List []CreativeDataMetrics `json:"list"`
	Sum  Metrics               `json:"sum"`
	PageInfo
}

type CreativeDataMetrics struct {
	CreativeID int64 `json:"creative_id"`
	Metrics
}
