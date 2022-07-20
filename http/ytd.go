package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shiningw/yadownloader/db"
	"github.com/shiningw/yadownloader/ytd"
)

type ytdController struct {
	ytd *ytd.YtdCmd
}

func (y ytdController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action := vars["action"]
	var resp apiResponse
	var data map[string]string
	body := requestBody{r.Body}
	err := body.getData(&data)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp = getApiResponse(err, nil)
		json.NewEncoder(w).Encode(resp)
		return
	}
	switch action {
	case "download":
		resp = y.download(data["text-input-value"])
	case "delete":
		gid := data["gid"]
		resp = y.delete(gid)
	}
	json.NewEncoder(w).Encode(resp)
}
func (y ytdController) download(url string) apiResponse {
	var resp apiResponse
	e := y.ytd.Download(url)
	if e != nil {
		resp = getApiResponse(e, nil)
	} else {
		resp = getApiResponse(e, y.ytd.Filename())
	}

	return resp
}

func (y ytdController) delete(gid string) apiResponse {
	var resp apiResponse
	var err error
	_, err = db.DeleteByGid(gid)
	if err != nil {
		resp = getApiResponse(err, nil)
	} else {
		resp = getApiResponse(err, gid)
	}
	return resp
}

type tableResp struct {
	TableResp
	data []db.YtdQueue
}

func (t *tableResp) TransForm() {
	var row *Tablerow
	if len(t.data) < 1 {
		return
	}
	t.TableResp.Status = true
	heading := Tableheading{
		"Filename",
		"Speed",
		"Progress",
		"Actions",
	}
	t.Title = heading

	for _, v := range t.data {
		row = &Tablerow{
			Filename: v.Filename,
			Speed:    v.Speed,
			Progress: v.Progress,
			Gid:      v.Gid,
			Status:   "",
			Actions:  []Actions{{Label: "delete", Path: "/ytd/action/delete"}},
		}
		t.Rows = append(t.Rows, row)
	}
}

func getYTDDownloads(w http.ResponseWriter, r *http.Request) {
	resp := &tableResp{
		TableResp{},
		db.GetYTDDownloads(),
	}
	resp.TransForm()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
