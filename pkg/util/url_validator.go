package util

import (
	"net/url"
)

func IsURLValid(URL string) (*url.URL, error) {
	parse, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	return parse, err
}
