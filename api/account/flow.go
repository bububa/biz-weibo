package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// Flow 获取广告账户流水信息
func Flow(clt *core.SDKClient, accessToken string, req *account.FlowRequest) (*account.FlowResponse, error) {
	var ret account.FlowResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
