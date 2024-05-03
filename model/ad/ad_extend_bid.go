package ad

import "github.com/bububa/biz-weibo/enum"

// AdExtendBid 一键起量对象
type AdExtendBid struct {
	// ID 计划一键起量记录唯一索引
	ID uint64 `json:"id,omitempty"`
	// AdID 计划Id
	AdID uint64 `json:"ad_id,omitempty"`
	// ExtendStatus 一键起量状态 详见：【附录-计划一键起量状态】
	ExtendStatus enum.AdExtendStatus `json:"extend_status,omitempty"`
	// ExtendEndType 起量关闭原因 详见：【附录-起量关闭原因】
	ExtendEndType enum.AdExtendEndType `json:"extend_end_type,omitempty"`
	// CreatedAt 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// DataVersion 数据版本
	DataVersion int `json:"data_version,omitempty"`
}
