package jeepay_go_sdk

type RefundRequest struct {
	PayOrderId   string `json:"payOrderId"`
	ExtParam     string `json:"extParam"`
	MchOrderNo   string `json:"mchOrderNo"`
	RefundReason string `json:"refundReason"`
	Sign         string `json:"sign"`
	ReqTime      string `json:"reqTime"`
	Version      string `json:"version"`
	ChannelExtra string `json:"channelExtra"`
	AppId        string `json:"appId"`
	MchRefundNo  string `json:"mchRefundNo"`
	ClientIp     string `json:"clientIp"`
	NotifyUrl    string `json:"notifyUrl"`
	SignType     string `json:"signType"`
	Currency     string `json:"currency"`
	MchNo        string `json:"mchNo"`
	RefundAmount int    `json:"refundAmount"`
}

type RefundResponse struct {
	ChannelOrderNo string `json:"channelOrderNo"`
	MchRefundNo    string `json:"mchRefundNo"`
	PayAmount      int    `json:"payAmount"`
	RefundAmount   int    `json:"refundAmount"`
	RefundOrderId  string `json:"refundOrderId"`
	State          int    `json:"state"`
}

type RefundQueryRequest struct {
	RefundOrderId string `json:"refundOrderId"`
	AppId         string `json:"appId"`
	Sign          string `json:"sign"`
	SignType      string `json:"signType"`
	ReqTime       string `json:"reqTime"`
	MchNo         string `json:"mchNo"`
	Version       string `json:"version"`
}

type RefundQueryResponse struct {
	AppId          string `json:"appId"`
	ChannelOrderNo string `json:"channelOrderNo"`
	CreatedAt      int64  `json:"createdAt"`
	Currency       string `json:"currency"`
	ExtParam       string `json:"extParam"`
	MchNo          string `json:"mchNo"`
	MchRefundNo    string `json:"mchRefundNo"`
	PayAmount      int    `json:"payAmount"`
	PayOrderId     string `json:"payOrderId"`
	RefundAmount   int    `json:"refundAmount"`
	RefundOrderId  string `json:"refundOrderId"`
	State          int    `json:"state"`
	SuccessTime    int64  `json:"successTime"`
}
