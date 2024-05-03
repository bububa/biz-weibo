package creative

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/creative"
)

// Info 获取创意详情
func Info(clt *core.SDKClient, accessToken string, req *creative.InfoRequest) (*creative.CreativeDetail, error) {
	var ret creative.InfoResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
