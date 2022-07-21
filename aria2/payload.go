package aria2

import (
	"bytes"
	"encoding/json"
)

type payload struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      string      `json:"id"`
}

func NewPayload(method string, args []interface{}) *payload {
	return &payload{
		Version: "2.0",
		Method:  method,
		Params:  args,
		Id:      "aria2go",
	}
}

func (p *payload) Encode() (*bytes.Buffer, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(p); err != nil {
		return nil, err
	}
	return &buf, nil
}

func BuildPayload(method string, args []interface{}) (*bytes.Buffer, error) {
	return NewPayload(method, args).Encode()
}
