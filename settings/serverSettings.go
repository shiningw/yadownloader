package settings

import (
	"encoding/json"
	"html/template"

	"github.com/shiningw/yadownloader/aria2"
	"github.com/shiningw/yadownloader/db"
)

type ServerSettings struct {
	Aria2    aria2.Aria2Settings `json:"aria2"`
	Name     string              `json:"name"`
	Counters db.Downloads        `json:"counters"`
	Error    error               `json:"error"`
}

func (s ServerSettings) String() string {
	d, _ := json.Marshal(s)
	return string(d)
}

func (s ServerSettings) JS() template.JS {
	d, _ := json.Marshal(s)
	return template.JS(d)
}

type Data struct {
	Data template.JS
}

func SettingsResp() ServerSettings {
	var serverSettings ServerSettings
	//stop if aria2 is not running or responding
	serverSettings.Aria2, serverSettings.Error = aria2.GetAria2Settings()
	if serverSettings.Error != nil {
		return serverSettings
	}
	serverSettings.Counters = db.GetDownloadCount(aria2.GetClient())
	return serverSettings
}
