package oauth

import "github.com/bububa/biz-weibo/model"

type GrantType string

const (
	AuthorizationCode GrantType = "authorization_code"
	RefreshToken      GrantType = "refresh_token"
)

// Token access_token令牌
type Token struct {
	// AccessToken
	AccessToken string `json:"access_token,omitempty"`
	// ExpiresIn
	ExpiresIn int64 `json:"expires_in,omitempty"`
	// RefreshToken
	RefreshToken string `json:"refresh_token,omitempty"`
	// RefreshTokenExpiresIn
	RefreshTokenExpiresIn int64 `json:"refresh_expires_in,omitempty"`
	// TokenType
	TokenType string `json:"token_type,omitempty"`
}

// TokenResponse access_token response
type TokenResponse struct {
	model.BaseResponse
	Token
}

// TokenRequest 创建access_token令牌 API Request
type TokenRequest struct {
	// ClientID 分配的appid
	ClientID string `json:"client_id,omitempty"`
	// GrandType 此处固定为authorization_code
	GrantType GrantType `json:"grant_type,omitempty"`
}

// URL implement Request interface
func (r TokenRequest) URL() string {
	return "oauth/token"
}

// AccessTokenRequest 创建access_token令牌 API Request
type AccessTokenRequest struct {
	TokenRequest
	// RedirectURI 创建应用时候填写的跳转url
	RedirectURI string `json:"redirect_uri,omitempty"`
	// Code 请求认证返回的code值
	Code string `json:"code,omitempty"`
}

// URL implement Request interface
func (r AccessTokenRequest) Payload() *model.Payload {
	p := new(model.Payload)
	p.AddValue("client_id", r.ClientID)
	p.AddValue("redirect_uri", r.RedirectURI)
	p.AddValue("grant_type", string(AuthorizationCode))
	p.AddValue("code", r.Code)
	return p
}

// RefreshTokenRequest 刷新access_token令牌 API Request
type RefreshTokenRequest struct {
	TokenRequest
	// RefreshToken refresh_token值
	RefreshToken string `json:"refresh_token,omitempty"`
}

// URL implement Request interface
func (r RefreshTokenRequest) Payload() *model.Payload {
	p := new(model.Payload)
	p.AddValue("client_id", r.ClientID)
	p.AddValue("grant_type", string(RefreshToken))
	p.AddValue("refresh_token", r.RefreshToken)
	return p
}
