package db

import "github.com/shiningw/aria2go/aria2"

type Downloads struct {
	Active   int `json:"active"`
	Waiting  int `json:"waiting"`
	Failed   int `json:"failed"`
	Complete int `json:"complete"`
	Total    int `json:"total,omitempty"`
	YTD      int `json:"ytd"`
}
type DownloadStatus int

const (
	Active DownloadStatus = iota + 1
	Waiting
	Complete
	Failed
)

type DownloadType int

const (
	Aria2 DownloadType = iota + 1
	YTD
)

func GetDownloadCount(aria2Client *aria2.Aria2) Downloads {
	downloads := Downloads{}
	downloads.Active = aria2Client.GetDownloadCount("active")
	downloads.Waiting = aria2Client.GetDownloadCount("waiting")
	downloads.Failed = aria2Client.GetDownloadCount("failed")
	downloads.Complete = aria2Client.GetDownloadCount("completed")
	downloads.Total = downloads.Active + downloads.Waiting + downloads.Failed + downloads.Complete
	downloads.YTD = GetCount(YTD)
	return downloads
}
