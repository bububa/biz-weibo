package enum

// SubjectObjType 投放标类型，详见：【附录-投放标的类型】
type SubjectObjType int

const (
	// SubjectObjType_DEFAULT 默认值
	SubjectObjType_DEFAULT SubjectObjType = -1
	// SubjectObjType_UNSPECIFIED 无
	SubjectObjType_UNSPECIFIED SubjectObjType = 0
	// SubjectObjType_LANDINGPAGE 落地页
	SubjectObjType_LANDINGPAGE SubjectObjType = 1
	// SubjectObjType_DONWLOAD_LINK 下载链接
	SubjectObjType_DONWLOAD_LINK SubjectObjType = 2
	// SubjectObjType_WECHAT_MINIAPP 微信小程序
	SubjectObjType_WECHAT_MINIAPP SubjectObjType = 3
	// SubjectObjType_PRIVATE_MESSAGGE 私信箱
	SubjectObjType_PRIVATE_MESSAGGE SubjectObjType = 4
	// SubjectObjType_DEEPLINK 直达链接
	SubjectObjType_DEEPLINK SubjectObjType = 5
	// SubjectObjType_WECHAT 微信号
	SubjectObjType_WECHAT SubjectObjType = 6
	// SubjectObjType_WEIBO_FOLLOW 微博关注
	SubjectObjType_WEIBO_FOLLOW SubjectObjType = 7
	// SubjectObjType_SHOP_PRODUCT 小店商品
	SubjectObjType_SHOP_PRODUCT SubjectObjType = 8
	// SubjectObjType_TOPIC_PAGE 话题页
	SubjectObjType_TOPIC_PAGE SubjectObjType = 9
)
