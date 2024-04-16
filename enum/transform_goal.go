package enum

// TransformGoal 转化目标，详见：【附录-转化目标】
type TransformGoal int

const (
	// TransformGoal_ACTIVATE 激活
	TransformGoal_ACTIVATE TransformGoal = 86004001
	// TransformGoal_APP_LAUNCH app唤起
	TransformGoal_APP_LAUNCH TransformGoal = 86004002
	// TransformGoal_WECHAT_FOLLOW 关注（微博关注）
	TransformGoal_WECHAT_FOLLOW TransformGoal = 86002003
	// TransformGoal_RESERVATION 预约
	TransformGoal_RESERVATION TransformGoal = 86004003
	// TransformGoal_MINIAPP_REGISTER 小程序注册
	TransformGoal_MINIAPP_REGISTER TransformGoal = 86005003
	// TransformGoal_PRIVATE_CONSULT 私信咨询
	TransformGoal_PRIVATE_CONSULT TransformGoal = 86005002
	// TransformGoal_FORM_SUBMIT 表单提交
	TransformGoal_FORM_SUBMIT TransformGoal = 86005001
	// TransformGoal_FIRST_PAY 首次付费
	TransformGoal_FIRST_PAY TransformGoal = 86007003
	// TransformGoal_PAY 每次付费
	TransformGoal_PAY TransformGoal = 86007005
	// TransformGoal_WECHAT_PAY 微信付费
	TransformGoal_WECHAT_PAY TransformGoal = 86002005
	// TransformGoal_MINIAPP_PAY 小程序付费
	TransformGoal_MINIAPP_PAY TransformGoal = 86005005
	// TransformGoal_MINIAPP_FANS 小程序加粉 (微信关注、微信加粉，合并至小程序加粉)
	TransformGoal_MINIAPP_FANS TransformGoal = 86005006
)
