package http

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/shiningw/yadownloader/aria2"
	"github.com/shiningw/yadownloader/helper"
)

type Actions struct {
	Label string `json:"name,omitempty"`
	Path  string `json:"path,omitempty"`
}
type TableResp struct {
	Results  []*aria2.StatusInfo `json:"-"`
	RespType string              `json:"-"`
	Rows     []*Tablerow         `json:"row"`
	Title    Tableheading        `json:"title"`
	Error    error               `json:"error"`
	Status   bool                `json:"status"`
}

type Tablerow struct {
	Filename string    `json:"filename"`
	Speed    string    `json:"speed,omitempty"`
	Progress string    `json:"progress,omitempty"`
	Status   string    `json:"status,omitempty"`
	Gid      string    `json:"data_gid,omitempty"`
	Actions  []Actions `json:"actions"`
}

type Tableheading []string

var tableHeading = Tableheading{
	"Filename",
	"Status",
	"Actions",
}

func NewTableResp(respType string, results []*aria2.StatusInfo) *TableResp {
	table := &TableResp{}
	if results == nil {
		table.Error = fmt.Errorf("no results")
		table.Status = false
		return table
	}
	table.RespType = respType
	table.Results = results
	table.Status = true
	table.Transform()
	return table
}

func (t *TableResp) Transform() {
	for _, info := range t.Results {
		row := &Tablerow{}
		//fmt.Printf("%+v\n", info)
		totalLength, _ := strconv.ParseFloat(info.TotalLength, 64)
		completedLength, _ := strconv.ParseFloat(info.CompletedLength, 64)
		//row.Gid = info.Following
		if row.Gid == "" {
			row.Gid = info.Gid
		}
		row.Filename = fmt.Sprintf("%s||%s", path.Base(info.Files[0].Path), helper.FormatBytes(totalLength))
		if strings.ToLower(info.Status) == "error" {
			row.Status = info.ErrorMessage
		} else {
			row.Status = info.Status
		}

		switch t.RespType {
		case "active":
			var percentage float64
			if totalLength > 0 {
				percentage = completedLength / totalLength
			}
			row.Progress = fmt.Sprintf("%s(%.2f%%)", helper.FormatBytes(completedLength), percentage*100)
			speed, _ := strconv.ParseFloat(info.DownloadSpeed, 64)
			row.Speed = helper.FormatBytes(speed)
			row.Actions = []Actions{{Label: "pause", Path: "/aria2/control/pause/" + row.Gid}}
			row.Status = ""
		case "waiting":
			row.Actions = []Actions{{Label: "unpause", Path: "/aria2/control/unpause/" + row.Gid}}
			row.Actions = append(row.Actions, Actions{Label: "delete", Path: "/aria2/control/remove/" + row.Gid})
		case "failed", "stopped":
			row.Actions = []Actions{{Label: "purge", Path: "/aria2/control/remove/" + row.Gid}}
		case "complete":
			row.Actions = []Actions{{Label: "delete", Path: "/aria2/control/purge/" + row.Gid}}
		}
		t.Rows = append(t.Rows, row)
	}
	if t.RespType == "active" {
		t.Title = Tableheading{"filename", "speed", "progress", "actions"}
	} else {
		t.Title = tableHeading
	}
}

type apiResponse struct {
	Status bool        `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
	*TableResp
}

func getApiResponse(err error, data interface{}) apiResponse {
	if err != nil {
		return apiResponse{Status: false, Error: err.Error(), Data: data}
	}
	return apiResponse{Status: true, Error: "", Data: data}
}
