package enum

// AdExtendEndType 起量关闭原因
type AdExtendEndType int

const (
	// AdExtendEndType_DEFAULT 默认状态
	AdExtendEndType_DEFAULT AdExtendEndType = 0
	// AdExtendEndType_USER_CLOSED 用户关闭
	AdExtendEndType_USER_CLOSED AdExtendEndType = 1
	// AdExtendEndType_LEARNING_SUCCESS 学习期成功
	AdExtendEndType_LEARNING_SUCCESS AdExtendEndType = 2
	// AdExtendEndType_LEARNING_FAILED 学习期失败
	AdExtendEndType_LEARNING_FAILED AdExtendEndType = 3
	// AdExtendEndType_EXPIRED 到时关闭
	AdExtendEndType_EXPIRED AdExtendEndType = 4
	// AdExtendEndType_REALTIME_CLOSED 实时流关闭
	AdExtendEndType_REALTIME_CLOSED AdExtendEndType = 5
)
