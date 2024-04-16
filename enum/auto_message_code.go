package enum

// AutoMessageCode 自动私信类型 ，详见：【附录-自动私信类型】
type AutoMessageCode int

const (
	// AutoMessageCode_COMMENT 评论
	AutoMessageCode_COMMENT AutoMessageCode = 14000005
	// AutoMessageCode_FORWARD 转发
	AutoMessageCode_FORWARD AutoMessageCode = 14000003
	// AutoMessageCode_LIKE 点赞
	AutoMessageCode_FAV AutoMessageCode = 14000098
	// AutoMessageCode_COLLECT 收藏
	AutoMessageCode_COLLECT AutoMessageCode = 14000045
)
