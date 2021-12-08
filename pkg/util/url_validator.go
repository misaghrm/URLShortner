package util

import (
	"regexp"
	"strings"
)

func IsURLValid(url string) bool {
	reg := regexp.MustCompile("((http|https|ftp)://)(www.)?[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)")
	url = strings.TrimSpace(url)
	return reg.MatchString(url)
}
