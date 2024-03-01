package report

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/report"
)

// CreativeData  广告创意数据
func CreativeData(clt *core.SDKClient, accessToken string, req *report.CreativeDataRequest) (report.CreativeData, error) {
	var ret report.CreativeDataResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return report.CreativeData{}, err
	}
	return ret.Data, nil
}
