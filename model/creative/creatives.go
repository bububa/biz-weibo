package creative

import (
	"encoding/json"
	"github.com/bububa/biz-weibo/model"
	"strconv"
)

// 广告创意请求结构体
type CreativesRequest struct {
	AdvertiserID     uint64            `json:"advertiser_id,omitempty"`     // 您管理的广告主id，为空则为token对应账户不使用请勿下发
	ConfiguredStatus []int             `json:"configured_status,omitempty"` // 开关状态 详见【附录-创意开关状态编码】 格式：configured_status=1,2
	EffectiveStatus  []int             `json:"effective_status,omitempty"`  // 创意状态 详见【附录-创意状态编码】
	Objective        []model.Objective `json:"objective,omitempty"`         // 营销目标 详见【附录-营销目标编码】
	CreativeIDList   []int             `json:"creative_id_list,omitempty"`  // 创意id数组
	BillingType      []int             `json:"billing_type,omitempty"`      // 出价方式 详见【附录-出价方式编码】
	Name             string            `json:"name,omitempty"`              // 创意名称
	AdID             int               `json:"ad_id,omitempty"`             // 计划id
	Page             int               `json:"page,omitempty"`              // 页码
	PageSize         int               `json:"page_size,omitempty"`         // 每页数据
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
		p.AddValue("ad_id", strconv.Itoa(r.AdID))
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
	List      []Creative `json:"list"`
	Total     int        `json:"total"`
	PageTotal int        `json:"page_total"`
	Page      int        `json:"page"`
}

type Creative struct {
	ID               int      `json:"id"`                // 创意id
	Name             string   `json:"name"`              // 创意名称
	ConfiguredStatus int      `json:"configured_status"` // 开关状态 详见【附录-创意开关状态编码】
	EffectiveStatus  int      `json:"effective_status"`  // 创意状态 【附录-创意状态编码】
	Version          int      `json:"version"`           // 创意版本，3（v3创意）、4（v4创意）
	BillingType      int      `json:"billing_type"`      // 出价方式 详见【附录-出价方式编码】
	CreativeType     int      `json:"creative_type"`     // 创意类型：1 新建创意、2 复用已有
	MediaMidComplex  int      `json:"media_mid_complex"` // 复用类型：1 素材复用 、2 博文复用
	MediaMidType     int      `json:"media_mid_type"`    // 博文类型：1 已有博文、 2 聚宝盆代投
	Objective        int      `json:"objective"`         // 营销目标 详见【附录-营销目标编码】
	PreviewMidURL    string   `json:"preview_mid_url"`   // 创意预览url
	CreatedAt        string   `json:"created_at"`        // 创建时间
	UpdatedAt        string   `json:"updated_at"`        // 更新时间
	Campaign         Campaign `json:"campaign"`          // 系列对象
	Ad               Ad       `json:"ad"`                // 计划对象
}

// 系列对象
type Campaign struct {
	ID              int    `json:"id"`               // 系列id
	Name            string `json:"name"`             // 系列名称
	EffectiveStatus int    `json:"effective_status"` // 系列状态
	Objective       int    `json:"objective"`        // 系列的营销目标
	LifetimeBudget  string `json:"lifetime_budget"`  // 系列日预算
}

// 计划对象
type Ad struct {
	ID              int    `json:"id"`               // 计划id
	Name            string `json:"name"`             // 计划名称
	EffectiveStatus int    `json:"effective_status"` // 计划状态
	BillingType     int    `json:"billing_type"`     // 出价方式
	Objective       int    `json:"objective"`        // 计划的营销目标
	DailyBudget     string `json:"daily_budget"`     // 计划日预算
	BidAmount       string `json:"bid_amount"`       // 出价方式
}
