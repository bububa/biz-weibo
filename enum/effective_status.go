package enum

// EffectiveStatus 计划状态（超粉系统)，详见：【附录-计划状态】
type EffectiveStatus int

const (
	// EffectiveStatus_FINISHED 结束状态
	EffectiveStatus_FINISHED EffectiveStatus = -1
	// EffectiveStatus_DELETED 删除
	EffectiveStatus_DELETED EffectiveStatus = -99
	// EffectiveStatus_DEFAULT 默认状态
	EffectiveStatus_DEFAULT EffectiveStatus = 0
	// EffectiveStatus_PENDING 待投放
	EffectiveStatus_PENDING EffectiveStatus = 1
	// EffectiveStatus_ACTIVE 投放中
	EffectiveStatus_ACTIVE EffectiveStatus = 2
	// EffectiveStatus_USER_PAUSED 用户暂停
	EffectiveStatus_USER_PAUSED EffectiveStatus = 3
	// EffectiveStatus_ABNORMAL_ACCOUNT 账户异常
	EffectiveStatus_ABNORMAL_ACCOUNT EffectiveStatus = 5
	// EffectiveStatus_SYSTEM_PAUSED 系列暂停
	EffectiveStatus_SYSTEM_PAUSED EffectiveStatus = 6
	// EffectiveStatus_INSUFFICIENT_BALANCE 余额不足
	EffectiveStatus_INSUFFICIENT_BALANCE EffectiveStatus = 7
	// EffectiveStatus_EXCEEDED_ACCOUNT_DAILY_BUDGET 到达账户日预算
	EffectiveStatus_EXCEEDED_ACCOUNT_DAILY_BUDGET EffectiveStatus = 8
	// EffectiveStatus_EXCEEDED_AD_DAILY_BUDGET 到达系列预算
	EffectiveStatus_EXCEEDED_AD_DAILY_BUDGET EffectiveStatus = 9
	// EffectiveStatus_EXCEED_CAMPAIGN_DAILY_BUDGET 到达计划日预算
	EffectiveStatus_EXCEED_CAMPAIGN_DAILY_BUDGET EffectiveStatus = 10
	// EffectiveStatus_NO_AVALIABLE_CREATIVE 无可用创意
	EffectiveStatus_NO_AVALIABLE_CREATIVE EffectiveStatus = 11
	// EffectiveStatus_FENTIAO_ORDER_CREATING 粉条订单创建中
	EffectiveStatus_FENTIAO_ORDER_CREATING EffectiveStatus = 12
	// EffectiveStatus_DATA_MARKET_PACKAGE_PROCESSING 数据市场包处理中
	EffectiveStatus_DATA_MARKET_PACKAGE_PROCESSING EffectiveStatus = 13
	// EffectiveStatus_ELECTRONIC_FENCE_FAILED 电子围栏生成失败
	EffectiveStatus_ELECTRONIC_FENCE_FAILED EffectiveStatus = 14
)
