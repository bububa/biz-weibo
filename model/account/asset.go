package account

import (
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// AssetRequest 获取账户资产 API Request
type AssetRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// URL implement Request interface
func (r AssetRequest) URL() string {
	return "v3/asset"
}

// Payload implement Request interface
func (r AssetRequest) Payload() *model.Payload {
	if r.AdvertiserID <= 0 {
		return nil
	}
	p := new(model.Payload)
	p.AddValue("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return p
}

// AssetResponse 获取账户资产 API Response
type AssetResponse struct {
	model.BaseResponse
	Asset
}

// Asset 账户资产
type Asset struct {
	// Balance 账户余额
	Balance string `json:"balance,omitempty"`
	// RealTimeConsume 当天实时消耗
	RealTimeConsume string `json:"real_time_consume,omitempty"`
	// Msg 成功返回“ok”；失败为具体失败原因
	Msg string `json:"msg,omitempty"`
}
