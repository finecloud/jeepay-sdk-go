package jeepay_go_sdk

type PayModel struct {
	Amount         *int    `json:"amount"`
	AppId          *string `json:"appId"`
	Body           *string `json:"body"`
	ChannelOrderNo *string `json:"channelOrderNo"`
	ClientIp       *string `json:"clientIp"`
	CreatedAt      *int64  `json:"createdAt"`
	Currency       *string `json:"currency"`
	ExtParam       *string `json:"extParam"`
	IfCode         *string `json:"ifCode"`
	MchNo          *string `json:"mchNo"`
	MchOrderNo     *string `json:"mchOrderNo"`
	PayOrderId     *string `json:"payOrderId"`
	State          *int    `json:"state"`
	Subject        *string `json:"subject"`
	SuccessTime    *int64  `json:"successTime"`
	WayCode        *string `json:"wayCode"`
	Sign           *string `json:"sign"`
	ReqTime        *int64  `json:"reqTime"`
	Version        *string `json:"version"`
	ChannelExtra   *string `json:"channelExtra"`
	NotifyUrl      *string `json:"notifyUrl"`
	SignType       *string `json:"signType"`
	ReturnUrl      *string `json:"returnUrl"`
	DivisionMode   *int    `json:"divisionMode"`
}

func NewPayModel() *PayModel {
	this := PayModel{}
	return &this
}
