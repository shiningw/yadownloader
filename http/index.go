package http

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shiningw/yadownloader/frontend"
	"github.com/shiningw/yadownloader/settings"
)

type indexController struct {
}

func (i indexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errors error
	serverSettings := settings.SettingsResp()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Security-Policy", `default-src script-src 'self'; style-src 'self' 'unsafe-inline';`)
	w.Header().Set("Cache-Control", "no-cache")
	tpl, errors := template.ParseFS(frontend.Value(), "tpl/index.tpl.html", "tpl/content.tpl.html")

	if errors != nil {
		log.Println(errors)
	}
	s := settings.Data{Data: serverSettings.JS()}
	errors = tpl.Execute(w, s)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusInternalServerError)
	}
}
