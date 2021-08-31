package oauth

import "github.com/bububa/biz-weibo/model"

type Scope string

const (
	ROScope Scope = "ads_read"
	RWScope Scope = "ads_management"
)

// AuthorizeRequest 授权请求参数
type AuthorizeRequest struct {
	// ClientID 分配的appid
	ClientID string `json:"client_id,omitempty"`
	// RedirectURI 创建应用时候填写的跳转url
	RedirectURI string `json:"redirect_uri,omitempty"`
	// ResponseType 此处固定为code
	ResponseCode string `json:"response_code,omitempty"`
	// Scope ads_read 只读， ads_management 读写
	Scope Scope `json:"scope,omitempty"`
	// State 用来跟踪授权状态的字符串（防跨站攻击）
	State string `json:"state,omitempty"`
}

func (r AuthorizeRequest) URL() string {
	return "oauth/authorize"
}

func (r AuthorizeRequest) Payload() *model.Payload {
	p := new(model.Payload)
	p.AddValue("client_id", r.ClientID)
	p.AddValue("redirect_uri", r.RedirectURI)
	p.AddValue("response_code", "code")
	p.AddValue("scope", string(r.Scope))
	if r.State != "" {
		p.AddValue("state", r.State)
	}
	return p
}
