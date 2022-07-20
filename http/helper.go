package http

import (
	"encoding/json"
	"errors"
	"net/http"
)

func deny(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(getApiResponse(errors.New("Unauthorized"), nil))
}
