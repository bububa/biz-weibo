package insights

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/insights"
)

// Effect 效果统计分析
func Effect(clt *core.SDKClient, accessToken string, req *insights.EffectRequest) (*insights.EffectResponse, error) {
	var ret insights.EffectResponse
	if err := clt.Do(accessToken, req, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
