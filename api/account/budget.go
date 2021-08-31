package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// Budget 修改账户日限额
func Budget(clt *core.SDKClient, accessToken string, req *account.BudgetRequest) (*account.Account, error) {
	var ret account.BudgetResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Account, nil
}
