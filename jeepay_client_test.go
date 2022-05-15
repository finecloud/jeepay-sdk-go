package jeepay_go_sdk

import (
	"testing"
)

// TestCreateClient 测试创建客户端
func TestCreateClient(t *testing.T) {
	configuration := Configuration{Host: ""}
	client := NewApiClient(&configuration)
	_ = client.PayApi.client
}
