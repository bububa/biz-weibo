package enum

// AdExtendStatus 一键起量状态
type AdExtendStatus int

const (
	AdExtendStatus_NOT_STARTED AdExtendStatus = 1
	AdExtendStatus_STARTED     AdExtendStatus = 2
	AdExtendStatus_RUNNING     AdExtendStatus = 3
	AdExtendStatus_COMPLETE    AdExtendStatus = 4
)
