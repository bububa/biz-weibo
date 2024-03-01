package oauth

import "github.com/bububa/biz-weibo/model"

type UserInfoRequest struct {
	AccessToken string `json:"access_token,omitempty"`
}

func (r UserInfoRequest) URL() string {
	return "oauth/get_user_info"
}

func (r UserInfoRequest) Payload() *model.Payload {
	p := new(model.Payload)
	p.AddValue("access_token", r.AccessToken)
	return p
}

type UserInfoResponse struct {
	model.BaseResponse
	Data UserInfo `json:"data"`
}

type UserInfo struct {
	Uid         uint64 `json:"uid"`
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
	UserType    int    `json:"user_type"`
}
