package campaign

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/campaign"
)

// Update 更新系列预算或名称
func Update(clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (*campaign.Campaign, error) {
	var ret campaign.UpdateResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Campaign, nil
}
