package jeepay_go_sdk

type PayCreateRequest struct {
	Amount       *int     `json:"amount"`
	AppId        *string  `json:"appId"`
	Body         *string  `json:"body"`
	ClientIp     *string  `json:"clientIp"`
	Currency     *string  `json:"currency"`
	ExtParam     *string  `json:"extParam"`
	MchOrderNo   *string  `json:"mchOrderNo"`
	MchNo        *string  `json:"mchNo"`
	Subject      *string  `json:"subject"`
	WayCode      *WayCode `json:"wayCode"`
	Sign         *string  `json:"sign"`
	ReqTime      *int64   `json:"reqTime"`
	Version      *string  `json:"version"`
	ChannelExtra *string  `json:"channelExtra"`
	NotifyUrl    *string  `json:"notifyUrl"`
	SignType     *string  `json:"signType"`
	ReturnUrl    *string  `json:"returnUrl"`
	DivisionMode *int     `json:"divisionMode"`
}
type PayQueryItem struct {
	Amount         *int   `json:"amount"`
	AppId          string `json:"appId"`
	Body           string `json:"body"`
	ChannelOrderNo string `json:"channelOrderNo"`
	ClientIp       string `json:"clientIp"`
	CreatedAt      *int64 `json:"createdAt"`
	Currency       string `json:"currency"`
	ExtParam       string `json:"extParam"`
	IfCode         string `json:"ifCode"`
	MchNo          string `json:"mchNo"`
	MchOrderNo     string `json:"mchOrderNo"`
	PayOrderId     string `json:"payOrderId"`
	State          *int   `json:"state"`
	Subject        string `json:"subject"`
	SuccessTime    int64  `json:"successTime"`
	WayCode        string `json:"wayCode"`
}

type BaseResponse struct {
	ErrCode     string `json:"errCode"`
	ErrMsg      string `json:"errMsg"`
	MchOrderNo  string `json:"mchOrderNo"`
	OrderState  *int   `json:"orderState"`
	PayOrderId  string `json:"payOrderId"`
	PayDataType string `json:"payDataType"`
	PayData     string `json:"payData"`
}

type PayChannelExtra struct {
	// 当 wayCode=ALI_JSAPI 时，channelExtra必须传buyerUserId，为支付宝用户ID，channelExtra示例数据如：
	BuyerUserId *string `json:"buyerUserId"`
	// 当 wayCode=AUTO_BAR 或 wayCode=ALI_BAR 或 wayCode=WX_BAR 或 wayCode=YSF_BAR 时，channelExtra必须传auth_code，为用户的付款码值，channelExtra示例数据如：
	AuthCode *string `json:"auth_code"`
	// 当 wayCode=WX_JSAPI 或 wayCode=WX_LITE 时，channelExtra必须传openid，为支付宝用户ID，channelExtra示例数据如：
	Openid *string `json:"openid"`
	// 当 wayCode=ALI_WAP 时，channelExtra可以传payDataType设置返回支付数据支付类型。此时payDataType可以为：form-返回自动跳转的支付表单,codeImgUrl-返回一个二维码图片URL,payUrl-返回支付链接，不传payDataType默认返回payUrl类型, channelExtra示例数据如：
	// 当 wayCode=ALI_PC 时，channelExtra可以传payDataType设置返回支付数据支付类型。此时payDataType可以为：form-返回自动跳转的支付表单,payUrl-返回支付链接，不传payDataType默认返回payUrl类型, channelExtra示例数据如：
	PayDataType *string `json:"payDataType"`
}

type PayQueryRequest struct {
	AppId      *string `json:"appId"`
	MchOrderNo *string `json:"mchOrderNo"`
	MchNo      *string `json:"mchNo"`
	Sign       *string `json:"sign"`
	ReqTime    *int64  `json:"reqTime"`
	Version    *string `json:"version"`
	SignType   *string `json:"signType"`
}

func NewPayModel() *PayCreateRequest {
	this := PayCreateRequest{}
	return &this
}

type OrderCloseRequest struct {
	PayOrderId string `json:"payOrderId"`
	AppId      string `json:"appId"`
	Sign       string `json:"sign"`
	SignType   string `json:"signType"`
	ReqTime    string `json:"reqTime"`
	MchNo      string `json:"mchNo"`
	Version    string `json:"version"`
}

type OrderNotifyRequest struct {
	Amount     int    `json:"amount"`
	Body       string `json:"body"`
	ClientIp   string `json:"clientIp"`
	CreatedAt  string `json:"createdAt"`
	Currency   string `json:"currency"`
	ExtParam   string `json:"extParam"`
	IfCode     string `json:"ifCode"`
	MchNo      string `json:"mchNo"`
	AppId      string `json:"appId"`
	MchOrderNo string `json:"mchOrderNo"`
	PayOrderId string `json:"payOrderId"`
	State      int    `json:"state"`
	Subject    string `json:"subject"`
	WayCode    string `json:"wayCode"`
	Sign       string `json:"sign"`
}

type OrderChannelRequest struct {
	MchNo       string `json:"mchNo"`
	AppId       string `json:"appId"`
	IfCode      string `json:"ifCode"`
	RedirectUrl string `json:"redirectUrl"`
	Sign        string `json:"sign"`
	SignType    string `json:"signType"`
	ReqTime     string `json:"reqTime"`
	Version     string `json:"version"`
}

type OrderChannelResponse struct {
	MchNo       string `json:"mchNo"`
	AppId       string `json:"appId"`
	IfCode      string `json:"ifCode"`
	RedirectUrl string `json:"redirectUrl"`
	Sign        string `json:"sign"`
	SignType    string `json:"signType"`
	ReqTime     string `json:"reqTime"`
	Version     string `json:"version"`
}

type OrderNotifyResponse struct {
	Amount     int    `json:"amount"`
	Body       string `json:"body"`
	ClientIp   string `json:"clientIp"`
	CreatedAt  string `json:"createdAt"`
	Currency   string `json:"currency"`
	ExtParam   string `json:"extParam"`
	IfCode     string `json:"ifCode"`
	MchNo      string `json:"mchNo"`
	AppId      string `json:"appId"`
	MchOrderNo string `json:"mchOrderNo"`
	PayOrderId string `json:"payOrderId"`
	State      int    `json:"state"`
	Subject    string `json:"subject"`
	WayCode    string `json:"wayCode"`
	Sign       string `json:"sign"`
}
