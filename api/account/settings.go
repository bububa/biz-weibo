package account

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/account"
)

// SettingsSet 同步用户配置到服务器
func SettingsSet(clt *core.SDKClient, accessToken string, req *account.SettingsSetRequest) (*account.Settings, error) {
	var ret account.SettingsResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Settings, nil
}

// Settings 获取用户配置
func Settings(clt *core.SDKClient, accessToken string, req *account.SettingsRequest) (*account.Settings, error) {
	var ret account.SettingsResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Settings, nil
}
