package logger

import "strings"

var (
	colorlessTag   = ""
	tagIndentWidth = 2
)

func WithTag(value string, f func() error) error {
	savedTag := colorlessTag
	colorlessTag = value
	err := f()
	colorlessTag = savedTag

	return err
}

func SetTag(value string) {
	colorlessTag = value
}

func formattedTag() string {
	if len(colorlessTag) == 0 {
		return ""
	}

	colorizedTag := colorize(colorlessTag, tagFormat...)

	return strings.Join([]string{colorizedTag, strings.Repeat(" ", tagIndentWidth)}, "")
}

func tagBlockWidth() int {
	if len(colorlessTag) == 0 {
		return 0
	} else {
		return len(formattedTag()) + tagIndentWidth
	}
}
