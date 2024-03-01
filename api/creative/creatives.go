package creative

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/creative"
)

func Creatives(clt *core.SDKClient, accessToken string, req *creative.CreativesRequest) (*creative.CreativeData, error) {
	var ret creative.CreativesResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret.Data, nil
}
