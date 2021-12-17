package util_test

import (
	"github.com/misaghrm/urlshortener/pkg/util"
	"testing"
)

func TestGetIntegerValue(t *testing.T) {
	str := map[string]int64{
		"4q9BHCRMSzE": 1471779799558000640,
		"4q9BKD177hY": 1471779875789475840,
	}
	for i, s := range str {
		t.Logf("Test %q : the %s given ", i, i)
		id := util.GetIntegerValue(i)
		if id != s {
			t.Fatalf("should be %q not %q", s, id)
		}
	}
}

func TestGetStringValue(t *testing.T) {
	str := map[int64]string{
		1471779799558000640: "4q9BHCRMSzE",
		1471779875789475840: "4q9BKD177hY",
	}
	for i, s := range str {
		t.Logf("Test %q : the %q given ", i, i)
		id := util.GetStringValue(i)
		if id != s {
			t.Fatalf("should be %q not %q", s, id)
		}
	}
}
