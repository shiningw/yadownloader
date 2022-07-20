package ytd

import "regexp"

var (
	downloadInfo = regexp.MustCompile(`\[(?P<module>(download|ExtractAudio|VideoConvertor|Merger|ffmpeg))\]((\s+|\s+Converting.*;\s+)Destination:\s+|\s+Merging formats into\s+\")(?P<filename>.*\.(?P<ext>(mp4|mp3|aac|webm|m4a|ogg|3gp|mkv|wav|flv)))`)

	siteInfo = regexp.MustCompile(`\[(?P<site>.+)]\s(?P<id>.+):\sDownloading\s.*`)
	progress = regexp.MustCompile(
		`\[download\]\s+` +
			`(?P<percentage>\d+(?:\.\d+)?%)` + //progress
			`\s+of\s+[~]?` +
			`(?P<filesize>\d+(?:\.\d+)?(?:K|M|G)iB)` + //file size
			`(?:\s+at\s+` +
			`(?P<speed>(\d+(?:\.\d+)?(?:K|M|G)iB/s)|Unknown speed))` + //speed
			`(?:\s+ETA\s+(?P<eta>([\d:]{2,8}|Unknown ETA)))?` + //estimated download time
			`(\s+in\s+(?P<totalTime>[\d:]{2,8}))?`)
)

type YtdHelper struct {
	FileInfo *regexp.Regexp
	SiteInfo *regexp.Regexp
	Progress *regexp.Regexp
}

func NewYtdHelper() *YtdHelper {
	return &YtdHelper{
		FileInfo: downloadInfo,
		SiteInfo: siteInfo,
		Progress: progress,
	}
}

func (y *YtdHelper) GetFileInfo(str string) map[string]string {
	return NamedMatches(y.FileInfo, str)
}

func (y *YtdHelper) GetSiteInfo(str string) map[string]string {
	return NamedMatches(y.SiteInfo, str)
}

func (y *YtdHelper) GetProgress(str string) map[string]string {
	return NamedMatches(y.Progress, str)
}

func NamedMatches(regex *regexp.Regexp, str string) map[string]string {
	match := regex.FindStringSubmatch(str)
	results := map[string]string{}
	for i, name := range match {
		results[regex.SubexpNames()[i]] = name
	}
	return results
}
