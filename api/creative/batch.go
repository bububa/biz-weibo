package creative

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/creative"
)

// Batch 批量获取创意详情
func Batch(clt *core.SDKClient, accessToken string, req *creative.BatchRequest) ([]creative.CreativeDetail, error) {
	var ret creative.BatchResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
