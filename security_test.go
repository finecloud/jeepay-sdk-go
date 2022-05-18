package jeepay_go_sdk

import "testing"

func TestSecurity(t *testing.T) {
	//amount=10000&clientIp=192.168.0.111&mchOrderNo=P0123456789101&notifyUrl=https://www.baidu.com&platId=1000&reqTime=20190723141000&returnUrl=https://www.baidu.com&version=1.0&key=EWEFD123RGSRETYDFNGFGFGSHDFGH

	s, err := encrypt(map[string]interface{}{
		"platId":     "1000",
		"mchOrderNo": "P0123456789101",
		"amount":     "10000",
		"clientIp":   "192.168.0.111",
		"returnUrl":  "https://www.baidu.com",
		"notifyUrl":  "https://www.baidu.com",
		"reqTime":    "20190723141000",
		"version":    "1.0",
	}, "EWEFD123RGSRETYDFNGFGFGSHDFGH", "MD5")

	if err != nil {
		t.Error(err)
	}

	t.Log(s)
}

func TestMd5(t *testing.T) {
	s := md5encrypt([]byte("amount=10000&clientIp=192.168.0.111&mchOrderNo=P0123456789101&notifyUrl=https://www.baidu.com&platId=1000&reqTime=20190723141000&returnUrl=https://www.baidu.com&version=1.0&key=EWEFD123RGSRETYDFNGFGFGSHDFGH"))
	t.Log(s)
}
