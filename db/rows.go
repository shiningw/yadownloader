package db

import "reflect"

type IRow interface {
	FieldNames() []string
	NumField() int
}
type DownloadsRow struct {
	Filename  string `json:"filename"`
	Uid       string
	Gid       string
	Url       string
	Type      int64
	Timestamp int64
	Status    int64
}

func (r DownloadsRow) FieldNames() []string {
	return StructNames(r)
}

func (r DownloadsRow) NumField() int {
	v := reflect.ValueOf(r)
	return v.NumField()
}

type Aria2Row struct {
	Gid  string
	Data string
}

func (r Aria2Row) FieldNames() []string {
	return StructNames(r)
}
func (r Aria2Row) NumField() int {
	v := reflect.ValueOf(r)
	return v.NumField()
}

type YtdRow struct {
	Gid      string
	Speed    string
	Progress string
	Filesize string
	Data     string
}

func (r YtdRow) FieldNames() []string {
	return StructNames(r)
}
func (r YtdRow) NumField() int {
	v := reflect.ValueOf(r)
	return v.NumField()
}

type YtdQueue struct {
	Gid       string
	Url       string
	Filename  string
	Status    int64
	Timestamp int64
	Speed     string
	Progress  string
}

func (r YtdQueue) FieldNames() []string {
	return StructNames(r)
}
func (r YtdQueue) NumField() int {
	v := reflect.ValueOf(r)
	return v.NumField()
}
