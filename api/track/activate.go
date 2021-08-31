package track

import (
	"github.com/bububa/biz-weibo/core"
	"github.com/bububa/biz-weibo/model/track"
)

// Activate 线索API数据上报
func Activate(clt *core.SDKClient, accessToken string, req *track.ActivateRequest) error {
	return clt.Do(accessToken, req, nil)
}
