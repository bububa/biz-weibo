package account

import (
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// LicensesRequest 资质列表 API Request
type LicensesRequest struct {
	// AdvertiserID 管理的广告主id，为空则为token对应账户
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数
	Page int `json:"page,omitempty"`
	// Limit 每页条数
	Limit int `json:"limit,omitempty"`
	// Name 资质名称
	Name string `json:"name,omitempty"`
	// Status 资质证明状态(0:已提交 ,1:待审,2:审核通过,3:审核未通过,4:过期,5:部分通过审核,6:已推送待审,9:已删除)
	Status int `json:"status,omitempty"`
}

// URL implement Request interface
func (r LicensesRequest) URL() string {
	return "v3/account/licenses"
}

// Payload implement Request interface
func (r LicensesRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID <= 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	if r.Page > 0 {
		p.AddValue("page", strconv.Itoa(r.Page))
	}
	if r.Limit > 0 {
		p.AddValue("limit", strconv.Itoa(r.Limit))
	}
	if r.Name != "" {
		p.AddValue("name", r.Name)
	}
	p.AddValue("status", strconv.Itoa(r.Status))
	return p
}

// LicensesResponse 资质列表 API Response
type LicensesResponse struct {
	model.BaseResponse
	// Total
	Total int `json:"total,omitempty"`
	// PageTotal
	PageTotal int `json:"pageTotal,omitempty"`
	// PageSize
	PageSize int `json:"pageSize,omitempty"`
	// Page
	Page int `json:"page,omitempty"`
	// List
	List []License `json:"list,omitempty"`
}

// License 资质
type License struct {
	// PackageID 资质包id，注意：此返回字段为创意创建资质对应字段licenses
	PackageID int64 `json:"packageId,omitempty"`
	// PackageName 资质名称
	PackageName string `json:"packageName,omitempty"`
	// Description 资质描述
	Description string `json:"description,omitempty"`
	// Industry 资质行业编码
	Industry int64 `json:"industry,omitempty"`
	// BizID
	BizID int64 `json:"bizId,omitempty"`
	// Uid
	Uid int64 `json:"uid,omitempty"`
	// EntityID
	EntityID int64 `json:"entityId,omitempty"`
	// Version
	Version int `json:"version,omitempty"`
	// Version30
	Version30 int `json:"version30,omitempty"`
	// Pending
	Pending int `json:"pending,omitempty"`
	// Status
	Status int `json:"status,omitempty"`
	// Status30
	Status30 int `json:"status30,omitempty"`
	// ValidityTime
	ValidityTime int64 `json:"validityTime,omitempty"`
	// CreateTime
	CreateTime int64 `json:"createTime,omitempty"`
	// UpdateTIme
	UpdateTime int64 `json:"updateTime,omitempty"`
	// Reason
	Reason string `json:"reason,omitempty"`
	// Reason30
	Reason30 string `json:"reason30,omitempty"`
	// Ext
	Ext string `json:"ext,omitempty"`
}
