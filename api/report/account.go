package report

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/report"
)

// AccountData 广告主数据
func AccountData(clt *core.SDKClient, accessToken string, req *report.AccountDataRequest) (report.AccountData, error) {
	var ret report.AccountDataResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return report.AccountData{}, err
	}
	return ret.Data, nil
}
