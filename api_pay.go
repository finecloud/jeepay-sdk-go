package jeepay_go_sdk

import (
	_nethttp "net/http"
)
import _context "context"

type PayApiService Service

func (p *PayApiService) CreateOrder(ctx _context.Context, query PayCreateRequest) (Response[BaseResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/pay/unifiedOrder", PayModel: model}
	return postExecute[Response[BaseResponse]](request)
}

func (p *PayApiService) QueryOrder(ctx _context.Context, query PayQueryRequest) (Response[PayQueryItem], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/pay/query", PayModel: model}
	return postExecute[Response[PayQueryItem]](request)
}

func (p *PayApiService) NotificationOrder(ctx _context.Context, query OrderNotifyRequest, path string) (OrderNotifyResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: path, PayModel: model}
	return postExecute[OrderNotifyResponse](request)
}

func (p *PayApiService) CloseOrder(ctx _context.Context, query OrderCloseRequest) (Response[BaseResponse], *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "/api/channelUserId/jump", PayModel: model}
	return postExecute[Response[BaseResponse]](request)
}

func (p *PayApiService) ChannelOrder(ctx _context.Context, query OrderCloseRequest) (OrderChannelResponse, *_nethttp.Response, error) {
	model := Struct2MapName(query)
	request := ApiRequest{ctx: ctx, ApiService: (*Service)(p), Path: "api/pay/close", PayModel: model}
	return postExecute[OrderChannelResponse](request)
}
