package util_test

import (
	"github.com/misaghrm/urlshortener/pkg/util"
	"testing"
)

func TestIsURLValid(t *testing.T) {
	urls := map[string]bool{
		"google.com":         true,
		"http://google.com":  true,
		"https://google.com": true,
		"htps://google.com":  false,
	}
	for s, b := range urls {
		t.Logf("The Given Url Is %s ", s)
		ok, _ := util.IsURLValid(s)
		if ok != b {
			t.Errorf("")
		}
	}
}
