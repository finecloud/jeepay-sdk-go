package jeepay_go_sdk

import (
	"context"
	"testing"
)

const (
	host   = "food-menu.proce.top:9216"
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
	newConfiguration.ApiId = "627d2377f4c6c710fdb9b680"
	newConfiguration.ApiKey = "RPuqcEqtbmAnxRVKFKIFJGWWOK7W7tAzwE3WllVZ1Xgp5wCXpoXGKtlktwszbQE9trJgNncREu1MssaPY7tUwomjt1pW16BZ9OfhXc0D858Vhtw02nR9sFm1X652Otzi"
	newConfiguration.Host = host
	newConfiguration.Scheme = schema
	client := NewApiClient(newConfiguration)
	amount := 4231
	mchno := "M1652368247"
	mchorderno := "test05"
	waycode := AliWap
	currency := "cny"
	subject := "测试"
	body := "测试"

	request := PayCreateRequest{
		Amount:     &amount,
		MchNo:      &mchno,
		MchOrderNo: &mchorderno,
		WayCode:    &waycode,
		Currency:   &currency,
		Subject:    &subject,
		Body:       &body,
	}

	execute, response, err := client.PayApi.CreateOrder(context.Background(), request)

	if err != nil {
		t.Error(err)
	}
	t.Log(execute)
	t.Log(response)
}
