package http

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/shiningw/yadownloader/aria2"
	"github.com/shiningw/yadownloader/frontend"
	"github.com/shiningw/yadownloader/ytd"
)

type FBConfig struct {
	Key     []byte `json:"key"`
	RootDir string `json:"rootDir"`
}

func RegisterRoutes(r *mux.Router, data FBConfig) {
	//log.Println(data.Server.Root, data.Settings.UserHomeBasePath)
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := authentication{data.Key, r}

			if r.URL.Path == "/downloader" {
				if !auth.auth() {
					http.Redirect(w, r, "/login/", http.StatusFound)
					return
				}
			}
			if strings.HasPrefix(r.URL.Path, "/downloader/") {
				if !auth.auth() {
					deny(r, w)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	})
	r.Handle("/downloader", indexController{})
	r.Handle("/downloader/aria2/data/{type}", aria2DataController{aria2.GetClient()})
	r.Handle("/downloader/aria2/action/{action}", aria2ActionControllerler{aria2.GetClient()})
	r.Handle("/downloader/aria2/control/{action}/{gid}", aria2ActionController{aria2.GetClient()})
	r.PathPrefix("/downloader/static").Handler(http.StripPrefix("/downloader/static/", assetController{frontend.Value()}))
	r.Handle("/downloader/counters", counters{aria2: aria2.GetClient()})
	r.Handle("/downloader/ytd/action/{action}", ytdController{ytd: ytd.NewYtdCmd(ytdOptions())})
	r.HandleFunc("/downloader/ytd/downloads", getYTDDownloads)
	r.Handle("/downloader/aria2/upload", torrentController{aria2: aria2.GetClient()})
}
func ytdOptions() ytd.YtdOptions {
	opts := ytd.NewYtdOptions(nil)
	opts.AddOption("--force-ipv4", false)
	opts.SetFormt("bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best")
	//opts.AddOption(url, true)
	//opts.AddOption("--OUTPUT", false)
	//opts.AddOption("/tmp/", false)
	//if config.FileExists(config.Aria2Config().DownloadDir) {
	//	os.Mkdir(config.Aria2Config().DownloadDir, 0755)
	//	}
	//opts.SetOutput(config.Aria2Config().DownloadDir + "/%(title)s.%(ext)s")
	return opts
}
