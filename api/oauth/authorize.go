package oauth

import (
	"github.com/bububa/biz-weibo/core"
	oauth "github.com/bububa/biz-weibo/model/oauth"
)

// Authorize 获取Authorization Code
func Authorize(clt *core.SDKClient, req *oauth.AuthorizeRequest) string {
	return clt.RequestURL(req.URL(), req.Payload())
}
