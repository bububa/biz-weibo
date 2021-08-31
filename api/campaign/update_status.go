package campaign

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/campaign"
)

// UpdateStatus 更新系列状态
func UpdateStatus(clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) ([]campaign.UpdateStatusResult, error) {
	var ret campaign.UpdateStatusResponse
	if err := clt.Do(accessToken, req, ret); err != nil {
		return nil, err
	}
	return []campaign.UpdateStatusResult(ret), nil
}
