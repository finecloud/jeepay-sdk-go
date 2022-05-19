package jeepay_go_sdk

type DivisionBindRequest struct {
	Version          string `json:"version"`
	ReqTime          string `json:"reqTime"`
	SignType         string `json:"signType"`
	Sign             string `json:"sign"`
	MchNo            string `json:"mchNo"`
	AppId            string `json:"appId"`
	IfCode           string `json:"ifCode"`
	ReceiverAlias    string `json:"receiverAlias"`
	ReceiverGroupId  string `json:"receiverGroupId"`
	AccType          string `json:"accType"`
	AccNo            string `json:"accNo"`
	AccName          string `json:"accName"`
	RelationType     string `json:"relationType"`
	RelationTypeName string `json:"relationTypeName"`
	DivisionProfit   string `json:"divisionProfit"`
}

type DivisionBindResponse struct {
	AccName          string  `json:"accName"`
	AccNo            string  `json:"accNo"`
	AccType          int     `json:"accType"`
	AppId            string  `json:"appId"`
	BindState        int     `json:"bindState"`
	DivisionProfit   float64 `json:"divisionProfit"`
	ErrCode          string  `json:"errCode"`
	ErrMsg           string  `json:"errMsg"`
	IfCode           string  `json:"ifCode"`
	MchNo            string  `json:"mchNo"`
	ReceiverAlias    string  `json:"receiverAlias"`
	ReceiverGroupId  int     `json:"receiverGroupId"`
	RelationType     string  `json:"relationType"`
	RelationTypeName string  `json:"relationTypeName"`
}

type DivisionExecRequest struct {
	Version                     string `json:"version"`
	ReqTime                     string `json:"reqTime"`
	SignType                    string `json:"signType"`
	Sign                        string `json:"sign"`
	MchNo                       string `json:"mchNo"`
	AppId                       string `json:"appId"`
	PayOrderId                  string `json:"payOrderId"`
	UseSysAutoDivisionReceivers string `json:"useSysAutoDivisionReceivers"`
	Receivers                   string `json:"receivers"`
}
