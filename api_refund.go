package jeepay_go_sdk

import (
	_context "context"
	_nethttp "net/http"
)

type RefundApiService Service

func (p *RefundApiService) RefundOrder(ctx _context.Context, query RefundRequest) (Response[RefundResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/refund/refundOrder", PayModel: model}
	return postExecute[Response[RefundResponse]](request)
}

func (p *RefundApiService) QueryRefundOrder(ctx _context.Context, query RefundQueryRequest) (Response[RefundQueryResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/refund/query", PayModel: model}
	return postExecute[Response[RefundQueryResponse]](request)
}

func (p *RefundApiService) RefundNotifyOrder(ctx _context.Context, query RefundRequest) (RefundQueryResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/refund/query", PayModel: model}
	return postExecute[RefundQueryResponse](request)
}
