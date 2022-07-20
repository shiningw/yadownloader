package ytd

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"path"

	"github.com/shiningw/yadownloader/config"
	"github.com/shiningw/yadownloader/db"
	"github.com/shiningw/yadownloader/helper"
)

var (
	TableName = "ytd"
)

type YtdCmd struct {
	YtdOptions
	filename string
}

const (
	maxReadLength = 512
)

func NewYtdCmd(options YtdOptions) *YtdCmd {
	//a := NewYtdCmd(options)
	a := &YtdCmd{options, ""}
	return a
}
func (y *YtdCmd) run() error {
	var bin string
	var err error
	if bin, err = exec.LookPath("yt-dlp"); err != nil {
		if bin, err = exec.LookPath("youtube-dl"); err != nil {
			return errors.New("yt-dlp or youtube-dl not found")
		}
	}
	cmd := exec.Command(bin, y.Options...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}
	y.read(stdout)
	b, _ := io.ReadAll(stderr)
	if len(b) > 0 {
		cmd.Process.Release()
		return errors.New(string(b))
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
func (y *YtdCmd) Download(url string) error {
	y.YtdOptions.SetUrl(url)
	y.AddOption("--force-ipv4", false)
	y.SetFormt("bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]/best")
	y.SetOutput(config.Aria2Config().DownloadDir + "/%(title)s.%(ext)s")
	return y.run()
}

func (y *YtdCmd) read(r io.Reader) {
	buf := make([]byte, maxReadLength)
	ytdHelper := NewYtdHelper()
	var row db.DownloadsRow
	for {
		n, err := r.Read(buf)
		if err != nil {
			break
		}
		line := string(buf[:n])
		info := ytdHelper.GetSiteInfo(line)
		if id, ok := info["id"]; ok {
			gid := helper.GenGid([]byte(id))
			row.Gid = gid
		}
		fileInfo := ytdHelper.GetFileInfo(line)
		if filename, ok := fileInfo["filename"]; ok {
			if filename != "" {
				y.filename = filename
			}
			row.Filename = path.Base(filename)
			row.Timestamp = helper.GetTimestamp()
			row.Type = int64(db.YTD)
			if moduleName, ok := fileInfo["module"]; ok {
				if moduleName == "download" {
					row.Url = y.YtdOptions.GetUrl()
					db.Save(row, "ytd")
				}
			}
		}
		progress := ytdHelper.GetProgress(line)
		if progress != nil {
			if progress["percentage"] != "" || progress["speed"] != "" {
				sql := fmt.Sprintf("UPDATE %s set filesize = ?,speed = ?,progress = ? WHERE gid = ?", TableName)
				db.Update(sql, []any{progress["filesize"], progress["speed"], progress["percentage"], row.Gid})
			}
		}
	}
}
func (y *YtdCmd) Filename() string {
	return y.filename
}
