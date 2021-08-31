package track

// Behavior 客户行为码
type Behavior int

const (
	// Behavior_FORM_SUBMIT 表单提交
	Behavior_FORM_SUBMIT Behavior = 1001
	// Behavior_PHONE_CALL 电话拨打
	Behavior_PHONE_CALL Behavior = 1002
	// Behavior_CONSULT 有效咨询
	Behavior_CONSULT Behavior = 1003
	// Behavior_WECHAT_COPY 微信复制
	Behavior_WECHAT_COPY Behavior = 1004
	// Behavior_LANDING_PAGE 落地页访问
	Behavior_LANDING_PAGE = 1005
	// Behavior_DOWNLOAD 下载开始
	Behavior_DOWNLOAD Behavior = 1006
	// Behavior_PURCHASE 商品购买
	Behavior_PURCHASE Behavior = 1007
	// Behavior_OTHERS 其他(other)
	Behavior_OTHERS Behavior = 1100
)
