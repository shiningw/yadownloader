package http

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
	"github.com/shiningw/yadownloader/aria2"
	"github.com/shiningw/yadownloader/config"
	"github.com/shiningw/yadownloader/db"
)

type aria2ActionControllerler struct {
	aria2 *aria2.Aria2
}

func (a aria2ActionControllerler) download(url string) (string, error) {
	a.aria2.SetDownloadDir(config.Aria2Config().DownloadDir)
	gid, err := a.aria2.Download(url)
	if err != nil {
		return "", err
	}
	return gid, nil
}

func (a aria2ActionControllerler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var p interface{}
	var resp apiResponse
	log.Println(config.Aria2Config())
	switch action {
	case "start":
		opts := aria2.NewRunOptions(config.Aria2Config().Token, config.Aria2Config().Port)
		opts.SetLogFile(config.Aria2Config().LogFile)
		opts.AddOption("--log-level=info")
		opts.SetInputFile(config.Aria2Config().InputFile)
		opts.SetSessionFile(config.Aria2Config().SessionFile)
		p = aria2.StartAria2(opts)
	case "stop":
		p = aria2.StopAria2(a.aria2)
	case "isrunning":
		p = aria2.Aria2IsRunning(a.aria2)
	case "version":
		p, _ = a.aria2.GetVersion()
	case "download":
		var data map[string]string
		body := requestBody{r.Body}
		var err error
		err = body.getData(&data)
		if err != nil {
			resp = apiResponse{Error: err.Error(), Status: false}

		} else {
			//a.aria2.SetDownloadDir(serverSettings.Aria2.DownloadDir)
			filename := path.Base(data["text-input-value"])
			//a.aria2.SetFilename(filename)
			gid, err := a.download(data["text-input-value"])
			if err != nil {
				resp = getApiResponse(err, nil)
			}
			if gid != "" {
				row := db.DownloadsRow{
					Filename:  filename,
					Gid:       gid,
					Uid:       "admin",
					Url:       data["text-input-value"],
					Type:      int64(db.Aria2),
					Timestamp: time.Now().Unix(),
					Status:    int64(db.Active),
				}
				db.Save(row, "aria2")
				resp = getApiResponse(nil, gid)
			}
		}
	}
	if action != "download" {
		if p == nil {
			resp = getApiResponse(nil, nil)
		} else {
			e := p.(error)
			resp = getApiResponse(e, nil)
		}
	}
	json.NewEncoder(w).Encode(resp)
}

type aria2DataController struct {
	aria2 *aria2.Aria2
}

func (a aria2DataController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	typ := vars["type"]
	var resp *TableResp
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var p interface{}
	switch typ {
	case "active":
		p, _ = a.aria2.TellActive(aria2.Filter)
	case "waiting":
		p, _ = a.aria2.TellWaiting(0, 999, aria2.Filter)
	case "failed", "stopped":
		p, _ = a.aria2.TellFailed(0, 999, aria2.Filter)
	case "complete":
		p, _ = a.aria2.TellCompleted(0, 999, aria2.Filter)
	default:
		p, _ = a.aria2.TellActive(aria2.Filter)
	}
	if p == nil {
		resp = &TableResp{Status: true}
	} else {
		resp = NewTableResp(typ, p.([]*aria2.StatusInfo))
	}
	json.NewEncoder(w).Encode(resp)
}

type aria2ActionController struct {
	*aria2.Aria2
}

// http handler for aria2 control
func (a aria2ActionController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]
	gid := vars["gid"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var err error
	var resp apiResponse
	switch action {
	case "pause":
		err = a.Pause(gid)
	case "unpause":
		err = a.Unpause(gid)
	case "remove":
		err = a.Remove(gid)
	case "purge":
		err = a.Purge(gid)
	}
	if err != nil {
		resp = getApiResponse(err, nil)
	} else {
		resp = getApiResponse(err, gid)
	}
	json.NewEncoder(w).Encode(resp)
}
