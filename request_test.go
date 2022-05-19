package jeepay_go_sdk

import (
	"encoding/json"
	"testing"
)

type Base struct {
	Key string `json:"Key"`
}

type BaseIn struct {
	Base
	Key string `json:"Key" json:"value,omitempty"`
}

func TestStruct(t *testing.T) {

	in := BaseIn{
		Key: "value_01",
		Base: Base{
			Key: "key_01",
		},
	}

	marshal, err := json.Marshal(in)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(marshal))
}
