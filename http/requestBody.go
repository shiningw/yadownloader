package http

import (
	"encoding/json"
	"io"
)

type requestBody struct {
	body io.ReadCloser
}

func (r *requestBody) getData(data *map[string]string) error {
	decoder := json.NewDecoder(r.body)
	err := decoder.Decode(&data)
	return err
}
