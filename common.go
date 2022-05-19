package jeepay_go_sdk

import (
	"encoding/json"
	"strconv"
)

type WayCode string

type PayDataType string

const (
	// QrChshier 聚合扫码(用户扫商家)
	QrChshier WayCode = "QR_CHSHIER"
	// AutoBar 聚合条码(商家扫用户)
	AutoBar WayCode = "AUTO_BAR"
	// AliBar 支付宝条码
	AliBar WayCode = "ALI_BAR"
	// AliJsapi 支付宝生活号
	AliJsapi WayCode = "ALI_JSAPI"
	// AliApp 支付宝APP
	AliApp WayCode = "ALI_APP"
	// AliWap 支付宝WAP
	AliWap WayCode = "ALI_WAP"
	// AliPc 支付宝PC网站
	AliPc WayCode = "ALI_PC"
	// AliQr 支付宝二维码
	AliQr WayCode = "ALI_QR"
	// WxBar 微信条码
	WxBar WayCode = "WX_BAR"
	// WxJsapi 微信公众号
	WxJsapi WayCode = "WX_JSAPI"
	// WxLite 微信小程序
	WxLite WayCode = "WX_LITE"
	// WxApp 微信APP
	WxApp WayCode = "WX_APP"
	// WxH5 微信H5
	WxH5 WayCode = "WX_H5"
	// WxNative 微信扫码
	WxNative WayCode = "WX_NATIVE"
	// YsfBar 云闪付条码
	YsfBar WayCode = "YSF_BAR"
	// YsfJsapi 云闪付jsapi
	YsfJsapi WayCode = "YSF_JSAPI"

	PayDataTypeForm   = "form"
	PayDataTypePayUrl = "payUrl"
)

type Response[V interface{}] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Sign string `json:"sign"`
	Data V      `json:"data"`
}

func Struct2MapName(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(obj)
	_ = json.Unmarshal(j, &m)
	return m
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
