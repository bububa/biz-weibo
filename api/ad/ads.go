package ad

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/ad"
)

// Ads 广告计划列表
func Ads(clt *core.SDKClient, accessToken string, req *ad.AdsRequest) (*ad.AdsResponse, error) {
	var ret ad.AdsResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
