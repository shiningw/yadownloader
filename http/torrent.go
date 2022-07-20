package http

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/shiningw/aria2go/aria2"
	"github.com/shiningw/yadownloader/config"
)

type torrentController struct {
	aria2 *aria2.Aria2
	r     *http.Request
	w     http.ResponseWriter
}

func (t torrentController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.r = r
	t.w = w
	switch r.Method {
	case "POST":
		t.receiveFile()
	default:
		json.NewEncoder(w).Encode(getApiResponse(errors.New("http: method not allowed"), nil))
	}
}

//receive file from client
func (t torrentController) receiveFile() {
	var resp apiResponse
	//set max upload file size to 3MB
	t.r.ParseMultipartForm(2 << 30)

	file, fileHeader, err := getUploadFile(t.r) //r.FormFile("torrentfile")
	if err != nil {
		json.NewEncoder(t.w).Encode(getApiResponse(err, nil))
		return
	}
	defer file.Close()
	//_, err = os.OpenFile(config.Aria2Config().TorrentDir+"/"+fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	buf := make([]byte, fileHeader.Size)
	n, err := file.Read(buf)
	if n != int(fileHeader.Size) {
		json.NewEncoder(t.w).Encode(getApiResponse(err, nil))
		return
	}

	content := base64.StdEncoding.EncodeToString(buf)
	saveFile(config.Aria2Config().TorrentDir+"/"+fileHeader.Filename, buf)

	gid, err := t.btDownload(content)
	if err == nil && gid != "" {
		resp = getApiResponse(nil, gid)
	} else {
		resp = getApiResponse(err, nil)
	}
	json.NewEncoder(t.w).Encode(resp)
}

func (t torrentController) btDownload(filename string) (gid string, err error) {
	var content string
	if len(filename) > 255 {
		content = filename
	} else {
		c, err := os.ReadFile(filename)
		if err != nil {
			return "", err
		}
		content = base64.StdEncoding.EncodeToString(c)
	}
	uris := []string{}
	t.aria2.SetDownloadDir(config.Aria2Config().DownloadDir)
	return t.aria2.AddTorrent(content, uris)
}

//save file to disk
func saveFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func getUploadFile(r *http.Request) (multipart.File, *multipart.FileHeader, error) {

	if r.MultipartForm != nil && r.MultipartForm.File != nil {

		for _, files := range r.MultipartForm.File {
			if len(files) < 1 {
				continue
			}
			f, err := files[0].Open()
			return f, files[0], err
		}
	}
	return nil, nil, errors.New("http: no file found")
}
