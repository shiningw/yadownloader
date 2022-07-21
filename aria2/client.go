package aria2

import (
	"errors"
	"reflect"
)

var Filter = []string{"status", "followedBy", "totalLength", "errorMessage", "dir", "uploadLength", "completedLength", "downloadSpeed", "files", "numSeeders", "connections", "gid", "following", "bittorrent"}

type Aria2 struct {
	*RpcClient
	options map[string]string
}

func NewAria2(u, token string) *Aria2 {
	r := NewRpcClient(u, token)
	return &Aria2{r, make(map[string]string)}
}
func (a *Aria2) SetDownloadDir(dir string) *Aria2 {
	a.options["dir"] = dir
	return a
}
func (a *Aria2) SetFilename(filename string) *Aria2 {
	a.options["out"] = filename
	return a
}
func (a *Aria2) Shutdown() (err error) {
	return a.Call("aria2.shutdown", nil)
}

func (a *Aria2) GetVersion() (res *VersionInfo, err error) {
	var result VersionInfo
	e := a.Call("aria2.getVersion", &result)
	if e != nil {
		return nil, e
	}
	return &result, nil
}

func (a *Aria2) GetSessionInfo() (res *SessionInfo, err error) {
	var result SessionInfo
	e := a.Call("aria2.getSessionInfo", &result)
	if e != nil {
		return nil, e
	}
	return &result, nil
}

func (a *Aria2) TellStatus(gid string, keys []string) (res *StatusInfo, err error) {
	var result StatusInfo
	e := a.Call("aria2.tellStatus", &result, gid, keys)
	if e != nil {
		return nil, e
	}
	return &result, nil
}
func (a *Aria2) TellActive(keys []string) (res []*StatusInfo, err error) {
	var result []*StatusInfo
	e := a.Call("aria2.tellActive", &result, keys)
	if e != nil {
		return nil, e
	}
	return result, nil
}

func (a *Aria2) TellWaiting(offset, num int, keys []string) (res []*StatusInfo, err error) {
	var result []*StatusInfo
	e := a.Call("aria2.tellWaiting", &result, offset, num, keys)
	if e != nil {
		return nil, e
	}
	return result, nil
}
func (a *Aria2) TellStopped(offset, num int, keys []string) (res []*StatusInfo, err error) {
	var result []*StatusInfo
	e := a.Call("aria2.tellStopped", &result, offset, num, keys)

	if e != nil {
		return nil, e
	}
	return result, nil
}
func (a *Aria2) TellFailed(offset, num int, keys []string) (res []*StatusInfo, err error) {
	data, err := a.TellStopped(offset, num, keys)
	var results []*StatusInfo
	for _, result := range data {
		if result.Status != "complete" && result.Status != "removed" {
			results = append(results, result)
		}
	}
	return results, err
}
func (a *Aria2) TellCompleted(offset, num int, keys []string) (res []*StatusInfo, err error) {
	data, err := a.TellStopped(offset, num, keys)
	result := FilterResult(data)
	return result, err
}

func (a *Aria2) GetDownloadCount(queue string) int {
	var data []*StatusInfo
	switch queue {
	case "active":
		data, _ = a.TellActive(Filter)
	case "waiting":
		data, _ = a.TellWaiting(0, 100, Filter)
	case "failed":
		data, _ = a.TellFailed(0, 100, Filter)
	case "completed":
		data, _ = a.TellCompleted(0, 100, Filter)
	}
	return len(data)
}

func (a *Aria2) Pause(gid string) (err error) {
	var id string
	return a.Call("aria2.pause", &id, gid)
}

func (a *Aria2) Unpause(gid string) (err error) {
	var id string
	return a.Call("aria2.unpause", &id, gid)
}

func (a *Aria2) Remove(gid string) (err error) {
	var id string
	return a.Call("aria2.remove", &id, gid)
}

func (a *Aria2) Purge(gid string) (err error) {
	var id string
	return a.Call("aria2.removeDownloadResult", &id, gid)
}
func (a *Aria2) IsRunning() error {
	var session SessionInfo
	a.Call("aria2.getSessionInfo", &session)
	if session.Id != "" {
		return nil
	}
	return errors.New("aria2 is not running")
}
func (a *Aria2) AddUri(uris []string) (gid string, err error) {
	var result string
	e := a.Call("aria2.addUri", &result, uris, a.options)
	if e != nil {
		return "", e
	}
	return result, nil
}

func (a *Aria2) Download(url any) (gid string, err error) {
	u := reflect.TypeOf(url)
	var uris []string
	if u.Kind() == reflect.String {
		uris = []string{url.(string)}
	}
	if u == reflect.TypeOf([]string{}) {
		uris = url.([]string)
	}
	return a.AddUri(uris)
}
func (a *Aria2) AddTorrent(torrent string, uris []string) (gid string, err error) {
	var result string
	e := a.Call("aria2.addTorrent", &result, torrent, uris, a.options)
	if e != nil {
		return "", e
	}
	return result, nil
}
func (a *Aria2) AddMetalink(metalink string, options map[string]string) (gid string, err error) {
	var result string
	e := a.Call("aria2.addMetalink", &result, metalink, options)
	if e != nil {
		return "", e
	}
	return result, nil
}
func FilterResult(info []*StatusInfo) []*StatusInfo {
	var results []*StatusInfo
	for _, result := range info {
		if result.Status != "error" {
			results = append(results, result)
		}
	}
	return results
}
