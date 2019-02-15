package logger

import (
	"fmt"
	"strings"
)

var (
	colorlessTag    = ""
	tagIndentWidth  = 2
	tagColorizeFunc = func(a ...interface{}) string { return fmt.Sprint(a...) }
)

func WithTag(value string, colorizeFunc func(...interface{}) string, f func() error) error {
	savedTag := colorlessTag
	savedColorizeFunc := tagColorizeFunc
	SetTag(value, colorizeFunc)
	err := f()
	SetTag(savedTag, savedColorizeFunc)
	return err
}

func SetTag(value string, colorizeFunc func(...interface{}) string) {
	colorlessTag = value
	tagColorizeFunc = colorizeFunc
}

func formattedTag() string {
	if len(colorlessTag) == 0 {
		return ""
	}

	colorizedTag := tagColorizeFunc(colorlessTag)

	return strings.Join([]string{colorizedTag, strings.Repeat(" ", tagIndentWidth)}, "")
}

func tagBlockWidth() int {
	if len(colorlessTag) == 0 {
		return 0
	} else {
		return len(colorlessTag) + tagIndentWidth
	}
}
