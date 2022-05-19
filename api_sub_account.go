package jeepay_go_sdk

import (
	_context "context"
	_nethttp "net/http"
)

type SubAccountApiService Service

func (p *SubAccountApiService) DivisionBindReceiver(ctx _context.Context, query DivisionBindRequest) (DivisionBindResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/division/receiver/bind", PayModel: model}
	return postExecute[DivisionBindResponse](request)
}

func (p *SubAccountApiService) DivisionExec(ctx _context.Context, query DivisionExecRequest) (Response[BaseResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/division/exec", PayModel: model}
	return postExecute[Response[BaseResponse]](request)
}
