package jeepay_go_sdk

import (
	"github.com/finecloud/jeepay-sdk-go/api"
	"log"
	"net/http"
	"net/http/httputil"
)

type APIClient struct {
	cfg *Configuration

	common Service

	PayApi *api.PayApiService

	RefundApi *api.RefundApiService

	SubAccountApi *api.SubAccountApiService

	TransferApi *api.TransferApiService
}

type Service struct {
	client *APIClient
}

// NewApiClient returns a new instance of the JeepayClient client.
func NewApiClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	c.PayApi = (*api.PayApiService)(&c.common)
	c.RefundApi = (*api.RefundApiService)(&c.common)
	c.SubAccountApi = (*api.SubAccountApiService)(&c.common)
	c.TransferApi = (*api.TransferApiService)(&c.common)

	return c

}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}
