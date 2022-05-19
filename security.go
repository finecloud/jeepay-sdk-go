package jeepay_go_sdk

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func encrypt(params map[string]interface{}, key string, securityType string) (string, error) {
	// key 按照 ascii吗从小到大排序
	keys := make([]string, 0)
	for key, _ := range params {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var buffer bytes.Buffer
	for _, item := range keys {
		value := params[item]
		if value == nil {
			continue
		}
		if value == "" {
			continue
		}
		buffer.WriteString(item)
		buffer.WriteString("=")
		buffer.WriteString(Strval(value))
		buffer.WriteString("&")
	}

	buffer.WriteString("key=")
	buffer.WriteString(key)
	if securityType == "MD5" {
		return md5encrypt(buffer.Bytes()), nil
	} else if securityType == "RSA" {
		return "", nil
	} else {
		return md5encrypt(buffer.Bytes()), nil
	}
}

func md5encrypt(buffer []byte) string {
	sum := md5.Sum(buffer)
	sec := fmt.Sprintf("%x", sum)
	return strings.ToUpper(sec)
}
