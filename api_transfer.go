package jeepay_go_sdk

import (
	_context "context"
	_nethttp "net/http"
)

type TransferApiService Service

func (p *TransferApiService) TransferOrder(ctx _context.Context, query TransferExecRequest) (Response[TransferExecResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/transferOrder", PayModel: model}
	return postExecute[Response[TransferExecResponse]](request)
}

func (p *TransferApiService) QueryTransferOrder(ctx _context.Context, query QueryTransferRequest) (Response[QueryTransferResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/transfer/query", PayModel: model}
	return postExecute[Response[QueryTransferResponse]](request)
}

func (p *TransferApiService) NotifyTransferOrder(ctx _context.Context, query RefundRequest, path string) (NotifyTransferResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: path, PayModel: model}
	return postExecute[NotifyTransferResponse](request)
}
