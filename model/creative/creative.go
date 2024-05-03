package creative

import (
	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model/ad"
	"github.com/bububa/biz-weibo/model/campaign"
)

// Creative 创意
type Creative struct {
	// ID 创意id
	ID uint64 `json:"id,omitempty"`
	// ConfiguredStatus 开关状态 详见【附录-创意开关状态编码】
	ConfiguredStatus enum.ConfiguredStatus `json:"configured_status,omitempty"`
	// EffectiveStatus 创意状态 【附录-创意状态编码】
	EffectiveStatus enum.CreativeEffectiveStatus `json:"effective_status,omitempty"`
	// Version 创意版本，3（v3创意）、4（v4创意）
	Version int `json:"version,omitempty"`
	// BillingType 出价方式 详见【附录-出价方式编码】
	BillingType enum.BillingType `json:"billing_type"`
	// CreativeType 创意类型：1 新建创意、2 复用已有
	CreativeType int `json:"creative_type,omitempty"`
	// MediaMidComplex 复用类型：1 素材复用 、2 博文复用、3聚宝盆素材复用
	MediaMidComplex int `json:"media_mid_complex,omitempty"`
	// MediaMidType 博文类型：1 已有博文、 2 聚宝盆代投
	MediaMidType int `json:"media_mid_type,omitempty"`
	// Objective 营销目标 详见【附录-营销目标编码】
	Objective enum.Objective `json:"objective"`
	// PreviewMidURL 创意预览url
	PreviewMidURL string `json:"preview_mid_url,omitempty"`
	// CreatedAt 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// Campaign 系列对象
	Campaign *campaign.Campaign `json:"campaign,omitempty"`
	// Ad 计划对象
	Ad *ad.Ad `json:"ad"`
}
