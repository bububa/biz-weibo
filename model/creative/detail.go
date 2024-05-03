package creative

import (
	"github.com/bububa/biz-weibo/enum"
	"github.com/bububa/biz-weibo/model/ad"
	"github.com/bububa/biz-weibo/model/campaign"
)

// CreativeDetail 创意
type CreativeDetail struct {
	// ID 创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// ConfiguredStatus 开关状态 详见【附录-创意开关状态编码】
	ConfiguredStatus enum.ConfiguredStatus `json:"configured_status,omitempty"`
	// EffectiveStatus 创意状态 【附录-创意状态编码】
	EffectiveStatus enum.CreativeEffectiveStatus `json:"effective_status,omitempty"`
	// Version 创意版本，3（v3创意）、4（v4创意）
	Version int `json:"version,omitempty"`
	// CreatedAt 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Name 创意名称
	Name string `json:"name,omitempty"`
	// CreativeType 创意类型：1 新建创意、2 复用已有
	CreativeType int `json:"creative_type,omitempty"`
	// MediaMidComplex 复用类型：1 素材复用 、2 博文复用、3聚宝盆素材复用
	MediaMidComplex int `json:"media_mid_complex,omitempty"`
	// MediaMidType 博文类型：1 已有博文、 2 聚宝盆代投
	MediaMidType int `json:"media_mid_type,omitempty"`
	// Mid 博文复用时博文的mid
	Mid string `json:"mid,omitempty"`
	// ComponentMountSwitch 组件挂载开关：0关闭，1开启 当所选聚宝盆代投博文为标的支持的样式（详见【附录-各标的下支持的创意样式】），且博文不包含组件时（详见【附录-各标的样式下支持的组件类型】），支持组件挂载
	ComponentMountSwitch *int `json:"component_mount_switch,omitempty"`
	// ShortlinkText 短链文案
	ShortlinkText string `json:"shortlink_text,omitempty"`
	// Media media对象
	Media *Media `json:"media,omitempty"`
	// BasicComponents 创意基础组件
	BasicComponents *BasicComponent `json:"basic_components,omitempty"`
	// ExtendComponents 附加组件，当使用到各组件key时，对应组件参数必填
	ExtendComponents *ExtendComponent `json:"extend_components,omitempty"`
	// ContainterType 打底博文类型，1 新建 2 已有
	ContainterType int `json:"containter_type,omitempty"`
	// SocialAssetsShare 二次传播：1 聚合沉淀， 0 独立积累
	SocialAssetsShare *int `json:"social_assets_share,omitempty"`
	// ContainerMid 已有打底博文的mid
	ContainerMid string `json:"container_mid,omitempty"`
	// ContainerMedia 打底博文对象，container_type为1时有效，如果不传则默认随机选取一组素材
	ContainerMedia *ContainerMedia `json:"container_media,omitempty"`
	// PublishType 发博类型：1 立即发博 2 定时发博
	PublishType int `json:"publish_type,omitempty"`
	// PublishDatetime 定时发博时间, 精确到分，须为当前时间两小时之后，格式："2020-06-12 12:50"
	PublishDatetime string `json:"publish_datetime,omitempty"`
	// AdOnly 是否为广告私博文，0否，1是
	AdOnly *int `json:"ad_only,omitempty"`
	// ShareScope 分享范围，container_type为1时有效，0 公开可见 1仅自己可见，默认为0
	ShareScope *int `json:"share_scope,omitempty"`
	// ComponentScope 评论设置，0：所有人可评论，4：仅自己可评论
	ComponentScope *int `json:"component_scope,omitempty"`
	// ComponentTop 是否评论置顶，false否，true是
	ComponentType bool `json:"component_type,omitempty"`
	// ComponentContent 评论置顶内容
	ComponentContent string `json:"component_content,omitempty"`
	// ComponentImage 评论图片id
	ComponentImage string `json:"component_image,omitempty"`
	// ComponentLinkText 评论链接文案，必须与shortlink_text保持一致
	ComponentLinkText string `json:"component_link_text,omitempty"`
	// ComponentUserID 置顶评论发评人
	ComponentUserID uint64 `json:"component_user_id,omitempty"`
	// ComponentID 已有置顶评论id
	ComponentID uint64 `json:"component_id,omitempty"`
	// ComponentLink 置顶评论链接文案
	ComponentLink string `json:"component_link,omitempty"`
	// ComponentImageList 原生广告置顶评论图片id
	ComponentImageList []string `json:"component_image_list,omitempty"`
	// Licenses 资质id （mid不是免审的需要上传资质），资质id，示例：["370074", "360330", "302098", "290348"]
	Licenses []string `json:"licenses,omitempty"`
	// MonitorType 监测类型，0：s2s，1：c2s
	MonitorType *int `json:"monitor_type,omitempty"`
	// ThirdPartyShow1 曝光监测地址1
	ThirdPartyShow1 string `json:"third_party_show_1,omitempty"`
	// ThirdPartyClick1 点击监测地址1
	ThirdPartyClick1 string `json:"third_party_click_1,omitempty"`
	// ThirdPartyShow2 曝光监测地址2
	ThirdPartyShow2 string `json:"third_party_show_2,omitempty"`
	// ThirdPartyClick2 点击监测地址2
	ThirdPartyClick2 string `json:"third_party_click_2,omitempty"`
	// VisibilityMonitorType 可见性监测类型：0 图文，1 视频
	VisibilityMonitorType *int `json:"visibility_monitor_type,omitempty"`
	// VisibilityMonitorURL 可见性监测地址
	VisibilityMonitorURL string `json:"visibility_monitor_url,omitempty"`
	// GameAppointmentURL 游戏预约监测链接地址
	GameAppointmentURL string `json:"game_appointment_url,omitempty"`
	// VideoMonitorList 视频播放时长监测，最多支持4组点位监测;示例：[{"video_min_duration":120,"video_url":"https://ad.weibo.com"}]
	VideoMonitorList []VideoMonitor `json:"video_monitor_list,omitempty"`
	// BillingType 出价方式 详见【附录-出价方式编码】
	BillingType enum.BillingType `json:"billing_type"`
	// 营销目标 详见【附录-营销目标编码】
	Objective enum.Objective `json:"objective"`
	// PreviewMidURL 创意预览url
	PreviewMidURL string `json:"preview_mid_url"`
	// Campaign 系列对象
	Campaign *campaign.Campaign `json:"campaign"`
	// Ad 计划对象
	Ad *ad.Ad `json:"ad"`
}

// Media media对象
type Media struct {
	// Images 图片数组的数组，示例：[["9b7f515dly1fyta5o6kdnj20ic09675e","9b7f515dly1fyta5o6kdnj20ic09675e"],["9b7f515dly1fyta5o6kdnj20ic09675e"]]
	Images [][]string `json:"images,omitempty"`
	// Video 视频对象数组
	Video []Video `json:"video,omitempty"`
	// Content 博文内容对象数组
	Content []MediaContent `json:"content,omitempty"`
}

// Video 视频对象
type Video struct {
	// Source 1 仅限超粉后台使用（ADS-API不可用） 2在线视频 3资产库(ADS-API使用)
	Source int `json:"source,omitempty"`
	// Cover 视频封面
	Cover string `json:"cover,omitempty"`
	// URL 视频url
	URL string `json:"url,omitempty"`
	// Fid 视频fid
	Fid string `json:"fid,omitempty"`
	// Orientation 视频横版、竖版 详见【附录-素材-视频版式编码】
	Orientation enum.Orientation `json:"orientation,omitempty"`
}

// MediaContent 博文内容对象
type MediaContent struct {
	// ActualContent 博文正文内容
	ActualContent string `json:"actual_content,omitempty"`
}

// BasicComponent 创意基础组件
type BasicComponent struct {
	// SideXComponent 边X边X
	SideXComponent *SideXComponent `json:"side_x_component,omitempty"`
	// ButtonComponent 按钮
	ButtonComponent *ButtonComponent `json:"button_component,omitempty"`
	// CardComponent 卡片
	CardComponent *CardComponent `json:"card_component,omitempty"`
	// ClickDirectJumpComponent 视图直跳 示例：{"component_id":1508,"direct_name":"立即咨询"}
	ClickDirectJumpComponent *ClickDirectJumpComponent `json:"click_direct_jump_component,omitempty"`
}

// ComponentBase
type ComponentBase struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
}

// SideXComponent 边X边X
type SideXComponent struct {
	// SideXDownload 边X边下，示例：{"component_id":1015,"lingdong_component_id":"333"}
	SideXDownload *SideX `json:"side_x_download,omitempty"`
	// SideXH5 边X边H5，示例：{"component_id":1014}
	SideX5H5 *SideX `json:"side_x_h5,omitempty"`
}

// SideX 边X边下，示例：{"component_id":1015,"lingdong_component_id":"333"}
type SideX struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// LingdongComponentID 灵动落地页ID
	LingdongComponentID int `json:"lingdong_component_id,omitempty"`
}

// ButtonComponent 按钮
type ButtonComponent struct {
	// Normal 普适，示例：{"component_id":1005,"button":"立即下载"}
	Normal *Button `json:"normal,omitempty"`
	// LongImage 长图，示例：{"component_id":60,"button":"立即参与"}
	LongImage *Button `json:"long_image,omitempty"`
	// Banner banner，示例：{"component_id":1010,"button":"立即购买"}
	Banner *Button `json:"banner,omitempty"`
}

// Button 按钮
type Button struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// Button 按钮文案
	Button string `json:"button,omitempty"`
}

// CardComponent 卡片
type CardComponent struct {
	// CardTitle 卡片标题，示例：{"component_id":1003,"card_title":"xxxx"}
	CardTitle *Card `json:"card_title,omitempty"`
	// CardDescription 卡片副标题，示例：{"component_id":1004,"card_description":"111"}
	CardDescription *Card `json:"card_description,omitempty"`
}

// Card 卡片
type Card struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// CardTitle  按钮文案
	CardTitle string `json:"card_title,omitempty"`
	// CardDescription 卡片副标题文案
	CardDescription string `json:"card_description,omitempty"`
}

// ClickDirectJumpComponent 视图直跳 示例：{"component_id":1508,"direct_name":"立即咨询"}
type ClickDirectJumpComponent struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// DirectName 视图直跳文案，详见【附录-视图直跳文案】
	DirectName string `json:"direct_name,omitempty"`
}

// ExtendComponent 附加组件，当使用到各组件key时，对应组件参数必填
type ExtendComponent struct {
	// BlueBarComponent 蓝条
	BlueBarComponent *ComponentBase `json:"blue_bar_component,omitempty"`
	// GestureComponent 手势互动
	GestureComponent *GestureComponent `json:"gesture_component,omitempty"`
	// GoodsWindowComponent 商品橱窗
	GoodsWindowComponent *ComponentBase `json:"goods_window_component,omitempty"`
	// VideoPopupComponent 视频弹窗
	VideoPopupComponent *VideoPopupComponent `json:"video_popup_component,omitempty"`
	// PicturePopupComponent 图片弹窗
	PicturePopupComponent *PicturePopupComponent `json:"picture_popup_component,omitempty"`
	// OnlineConsultComponent 在线咨询
	OnlineConsultComponent *OnlineConsultComponent `json:"online_consult_component,omitempty"`
	// MessageTipsComponent 私信预填文案
	MessageTipsComponent *MessageTipsComponent `json:"message_tips_component,omitempty"`
	// HotContentComponent 热门内容 示例：{"component_id":1023,"nick_name_show":1,"nick_name_uid":2608812381,"nick_name":"用户昵称","hot_content_text":"热门内容的文案"}
	HotContentComponent *HotContentComponent `json:"hot_content_component,omitempty"`
}

// GestureComponent 手势互动
type GestureComponent struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// GestureType 手势类型，1-长按， 2-横划
	GestureType int `json:"gesture_type,omitempty"`
	// GestureDisplayStartTime 进入时间，最小3，最大12
	GestureDisplayStartTime int `json:"gesture_display_start_time,omitempty"`
	// GestureDurationTime 持续时间，最小2，最大10
	GestureDurationTime int `json:"gesture_duration_time,omitempty"`
}

// VideoPopupComponent 视频弹窗
type VideoPopupComponent struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// VideoPopupType 1-小图弹窗，2-整图弹窗
	VideoPopupType int `json:"video_popup_type,omitempty"`
	// VideoPopupImage 弹窗图片pid
	VideoPopupImage string `json:"video_popup_image,omitempty"`
	// VideoPopupImageType 1自定义 3模板
	VideoPopupImageType int `json:"video_popup_image_type,omitempty"`
	// VideoPopupTitle 弹窗标题
	VideoPopupTitle string `json:"video_popup_title,omitempty"`
	// VideoPopupDescription 弹窗描述
	VideoPopupDescription string `json:"video_popup_description,omitempty"`
}

// PicturePopupComponent 图片弹窗
type PicturePopupComponent struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// PicturePopupImage 弹窗图片pid
	PicturePopupImage string `json:"picture_popup_image,omitempty"`
	// PicturePopupTItle 弹窗标题
	PicturePopupTItle string `json:"picture_popup_t_itle,omitempty"`
	// PicuturePopupPrice 弹窗价格
	PicuturePopupPrice string `json:"picuture_popup_price,omitempty"`
}

// OnlineConsultComponent 在线咨询
type OnlineConsultComponent struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// LingdongKuuid  灵动私信页id
	LingdongKuuid string `json:"lingdong_kuuid,omitempty"`
}

// MessageTipsComponent 私信预填文案
type MessageTipsComponent struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// MessageTipsText 预填文案
	MessageTipsText string `json:"message_tips_text,omitempty"`
}

// ContainerMedia 打底博文对象，container_type为1时有效，如果不传则默认随机选取一组素材
type ContainerMedia struct {
	// ContainerImage  pid数组，示例：["aaaa","ccccc"]
	ContainerImage []string `json:"container_image,omitempty"`
	// ContainerVideo 打底博文视频对象
	ContainerVideo *Video `json:"container_video,omitempty"`
	// ContainerContent 打底博文内容
	ContainerContent *MediaContent `json:"container_content,omitempty"`
}

// VideoMonitor 视频播放时长监测
type VideoMonitor struct {
	// VideoMinDuration 视频播放监测点位
	VideoMinDuration int64 `json:"video_min_duration,omitempty"`
	// VideoURL 视频播放时长监测地址
	VideoURL string `json:"video_url,omitempty"`
}

// HotContentComponent 热门内容
type HotContentComponent struct {
	// ComponentID 组件id 详见【附录-组件编码】
	ComponentID int `json:"component_id,omitempty"`
	// NickNameShow 昵称设置，1：展示 0：隐藏
	NickNameShow *int `json:"nick_name_show,omitempty"`
	// NickNameUid 昵称uid，当前广告主或同资质同通路uid。【广告投放-创意管理-获取发评人列表】 说明：nick_name_show=0时，不需要设置这个字段
	NickNameUid uint64 `json:"nick_name_uid,omitempty"`
	// NickName nick_name_uid对应的昵称，说明：nick_name_show=0时，不需要设置这个字段
	NickName string `json:"nick_name,omitempty"`
	// HotContentText 文案内容
	HotContentText string `json:"hot_content_text,omitempty"`
}
