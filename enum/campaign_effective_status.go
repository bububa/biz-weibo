package enum

// CampaignEffectiveStatus 系列内部状态
type CampaignEffectiveStatus int

const (
	// CampaignEffectiveStatus_PAUSED 暂停
	CampaignEffectiveStatus_PAUSED CampaignEffectiveStatus = 0
	// CampaignEffectiveStatus_NORMAL 运行
	CampaignEffectiveStatus_NORMAL CampaignEffectiveStatus = 1
	// CampaignEffectiveStatus_DELETED 删除
	CampaignEffectiveStatus_DELETED CampaignEffectiveStatus = 2
)
