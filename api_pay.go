package jeepay_go_sdk

import (
	"bytes"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"time"
)
import _context "context"

type PayApiService Service

type OrderCreateRequest struct {
	ctx        _context.Context
	ApiService *PayApiService
	PayModel   *PayModel
}

func (r OrderCreateRequest) Execute() (PayModel, *_nethttp.Response, error) {
	return r.ApiService.PostOrderExecute(r)
}

// CreateOrder Unified order API  [https://docs.jeequan.com/docs/jeepay/payment_api]
func (p *PayApiService) CreateOrder(ctx _context.Context) OrderCreateRequest {
	return OrderCreateRequest{ctx: ctx, ApiService: p}
}

func (p *PayApiService) PostOrderExecute(r OrderCreateRequest) (PayModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PayModel
	)
	if r.PayModel == nil {
		return localVarReturnValue, nil, reportError("dAGRun is required and must be specified")
	}

	path, err := p.client.cfg.GetLocalBasePath()
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarPath := path + "/api/pay/unifiedOrder"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	signType := "MD5"
	r.PayModel.SignType = &signType
	version := "1.0"
	r.PayModel.Version = &version
	micro := time.Now().UnixMicro()
	r.PayModel.ReqTime = &micro
	id := p.Configuration.ApiId
	r.PayModel.AppId = &id
	mapData := Struct2MapName(r.PayModel)
	sign, err := encrypt(mapData, p.Configuration.ApiKey, p.Configuration.SignType)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	r.PayModel.Sign = &sign

	// body params
	localVarPostBody = r.PayModel

	req, err := p.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := p.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = p.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = p.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = p.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = p.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = p.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	err = p.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrderCheckRequest struct {
	PayOrderId string `json:"payOrderId"`
	AppId      string `json:"appId"`
	Sign       string `json:"sign"`
	SignType   string `json:"signType"`
	ReqTime    string `json:"reqTime"`
	MchNo      string `json:"mchNo"`
	Version    string `json:"version"`
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
