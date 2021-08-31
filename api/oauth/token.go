package oauth

import (
	"github.com/bububa/biz-weibo/core"
	oauth "github.com/bububa/biz-weibo/model/oauth"
)

// AccessToken 创建access_token令牌
func AccessToken(clt *core.SDKClient, req *oauth.AccessTokenRequest) (*oauth.Token, error) {
	var ret oauth.TokenResponse
	if err := clt.Do("", req, &ret); err != nil {
		return nil, err
	}
	return &ret.Token, nil
}

// RefreshToken 刷新access_token令牌
func RefreshToken(clt *core.SDKClient, req *oauth.RefreshTokenRequest) (*oauth.Token, error) {
	var ret oauth.TokenResponse
	if err := clt.Do("", req, &ret); err != nil {
		return nil, err
	}
	return &ret.Token, nil
}
