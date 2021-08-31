package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// ReverseSpendCap 修改账户次日日限额
func ReverseSpendCap(clt *core.SDKClient, accessToken string, req *account.ReverseSpendCapRequest) (*account.Account, error) {
	var ret account.ReverseSpendCapResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Account, nil
}
