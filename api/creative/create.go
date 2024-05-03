package creative

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/creative"
)

// Create 创建创意详
func Create(clt *core.SDKClient, accessToken string, req *creative.CreateRequest) (uint64, error) {
	var ret creative.CreateResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return 0, err
	}
	return ret.Data, nil
}
