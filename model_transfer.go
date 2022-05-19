package jeepay_go_sdk

type TransferExecRequest struct {
	IfCode       string `json:"ifCode"`
	EntryType    string `json:"entryType"`
	Amount       int    `json:"amount"`
	AccountName  string `json:"accountName"`
	MchOrderNo   string `json:"mchOrderNo"`
	Sign         string `json:"sign"`
	TransferDesc string `json:"transferDesc"`
	ReqTime      string `json:"reqTime"`
	Version      string `json:"version"`
	AppId        string `json:"appId"`
	AccountNo    string `json:"accountNo"`
	ClientIp     string `json:"clientIp"`
	SignType     string `json:"signType"`
	Currency     string `json:"currency"`
	MchNo        string `json:"mchNo"`
}

type TransferExecResponse struct {
	AccountNo      string `json:"accountNo"`
	Amount         int    `json:"amount"`
	ChannelOrderNo string `json:"channelOrderNo"`
	MchOrderNo     string `json:"mchOrderNo"`
	State          int    `json:"state"`
	TransferId     string `json:"transferId"`
}

type QueryTransferRequest struct {
	AppId      string `json:"appId"`
	Sign       string `json:"sign"`
	SignType   string `json:"signType"`
	ReqTime    string `json:"reqTime"`
	TransferId string `json:"transferId"`
	MchNo      string `json:"mchNo"`
	Version    string `json:"version"`
}

type QueryTransferResponse struct {
	AccountNo    string `json:"accountNo"`
	Amount       int    `json:"amount"`
	AppId        string `json:"appId"`
	CreatedAt    int64  `json:"createdAt"`
	Currency     string `json:"currency"`
	EntryType    string `json:"entryType"`
	ErrCode      string `json:"errCode"`
	ErrMsg       string `json:"errMsg"`
	IfCode       string `json:"ifCode"`
	MchNo        string `json:"mchNo"`
	MchOrderNo   string `json:"mchOrderNo"`
	State        int    `json:"state"`
	TransferDesc string `json:"transferDesc"`
	TransferId   string `json:"transferId"`
}

type NotifyTransferResponse struct {
	AccountNo    string `json:"accountNo"`
	Amount       int    `json:"amount"`
	AppId        string `json:"appId"`
	CreatedAt    int64  `json:"createdAt"`
	Currency     string `json:"currency"`
	EntryType    string `json:"entryType"`
	ErrCode      string `json:"errCode"`
	ErrMsg       string `json:"errMsg"`
	IfCode       string `json:"ifCode"`
	MchNo        string `json:"mchNo"`
	MchOrderNo   string `json:"mchOrderNo"`
	State        int    `json:"state"`
	TransferDesc string `json:"transferDesc"`
	TransferId   string `json:"transferId"`
}
