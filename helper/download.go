package helper

type Downloads struct {
	Active   int `json:"active"`
	Waiting  int `json:"waiting"`
	Failed   int `json:"failed"`
	Complete int `json:"complete"`
	Total    int `json:"total,omitempty"`
	YTD      int `json:"ytd,omitempty"`
}
type DownloadStatus int

const (
	Active DownloadStatus = iota + 1
	Waiting
	Complete
	Failed
)

type DownloadType int64

const (
	Aria2 DownloadType = iota + 1
	YTD
)
