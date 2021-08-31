package campaign

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/campaign"
)

// Campaigns 广告系列列表
func Campaigns(clt *core.SDKClient, accessToken string, req *campaign.CampaignsRequest) (*campaign.CampaignsResponse, error) {
	var ret campaign.CampaignsResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
