package creative

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/creative"
)

// Creatives 获取创意列表
func Creatives(clt *core.SDKClient, accessToken string, req *creative.CreativesRequest) (*creative.CreativesData, error) {
	var ret creative.CreativesResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
