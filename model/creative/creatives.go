package creative

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model"
)

// 广告创意请求结构体
type CreativesRequest struct {
	AdvertiserID     uint64                         `json:"advertiser_id,omitempty"`     // 您管理的广告主id，为空则为token对应账户不使用请勿下发
	ConfiguredStatus []enum.ConfiguredStatus        `json:"configured_status,omitempty"` // 开关状态 详见【附录-创意开关状态编码】 格式：configured_status=1,2
	EffectiveStatus  []enum.CreativeEffectiveStatus `json:"effective_status,omitempty"`  // 创意状态 详见【附录-创意状态编码】
	Objective        []enum.Objective               `json:"objective,omitempty"`         // 营销目标 详见【附录-营销目标编码】
	CreativeIDList   []uint64                       `json:"creative_id_list,omitempty"`  // 创意id数组
	BillingType      []enum.BillingType             `json:"billing_type,omitempty"`      // 出价方式 详见【附录-出价方式编码】
	Name             string                         `json:"name,omitempty"`              // 创意名称
	AdID             uint64                         `json:"ad_id,omitempty"`             // 计划id
	Page             int                            `json:"page,omitempty"`              // 页码
	PageSize         int                            `json:"page_size,omitempty"`         // 每页数据
}

// URL implement Request interface
func (r CreativesRequest) URL() string {
	return "v4/creatives"
}

// Payload implement Request interface
func (r CreativesRequest) Payload() *model.Payload {
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
	if len(r.Objective) > 0 {
		buf, _ := json.Marshal(r.Objective)
		p.AddValue("objectives", string(buf))
	}
	if len(r.CreativeIDList) > 0 {
		buf, _ := json.Marshal(r.CreativeIDList)
		p.AddValue("creative_id_list", string(buf))
	}
	if len(r.BillingType) > 0 {
		buf, _ := json.Marshal(r.BillingType)
		p.AddValue("billing_type", string(buf))
	}
	if len(r.ConfiguredStatus) > 0 {
		buf, _ := json.Marshal(r.ConfiguredStatus)
		p.AddValue("configured_status", string(buf))
	}
	if r.AdID > 0 {
		p.AddValue("ad_id", strconv.FormatUint(r.AdID, 10))
	}
	if r.Page > 0 {
		p.AddValue("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		p.AddValue("page_size", strconv.Itoa(r.PageSize))
	}
	return p
}

type CreativesResponse struct {
	model.BaseResponse
	Data CreativeData `json:"data"`
}

type CreativeData struct {
	List []Creative `json:"list,omitempty"`
	model.PageInfo
}
