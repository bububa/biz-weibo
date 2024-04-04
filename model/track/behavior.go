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
	// Behavior_FINANCE_COMPLETE 金融-完件
	Behavior_FINANCE_COMPLETE Behavior = 1024
	// Behavior_FINANCE_CREDITE 金融-授信
	Behavior_FINANCE_CREDITE Behavior = 1025
	// Behavior_FINANCE_LOAN 金融-放款
	Behavior_FINANCE_LOAN Behavior = 1026
	// Behavior_PAID_ECOMMERCE_ORDER 付款订单商品数
	Behavior_PAID_ECOMMERCE_ORDER Behavior = 2001
	// Behavior_MINIAPP_ACTIVE 激活(小程序)
	Behavior_MINIAPP_ACTIVE Behavior = 3001
	// Behavior_MINIAPP_REGISTER 注册(小程序)
	Behavior_MINIAPP_REGISTER Behavior = 3002
	// Behavior_MINIAPP_PAY 付费(小程序&微信)
	Behavior_MINIAPP_PAY Behavior = 3003
	// Behavior_WECHAT_FOLLOW 微信关注
	Behavior_WECHAT_FOLLOW Behavior = 3004
)
