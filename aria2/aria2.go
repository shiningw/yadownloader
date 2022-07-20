package aria2

import (
	"fmt"
	"path"

	"github.com/shiningw/aria2go/client"

	"github.com/shiningw/aria2go/aria2"
	"github.com/shiningw/yadownloader/config"
)

type Aria2Settings struct {
	Version     *client.VersionInfo `json:"version"`
	Status      bool                `json:"status"`
	DownloadDir string              `json:"download_dir"`
	TorrentDir  string              `json:"torrent_dir"`
	Port        string              `json:"-"`
	Token       string              `json:"-"`
}

var aria2RpcUrl = fmt.Sprintf("http://localhost:%s/jsonrpc", config.Aria2Config().Port)
var Client = aria2.NewAria2(aria2RpcUrl, config.Aria2Config().Token)

func Download(url string, dir string) (string, error) {
	Client.SetDownloadDir(dir)
	filename := path.Base(url)
	Client.SetFilename(filename)
	return Client.AddUri([]string{url})
}

func GetAria2Settings() (Aria2Settings, error) {
	var e error
	var version *client.VersionInfo
	aria2Settings := Aria2Settings{Status: false, Version: &client.VersionInfo{}}
	aria2Settings.DownloadDir = config.Aria2Config().DownloadDir
	aria2Settings.TorrentDir = config.Aria2Config().TorrentDir
	aria2Settings.Port = config.Aria2Config().Port
	aria2Settings.Token = config.Aria2Config().Token
	if e = Client.IsRunning(); e != nil {
		aria2Settings.Status = false
	} else {
		aria2Settings.Status = true
	}

	if version, e = Client.GetVersion(); e == nil {
		aria2Settings.Version = version
	}

	return aria2Settings, e
}
