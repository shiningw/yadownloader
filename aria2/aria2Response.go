package aria2

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Aria2Response struct {
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
	Id      string           `json:"id"`
}

func NewAria2Response(body io.Reader) (*Aria2Response, error) {
	r := Aria2Response{}
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *Aria2Response) Decode(reply interface{}) error {
	if r.Error != nil {
		jsonErr := &Error{}
		if err := json.Unmarshal(*r.Error, jsonErr); err != nil {
			return &Error{
				Code:    E_SERVER,
				Message: string(*r.Error),
			}
		}
		return jsonErr
	}

	if r.Result == nil {
		return errors.New("result is null")
	}

	return json.Unmarshal(*r.Result, reply)
}

func GetAria2Response(resp *http.Response, reply interface{}) error {
	result, err := NewAria2Response(resp.Body)
	if err != nil {
		return err
	}
	err = result.Decode(&reply)
	return err
}
