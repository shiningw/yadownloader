package helper

import (
	"crypto/md5"
	"fmt"
	"math"
	"os/exec"
	"time"
)

func GenGid(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func FormatBytes(bytes float64) string {
	if bytes < 1 {
		return "0B"
	}
	suffixes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	base := math.Log(bytes) / math.Log(1024)
	n := math.Pow(1024, base-math.Floor(base))
	return fmt.Sprintf("%.2f %s", n, suffixes[int(base)])
}

func LookupBinary(name string) (string, error) {
	return exec.LookPath(name)
}
