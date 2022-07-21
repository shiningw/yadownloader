package http

import (
	"encoding/json"
	"net/http"

	"github.com/shiningw/yadownloader/aria2"
	"github.com/shiningw/yadownloader/db"
)

type counters struct {
	aria2 *aria2.Aria2
}

func (c counters) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	downloads := db.GetDownloadCount(c.aria2)
	json.NewEncoder(w).Encode(downloads)
}
