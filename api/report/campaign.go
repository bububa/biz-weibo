package report

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/report"
)

// CampaignData 广告系列数据
func CampaignData(clt *core.SDKClient, accessToken string, req *report.CampaignDataRequest) (report.CampaignData, error) {
	var ret report.CampaignDataResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return report.CampaignData{}, err
	}
	return ret.Data, nil
}
