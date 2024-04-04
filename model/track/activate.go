package track

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/biz-weibo/model"
)

// ActivateRequest 线索API数据上报 API Request
type ActivateRequest struct {
	// AdvertiserID 您管理的广告主id，为空则为token对应账户不使用请勿下发
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// Time 转化时间，如：1534730754456(毫秒)
	Time int64 `json:"time,omitempty"`
	// MarkID mark_id允许为空，此时为自然量
	MarkID string `json:"mark_id,omitempty"`
	// Behavior 客户行为码，behavior编码见下文
	Behavior Behavior `json:"behavior,omitempty"`
	// Host 回传来源，默认请填写您的域名，如：weibo.com
	Host string `json:"host,omitempty"`
	// Imp 小程序同步监测相关参数，请求唯一id，微信小程序必须有此字段，需要urlencode
	Imp string `json:"imp,omitempty"`
	// PaidAmount 小程序同步监测相关参数，behavior=3003时可选，单位：元，保留小数点后两位，精确到分。behavior行为码，详见【附录-behavior行为码】
	PaidAmount float64 `json:"paid_amount,omitempty"`
	// Score 表单提交相关参数，behavior=1001时可选
	// 意向度从高到低，分别回传5/4/3/2，四个等级。
	// 如少于四个意向度，则自己定义，例如3/2，如多于四个意向度，则最多定义四个。
	// 回传1是负向（无意向），例如选择了“不参与活动”，或直接点击关闭意向表单
	Score int `json:"score,omitempty"`
	// TechnologyProvider 技术服务商名，默认为空，使用三方技术服务时必填，未填写将影响实际结算
	TechnologyProvider string `json:"technology_provider,omitempty"`
	// ItemOrderPay 付款订单商品数相关参数，behavior编码2001必选，参数格式：item_order_pay=urlencode(json)。
	ItemOrderPay *ItemOrderPay `json:"item_order_pay,omitempty"`
}

// ItemOrderPay 付款订单商品数相关参数，behavior编码2001必选，参数格式：item_order_pay=urlencode(json)。
type ItemOrderPay struct {
	// ClickTime 点击时间，时间戳，示例"1534730754456"
	ClickTime string `json:"click_time,omitempty"`
	// OrderCreateTime 创建时间，时间戳，示例"1534730754456"
	OrderCreateTime string `json:"order_create_time,omitempty"`
	// OrderPayTime 付款时间，时间戳，示例"1534730754456"
	OrderPayTime string `json:"order_pay_time,omitempty"`
	// OrderPayClickTime 订单付款时间，时间戳，示例"1534730754456"
	OrderPayClickTime string `json:"order_pay_click_time,omitempty"`
	// ItemName 商品名称
	ItemName string `json:"item_name,omitempty"`
	// ItemPrice 商品单价 单位：元，保留小数点后两位，精确到分
	ItemPrice float64 `json:"item_price,omitempty"`
	// PaymentAmount 商品付款金额 单位：元，保留小数点后两位，精确到分
	PaymentAmount float64 `json:"payment_amount,omitempty"`
	// QuantityOfItem 商品数量
	QuantityOfItem int `json:"quantity_of_item,omitempty"`
	// Category 类目名称
	Category string `json:"category,omitempty"`
	// OrderStatus 订单状态，只支持【已付款】
	OrderStatus string `json:"order_status,omitempty"`
	// OrderID 付款订单号
	OrderID string `json:"order_id,omitempty"`
	// ServiceProvicer 服务商名称
	ServiceProvicer string `json:"service_provicer,omitempty"`
	// OrderType 1代表定金订单，2代表尾款订单，0代表常规订单
	OrderType int `json:"order_type,omitempty"`
	// IsSameItem 1代表主推商品，2代表同店跨品商品，0代表未知
	IsSameItem int `json:"is_same_item,omitempty"`
}

// URL implement Request interface
func (r ActivateRequest) URL() string {
	return "v3/track/activate"
}

// Payload implement Request interface
func (r ActivateRequest) Payload() *model.Payload {
	p := new(model.Payload)
	if r.AdvertiserID > 0 {
		p.AddValue("advertiser_id", strconv.FormatInt(r.AdvertiserID, 10))
	}
	p.AddValue("time", strconv.FormatInt(r.Time, 10))
	if r.MarkID != "" {
		p.AddValue("mark_id", r.MarkID)
	}
	p.AddValue("behavior", strconv.Itoa(int(r.Behavior)))
	if r.Host != "" {
		p.AddValue("host", r.Host)
	}
	if r.Imp != "" {
		p.AddValue("imp", r.Imp)
	}
	if r.PaidAmount > 1e-15 {
		p.AddValue("pay_amount", strconv.FormatFloat(r.PaidAmount, 'f', -1, 64))
	}
	if r.Score > 0 {
		p.AddValue("score", strconv.Itoa(r.Score))
	}
	if r.TechnologyProvider != "" {
		p.AddValue("technology_provider", r.TechnologyProvider)
	}
	if r.ItemOrderPay != nil {
		bs, _ := json.Marshal(r.ItemOrderPay)
		p.AddValue("item_order_pay", string(bs))
	}
	return p
}

type ActivateV4Request struct {
	ActivateRequest
}

// URL implement Request interface
func (r ActivateV4Request) URL() string {
	return "v4/track/activate"
}
