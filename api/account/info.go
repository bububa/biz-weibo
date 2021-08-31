package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// Info 获取广告账户信息
func Info(clt *core.SDKClient, accessToken string, advertiserID int64) (*account.Account, error) {
	var (
		ret account.InfoResponse
		req = &account.InfoRequest{
			AdvertiserID: advertiserID,
		}
	)
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Account, nil
}
