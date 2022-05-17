package jeepay_go_sdk

import (
	"context"
	"testing"
)

const (
	host   = "food-menu.proce.top:90"
	schema = "http"
)

// TestCreateClient 测试创建客户端
func TestCreateClient(t *testing.T) {
	newConfiguration := NewConfiguration()
	newConfiguration.ApiId = "test"
	newConfiguration.ApiKey = "test"
	newConfiguration.Host = host
	newConfiguration.Scheme = schema
	client := NewApiClient(newConfiguration)
	_ = client.PayApi.client
}

func TestPayApi(t *testing.T) {
	newConfiguration := NewConfiguration()
	newConfiguration.ApiId = "test"
	newConfiguration.ApiKey = "test"
	newConfiguration.Host = host
	newConfiguration.Scheme = schema
	client := NewApiClient(newConfiguration)
	order := client.PayApi.CreateOrder(context.Background())
	order.PayModel = &PayModel{
		Amount:    100,
		Subject:   "测试",
		Body:      "测试",
		NotifyUrl: "http://www.baidu.com",
	}

	execute, response, err := client.PayApi.PostOrderExecute(order)
	if err != nil {
		t.Error(err)
	}
	t.Log(execute)
	t.Log(response)
}
