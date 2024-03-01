package report

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/report"
)

// AdData 广告计划数据
func AdData(clt *core.SDKClient, accessToken string, req *report.AdDataRequest) (report.AdData, error) {
	var ret report.AdDataResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return report.AdData{}, err
	}
	return ret.Data, nil
}
