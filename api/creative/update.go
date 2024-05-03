package creative

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/creative"
)

// Update 创意编辑
func Update(clt *core.SDKClient, accessToken string, req *creative.UpdateRequest) (uint64, error) {
	var ret creative.UpdateResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return 0, err
	}
	return ret.Data, nil
}
