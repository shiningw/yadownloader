package ytd

import (
	"sort"
	"strings"
)

type YtdOptions struct {
	Options []string `json:"options"`
}

func NewYtdOptions(opts []string) YtdOptions {
	var options []string
	if opts == nil {
		options = Defaults
	} else {
		options = append(Defaults, opts...)
	}
	o := YtdOptions{Options: options}
	return o
}
func (o *YtdOptions) AddOption(option string, prepend bool) {
	if option[:2] != "--" && !strings.HasPrefix(option, "http") {
		option = "--" + option
	}
	if i := o.hasOption(option); i != -1 {
		o.Options[i] = option
	} else {
		if prepend {
			o.Options = append([]string{option}, o.Options...)
		} else {
			o.Options = append(o.Options, option)
		}
	}
}

func (o *YtdOptions) SetUrl(url string) {
	o.AddOption(url, true)
}

func (o *YtdOptions) SetOutput(dir string) {
	o.Options = append(o.Options, "--output")
	o.Options = append(o.Options, dir)
}

func (o *YtdOptions) SetFormt(format string) {
	o.Options = append(o.Options, "--format")
	o.Options = append(o.Options, format)
}

func (o *YtdOptions) GetUrl() string {
	url := o.Options[0]
	if strings.HasPrefix(url, "http") {
		return url
	}
	return ""
}

func (o *YtdOptions) hasOption(option string) int {
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
func (o *YtdOptions) findOption(option string) int {
	sort.Slice(o.Options, func(i, j int) bool { return o.Options[i] < o.Options[j] })
	i := sort.SearchStrings(o.Options, option)
	if i < len(o.Options) && o.Options[i] == option {
		return i
	}
	return -1
}

var Defaults = []string{
	"--no-mtime",
	"--ignore-errors",
}
