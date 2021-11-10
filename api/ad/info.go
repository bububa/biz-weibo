package ad

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/ad"
)

// Info 广告计划详情
func Info(clt *core.SDKClient, accessToken string, req *ad.InfoRequest) (*ad.Ad, error) {
	var ret ad.InfoResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Ad, nil
}
