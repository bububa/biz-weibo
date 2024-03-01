package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// DeleteReverseBudget 删除账户次日日限额
func DeleteReverseBudget(clt *core.SDKClient, accessToken string, advertiserID uint64) (*account.Account, error) {
	var (
		ret account.DeleteReverseBudgetResponse
		req = &account.DeleteReverseBudgetRequest{
			AdvertiserID: advertiserID,
		}
	)
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Account, nil
}
