package oauth

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/oauth"
)

// UserInfo 获取授权User信息
func UserInfo(clt *core.SDKClient, req *oauth.UserInfoRequest) (*oauth.UserInfo, error) {
	var ret oauth.UserInfoResponse
	if err := clt.Do("", req, &ret); err != nil {
		return nil, err
	}
	return &ret.Data, nil
}
