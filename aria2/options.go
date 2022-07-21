package aria2

import (
	"sort"
	"strconv"
	"strings"
)

type RunOptions struct {
	Options []string `json:"options"`
}

func NewRunOptions(token, rpcPort string) RunOptions {
	o := RunOptions{Options: Defaults}
	o.SetSecret(token)
	o.SetRPCPort(rpcPort)
	return o
}
func (o *RunOptions) AddOption(option string) {
	if option[:2] != "--" {
		option = "--" + option
	}
	//replace it if it already exists
	if i := o.hasOption(option); i != -1 {
		o.Options[i] = option
	} else {
		o.Options = append(o.Options, option)
	}
}
func (o *RunOptions) SetSecret(token string) {
	o.AddOption("--rpc-secret=" + token)
}
func (o *RunOptions) SetRPCPort(port string) {
	o.AddOption("--rpc-listen-port=" + port)
}
func (o *RunOptions) SetSessionFile(filename string) {
	o.AddOption("--save-session=" + filename)
}
func (o *RunOptions) SetLogFile(filename string) {
	o.AddOption("--log=" + filename)
}
func (o *RunOptions) SetLogLevel(level string) {
	o.AddOption("--log-level=" + level)
}
func (a Aria2Cmd) GetCmd() string {
	return strings.Join(a.Options, " ")
}
func (o *RunOptions) SetInputFile(filename string) {
	o.AddOption("--input-file=" + filename)
}
func (o *RunOptions) SetUploadSpeedLimit(limit uint) {
	o.AddOption("--max-overall-upload-limit=" + strconv.Itoa(int(limit)))
}
func (o *RunOptions) SetDownloadSpeedLimit(limit uint) {
	o.AddOption("--max-overall-download-limit=" + strconv.Itoa(int(limit)))
}
func (o *RunOptions) hasOption(option string) int {
	var index int
	if index = strings.Index(option, "="); index == -1 {
		return -1
	}
	optName := option[:index]
	for _, opt := range o.Options {
		if strings.HasPrefix(opt, optName) {
			return o.findOption(opt)
		}
	}
	return -1

}
func (o *RunOptions) findOption(option string) int {
	sort.Slice(o.Options, func(i, j int) bool { return o.Options[i] < o.Options[j] })
	i := sort.SearchStrings(o.Options, option)
	if i < len(o.Options) && o.Options[i] == option {
		return i
	}
	return -1
}

var Defaults = []string{

	"--log-level=debug",
	"--continue",
	"--daemon=true",
	"--enable-rpc=true",
	"--listen-port=51413",
	"--follow-torrent=true",
	"--enable-dht=true",
	"--enable-peer-exchange=true",
	"--peer-id-prefix=-TR2770-",
	"--user-agent=Transmission/2.77",
	"--seed-ratio=1.0",
	"--bt-seed-unverified=true",
	"--max-connection-per-server=4",
	"--max-concurrent-downloads=10",
	"--check-certificate=false",
}
