package processor

import (
	"regexp"
	"strings"
)

func Quote(s string) string {
	s = strings.TrimSpace(s)
	pattern := regexp.MustCompile(`\s*?([,.:;!]+)`)
	s = pattern.ReplaceAllString(s, "$1")
	pattern1 := regexp.MustCompile(`([,.:;!]+)([A-Za-z])`)
	s = pattern1.ReplaceAllString(s, "$1 $2")
	pattern2 := regexp.MustCompile(`'\s*(.*?)\s*'`)
	s = pattern2.ReplaceAllString(s, "'$1'")
	pattern3 := regexp.MustCompile(`"\s*(.*?)\s*"`)
	s = pattern3.ReplaceAllString(s, `"$1"`)
	return s
}
