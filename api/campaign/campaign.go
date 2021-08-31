package campaign

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/campaign"
)

// Campaign 广告系列详情
func Campaign(clt *core.SDKClient, accessToken string, req *campaign.CampaignRequest) (*campaign.Campaign, error) {
	var ret campaign.CampaignResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Campaign, nil
}
