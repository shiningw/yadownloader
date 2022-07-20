package http

import (
	"io/fs"
	"net/http"
	"strings"
)

type assetController struct {
	fs fs.FS
}

func (s assetController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if strings.HasSuffix(r.URL.Path, ".js") || strings.HasSuffix(r.URL.Path, ".map") {
		fileContents, err := fs.ReadFile(s.fs, r.URL.Path+".gz")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")

		if _, err := w.Write(fileContents); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	http.FileServer(http.FS(s.fs)).ServeHTTP(w, r)
}
