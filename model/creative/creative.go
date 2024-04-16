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
	// Name 创意名称
	Name string `json:"name,omitempty"`
	// ConfiguredStatus 开关状态 详见【附录-创意开关状态编码】
	ConfiguredStatus enum.ConfiguredStatus `json:"configured_status"`
	// EffectiveStatus 创意状态 【附录-创意状态编码】
	EffectiveStatus enum.CreativeEffectiveStatus `json:"effective_status"`
	// Version 创意版本，3（v3创意）、4（v4创意）
	Version int `json:"version"`
	// BillingType 出价方式 详见【附录-出价方式编码】
	BillingType enum.BillingType `json:"billing_type"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeType 创意类型：1 新建创意、2 复用已有
	CreativeType int `json:"creative_type"`
	// MediaMidComplex 复用类型：1 素材复用 、2 博文复用
	MediaMidComplex int `json:"media_mid_complex"`
	// MediaMidType 博文类型：1 已有博文、 2 聚宝盆代投
	MediaMidType int `json:"media_mid_type"`
	// Mid 博文复用时博文的mid
	Mid string `json:"mid,omitempty"`
	// ComponentMountSwitch 组件挂载开关：0关闭，1开启 当所选聚宝盆代投博文为标的支持的样式（详见【附录-各标的下支持的创意样式】），且博文不包含组件时（详见【附录-各标的样式下支持的组件类型】），支持组件挂载
	ComponentMountSwitch int `json:"component_mount_switch"`
	// ShortlinkText 短链文案
	ShortlinkText string `json:"shortlink_text"`
	// 营销目标 详见【附录-营销目标编码】
	Objective enum.Objective `json:"objective"`
	// PreviewMidURL 创意预览url
	PreviewMidURL string `json:"preview_mid_url"`
	// CreatedAt 创建时间
	CreatedAt string `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt string `json:"updated_at"`
	// Campaign 系列对象
	Campaign *campaign.Campaign `json:"campaign"`
	// Ad 计划对象
	Ad *ad.Ad `json:"ad"`
}
