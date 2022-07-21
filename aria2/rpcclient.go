package aria2

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

type RpcClient struct {
	uri   string
	c     *http.Client
	token string
	req   *Request
}

func NewRpcClient(u, token string) *RpcClient {
	urlParts, err := url.Parse(u)
	if err != nil {
		log.Panic(err)
	}
	c := &http.Client{}
	r := &RpcClient{uri: u, c: c, token: token, req: NewRequest("POST", *urlParts, nil)}
	return r
}

func (r *RpcClient) send(payload *bytes.Buffer) (*http.Response, error) {
	r.req.SetBody(payload)
	resp, err := r.c.Do(r.req.Request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RpcClient) Call(method string, reply interface{}, args ...interface{}) (err error) {
	var data []interface{}
	if r.token == "" {
		if args != nil {
			data = args
		}
	} else {
		tokenArg := "token:" + r.token
		data = append([]interface{}{tokenArg}, args...)
	}
	payload, err := BuildPayload(method, data)
	//log.Println(payload.String())

	if err != nil {
		return err
	}
	resp, err := r.send(payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = GetAria2Response(resp, &reply)
	if err != nil {
		return err
	}
	return nil
}
