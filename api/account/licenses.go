package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// Licenses 资质列表
func Licenses(clt *core.SDKClient, accessToken string, req *account.LicensesRequest) (*account.LicensesResponse, error) {
	var ret account.LicensesResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
