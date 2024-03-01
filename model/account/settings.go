package account

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// SettingsRequest 用户配置 API Request
type SettingsRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Key 配置索引
	Key string `json:"key,omitempty"`
}

// URL implement Request interface
func (r SettingsRequest) URL() string {
	return "v3/account/settings"
}

// Payload implement Request interface
func (r SettingsRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	}
	p.AddValue("key", r.Key)
	return p
}

// SettingsSetRequest 同步用户配置到服务器 API Request
type SettingsSetRequest struct {
	SettingsRequest
	// Value 配置内容
	Value string `json:"value,omitempty"`
}

// Payload implement Request interface
func (r SettingsSetRequest) Payload() *model.Payload {
	buf, _ := json.Marshal(r)
	return model.NewPostPayload(buf)
}

// SettingsResponse 用户配置 API Response
type SettingsResponse struct {
	model.BaseResponse
	Settings
}

// Settings 用户配置
type Settings struct {
	// ID
	ID int64 `json:"id,omitempty"`
	// CustomerID
	CustomerID int64 `json:"customer_id,omitempty"`
	// Sid
	Sid int64 `json:"sid,omitempty"`
	// CloudKey
	CloudKey string `json:"cloud_key,omitempty"`
	// CloudVal
	CloudVal string `json:"cloud_val,omitempty"`
	// CreatedAt
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt
	UpdatedAt string `json:"updated_at,omitempty"`
}
