package creative

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/creative"
)

// AdRec 智能创意文案
func AdRec(clt *core.SDKClient, accessToken string, req *creative.AdRecRequest) ([]creative.AdRec, error) {
	var ret creative.AdRecResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return ret.Data, nil
}
