package track

import (
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// ActivateRequest 线索API数据上报 API Request
type ActivateRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// Time 转化时间，如：1534730754456(毫秒)
	Time int64 `json:"time,omitempty"`
	// MarkID mark_id允许为空，此时为自然量
	MarkID string `json:"mark_id,omitempty"`
	// Behavior 客户行为码，behavior编码见下文
	Behavior Behavior `json:"behavior,omitempty"`
	// Host 回传来源，默认请填写您的域名，如：weibo.com
	Host string `json:"host,omitempty"`
}

// URL implement Request interface
func (r ActivateRequest) URL() string {
	return "v3/track/activate"
}

// Payload implement Request interface
func (r ActivateRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	}
	p.AddValue("time", strconv.FormatInt(r.Time, 10))
	if r.MarkID != "" {
		p.AddValue("mark_id", r.MarkID)
	}
	p.AddValue("behavior", strconv.Itoa(int(r.Behavior)))
	if r.Host != "" {
		p.AddValue("host", r.Host)
	}
	return p
}
