package aria2

import (
	"fmt"
	"sync"

	"github.com/shiningw/yadownloader/config"
)

var (
	client *Aria2
)
var once sync.Once

func GetClient() *Aria2 {
	aria2RpcUrl := fmt.Sprintf("http://localhost:%s/jsonrpc", config.Aria2Config().Port)
	once.Do(func() {
		client = NewAria2(aria2RpcUrl, config.Aria2Config().Token)
	})
	return client
}

type Aria2Settings struct {
	Version     *VersionInfo `json:"version"`
	Status      bool         `json:"status"`
	DownloadDir string       `json:"download_dir"`
	TorrentDir  string       `json:"torrent_dir"`
	Port        string       `json:"-"`
	Token       string       `json:"-"`
	Error       string       `json:"error"`
}

func GetAria2Settings() (Aria2Settings, error) {

	var e error
	var version *VersionInfo

	aria2Settings := Aria2Settings{Status: false, Version: &VersionInfo{}}
	aria2Settings.DownloadDir = config.Aria2Config().DownloadDir
	aria2Settings.TorrentDir = config.Aria2Config().TorrentDir
	aria2Settings.Port = config.Aria2Config().Port
	aria2Settings.Token = config.Aria2Config().Token
	if e = client.IsRunning(); e != nil {
		aria2Settings.Status = false
		aria2Settings.Error = e.Error()
	} else {
		aria2Settings.Status = true
	}

	if version, e = client.GetVersion(); e == nil {
		aria2Settings.Version = version
	}

	return aria2Settings, e
}
