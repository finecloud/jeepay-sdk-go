package api

import jeepaygosdk "github.com/finecloud/jeepay-sdk-go"
import _context "context"

type PayApiService jeepaygosdk.Service

type UnifiedOrderCreateRequest struct {
	ctx          _context.Context
	ApiService   *PayApiService
	Amount       int    `json:"amount"`
	ExtParam     string `json:"extParam"`
	MchOrderNo   string `json:"mchOrderNo"`
	Subject      string `json:"subject"`
	WayCode      string `json:"wayCode"`
	Sign         string `json:"sign"`
	ReqTime      string `json:"reqTime"`
	Body         string `json:"body"`
	Version      string `json:"version"`
	ChannelExtra string `json:"channelExtra"`
	AppId        string `json:"appId"`
	ClientIp     string `json:"clientIp"`
	NotifyUrl    string `json:"notifyUrl"`
	SignType     string `json:"signType"`
	Currency     string `json:"currency"`
	ReturnUrl    string `json:"returnUrl"`
	MchNo        string `json:"mchNo"`
	DivisionMode int    `json:"divisionMode"`
}

// UnifiedOrder Unified order API  [https://docs.jeequan.com/docs/jeepay/payment_api]
func (p *PayApiService) UnifiedOrder(ctx _context.Context) UnifiedOrderCreateRequest {

	return UnifiedOrderCreateRequest{ctx: ctx, ApiService: p}
}
