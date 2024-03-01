package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// Info 获取广告账户信息
func ManageList(clt *core.SDKClient, accessToken string, req *account.ManageListRequest) (*account.ManageListResponse, error) {
	ret := new(account.ManageListResponse)
	if err := clt.Do(accessToken, req, ret); err != nil {
		return nil, err
	}
	return ret, nil
}
