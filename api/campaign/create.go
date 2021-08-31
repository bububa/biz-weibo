package campaign

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/campaign"
)

// Create 创建广告系列
func Create(clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (*campaign.Campaign, error) {
	var ret campaign.CreateResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Campaign, nil
}
