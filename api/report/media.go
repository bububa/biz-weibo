package report

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/report"
)

// MediaData  广告素材数据
func MediaData(clt *core.SDKClient, accessToken string, req *report.MediaDataRequest) (report.MediaData, error) {
	var ret report.MediaDataResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return report.MediaData{}, err
	}
	return ret.Data, nil
}
