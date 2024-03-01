package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// Asset 获取账户资产
func Asset(clt *core.SDKClient, accessToken string, advertiserID uint64) (*account.Asset, error) {
	var (
		ret account.AssetResponse
		req = &account.AssetRequest{
			AdvertiserID: advertiserID,
		}
	)
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Asset, nil
}
