package enum

// CreativeEffectiveStatus 创意状态
type CreativeEffectiveStatus int

const (
	// CreativeEffectiveStatus_AUDITING 审核中
	CreativeEffectiveStatus_AUDITING CreativeEffectiveStatus = 0
	// CreativeEffectiveStatus_NORMAL 正常
	CreativeEffectiveStatus_NORMAL CreativeEffectiveStatus = 1
	// CreativeEffectiveStatus_AUDIT_ABNORMAL 异常 审核异常
	CreativeEffectiveStatus_AUDIT_ABNORMAL CreativeEffectiveStatus = 2
	// CreativeEffectiveStatus_UID_ABNORMAL 异常 uid/mid异常
	CreativeEffectiveStatus_UID_ABNORMAL CreativeEffectiveStatus = 3
	// CreativeEffectiveStatus_PAUSED 暂停
	CreativeEffectiveStatus_PAUSED CreativeEffectiveStatus = 4
	// CreativeEffectiveStatus_PUBLISH_ABNORMAL 异常 （自动）发布异常
	CreativeEffectiveStatus_PUBLISH_ABNORMAL CreativeEffectiveStatus = 6
	// CreativeEffectiveStatus_PREPUBLISH 预发布
	CreativeEffectiveStatus_PREPUBLISH CreativeEffectiveStatus = 7
	// CreativeEffectgiveStatus_DELETED 删除
	CreativeEffectiveStatus_DELETED CreativeEffectiveStatus = 9
	// CreativeEffectiveStatus_DRAFT 草稿
	CreativeEffectiveStatus_DRAFT CreativeEffectiveStatus = 12
	// CreativeEffectiveStatus_FAILED 创意生成失败
	CreativeEffectiveStatus_FAILED CreativeEffectiveStatus = 13
)
