package service

import "testing"

func TestBuildRangeCode(t *testing.T) {
	ch := BuildRangeCode()
	for code := range ch {
		t.Logf("code: %s\r\n", code)
	}
}
