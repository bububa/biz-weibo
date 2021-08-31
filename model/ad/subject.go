package ad

// Subject 标的对象
type Subject struct {
	// ObjType 注此处创意层特殊逻辑：根据direct_type，obj_type51为app直达，52为电商直达
	ObjType int `json:"obj_type,omitempty"`
	// ObjURL 落地页地址
	ObjURL string `json:"obj_url,omitempty"`
	// ObjURLExtendID 落地页地址id
	ObjURLExtendID int64 `json:"obj_url_extend_id,omitempty"`
	// ObjURLExtendType 落地页类型 1cpl 2灵动
	ObjURLExtendType int `json:"obj_url_extend_type,omitempty"`
	// ObjURLType 落地页地址类型，后端用到，收集销售线索表单提交时创意下发。默认空，cpl、luming（灵动）、land_page（即js）、api（线索api）
	ObjURLType string `json:"obj_url_type,omitempty"`
	// AppType app类型,默认-1，0 Android,1 ios,
	AppType int `json:"app_type,omitempty"`
	// AppID app id
	AppID int64 `json:"app_id,omitempty"`
	// IsAppStoreDirect 是否开启应用商店直达，0不开启，1开启
	IsAppStoreDirect int `json:"is_app_store_direct,omitempty"`
	// URLH5 h5链接地址
	URLH5 string `json:"url_h5,omitempty"`
	// AppWm 渠道号
	AppWm string `json:"app_wm,omitempty"`
	// DirectURL 直达链接地址
	DirectURL string `json:"direct_url,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
	// IsDirect 是否开启直达链接， 0不开启， 1开启
	IsDirect int `json:"is_direct,omitempty"`
	// DirectType 直达类型 0 未使用 1 app店直达 2 电商直达
	DirectType int `json:"direct_type,omitempty"`
	// ConvertMonitorType 转化监测类型 0默认无， 1. 应用api、2. js、 3. 线索api、 4微信转化api
	ConvertMonitorType int `json:"convert_monitor_type,omitempty"`
	// CollectType 归因方式 0默认无， 10000 点击， 11000 互动
	CollectType int `json:"collect_type,omitempty"`
	// MonitorAPI 监测api地址
	MonitorAPI string `json:"monitor_api,omitempty"`
	// MiniProgram 小程序信息包括：{"mini_program_id":程序原始id,"mini_program_address":"小程序地址","mini_program_name":"小程序名称"}
	MiniProgram string `json:"mini_program,omitempty"`
}
