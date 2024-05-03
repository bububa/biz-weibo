package creative

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/creative"
)

// UpdateStatusBatch 创意状态更新
func UpdateStatusBatch(clt *core.SDKClient, accessToken string, req *creative.UpdateStatusBatchRequest) ([]creative.UpdateResult, error) {
	var ret creative.UpdateStatusBatchResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
