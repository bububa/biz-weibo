package enum

// DeliverScope 投放范围（历史原因保留字段，1：微博站内、2：优选流量、3：新浪站内），建议使用详情接口获取投放范围，详见：获取广告计划详情，【附录-投放范围】
type DeliverScope int

const (
	// DeliverScope_WEIBO 微博站内
	DeliverScope_WEIBO DeliverScope = 1
	// DeliverScope_SMART 优选流量
	DeliverScope_SMART DeliverScope = 2
	// DeliverScope_SINA 新浪站内
	DeliverScope_SINA DeliverScope = 3
)
