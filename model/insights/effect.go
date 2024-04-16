package insights

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// EffectRequest 效果统计分析
type EffectRequest struct {
	// AdvertiserID 广告主id，此为您管理的广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data
	Data EffectRequestData `json:"data,omitempty"`
}

// EffectRequestData
type EffectRequestData struct {
	// TimeInterval 查询的开始时间和结束时间
	TimeInterval []string `json:"time_interval,omitempty"`
	// Field 返回的指标字段
	Field []string `json:"field,omitempty"`
	// Page 请求的页数(从1开始)
	Page int `json:"page,omitempty"`
	// Rows 每页的行数
	Rows int `json:"rows,omitempty"`
	// Granularity 数据group by的粒度
	Granularity []string `json:"granularity,omitempty"`
	// OrderBy 排序的字段
	OrderBy []string `json:"order_by,omitempty"`
	// OrderMode 排序方式，asc or desc
	OrderMode string `json:"order_mode,omitempty"`
	// Param 查询结果的限制条件，key为筛选字段，不填为不限
	Param map[string][]int64 `json:"param,omitempty"`
}

// URL implement Request interface
func (r EffectRequest) URL() string {
	return "v3/insights/effect"
}

// Payload implement Request interface
func (r EffectRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	buf, _ := json.Marshal(r.Data)
	p.AddValue("data", string(buf))
	return p
}

// EffectResponse 效果统计分析 API Response
type EffectResponse struct {
	model.BaseResponse
	Pages int       `json:"pages,omitempty"`
	Page  int       `json:"page,omitempty"`
	Rows  int       `json:"rows,omitempty"`
	Sum   *Metrics  `json:"sum,omitempty"`
	List  []Metrics `json:"list,omitempty"`
}
