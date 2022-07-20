package config

import (
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

type Config struct {
	Cert    string `json:"cert"`
	Keyfile string `json:"keyfile"`
	Logfile string `json:"logfile"`
	Aria2   aria2  `json:"aria2"`
	Debug   bool   `json:"debug"`
	HTTPS   bool   `json:"https"`
	Dsn     string `json:"dsn"`
	Server  server `json:"server"`
}

type server struct {
	RootDir string `json:"rootdir"`
}

type aria2 struct {
	Token       string `json:"token"`
	Port        string `json:"port"`
	DownloadDir string `json:"download_dir"`
	TorrentDir  string `json:"torrent_dir"`
	SessionFile string `json:"session_file"`
	InputFile   string `json:"input_file"`
	LogFile     string `json:"log_file"`
}

var AppConfig Config
var once sync.Once

var defaults = Config{
	Dsn: "/etc/yadownloader/downloader.db",
	Aria2: aria2{
		Token:       "yadownloader123",
		Port:        "6800",
		DownloadDir: "/etc/yadownloader/dl/downloads",
		TorrentDir:  "/etc/yadownloader/dl/torrents",
		SessionFile: "/etc/yadownloader/aria2/aria2.session",
		InputFile:   "/etc/yadownloader/aria2/aria2.session",
		LogFile:     "/etc/yadownloader/aria2/aria2.log",
	},
	Debug: false,
	HTTPS: false,
}

func GetConfig() Config {
	once.Do(func() {
		initConfig()
	})
	return AppConfig
}

func Aria2Config() aria2 {
	return GetConfig().Aria2
}

func initConfig() {
	var err error
	AppConfig, err = parseConfig("/etc/yadownloader/config.json")
	if err != nil {
		log.Printf("%s\nUse default config!", err)
		AppConfig = defaults
	}
	dir := path.Dir(AppConfig.Aria2.SessionFile)
	if !FileExists(dir) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	if !FileExists(AppConfig.Aria2.SessionFile) {
		log.Println("aria2 session file not exists, create it")
		if _, err = os.Create(AppConfig.Aria2.SessionFile); err != nil {
			log.Fatal(err)
		}
	}
	if !FileExists(AppConfig.Aria2.DownloadDir) {
		log.Println(AppConfig.Aria2.DownloadDir, " not exists, create it")
		if err = os.MkdirAll(AppConfig.Aria2.DownloadDir, 0755); err != nil {
			log.Fatal(err)
		}
	}
	if !FileExists(AppConfig.Aria2.TorrentDir) {
		log.Println(AppConfig.Aria2.TorrentDir, " not exists, create it")
		if err = os.MkdirAll(AppConfig.Aria2.TorrentDir, 0755); err != nil {
			log.Fatal(err)
		}
	}
}

func parseConfig(s string) (Config, error) {
	var config Config
	if !FileExists(s) {
		return config, errors.New(s + " config file not exists")
	}
	data, err := os.ReadFile(s)
	if err != nil {
		return config, err
	}
	if _, err := CheckJson(string(data)); err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, err
}

func CheckJson(s string) (bool, error) {
	content := strings.NewReader(s)
	dec := json.NewDecoder(content)
	for {
		_, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, fs.ErrNotExist)
}
