package jeepay_go_sdk

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"time"
)

type BaseRequest[T any] struct {
	Amount       *int     `json:"amount"`
	AppId        *string  `json:"appId"`
	ClientIp     *string  `json:"clientIp"`
	ExtParam     *string  `json:"extParam"`
	MchOrderNo   *string  `json:"mchOrderNo"`
	MchNo        *string  `json:"mchNo"`
	Subject      *string  `json:"subject"`
	WayCode      *WayCode `json:"wayCode"`
	Sign         *string  `json:"sign"`
	ReqTime      *int64   `json:"reqTime"`
	Version      *string  `json:"version"`
	ChannelExtra *string  `json:"channelExtra"`
	SignType     *string  `json:"signType"`
	DivisionMode *int     `json:"divisionMode"`
}

type ApiRequest struct {
	ctx        _context.Context
	ApiService *Service
	PayModel   map[string]interface{}
	Path       string
}

// postExecute executes the request and returns response or error.
func postExecute[R any](request ApiRequest) (R, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  R
	)
	if request.PayModel == nil {
		return localVarReturnValue, nil, reportError("dAGRun is required and must be specified")
	}

	var p = request.ApiService
	baseUrl, err := p.client.cfg.GetLocalBasePath()
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarPath := baseUrl + request.Path

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

	service := request.ApiService
	if service == nil {
		return localVarReturnValue, nil, reportError("ApiService is required and must be specified")
	}

	request.PayModel["signType"] = "MD5"
	request.PayModel["version"] = "1.0"
	request.PayModel["reqTime"] = time.Now().UnixMicro()
	request.PayModel["appId"] = p.Configuration.ApiId
	//mapData := Struct2MapName(request.PayModel)
	sign, err := encrypt(request.PayModel, p.Configuration.ApiKey, p.Configuration.SignType)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	request.PayModel["sign"] = sign

	// body params
	localVarPostBody = request.PayModel

	req, err := p.client.prepareRequest(request.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
