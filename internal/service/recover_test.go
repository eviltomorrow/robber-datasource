package service

import "testing"

func TestFetchMetadataFromLog(t *testing.T) {
	data, err := FetchMetadataFromLog("../../tests/log/data.log")
	if err != nil {
		t.Fatal(err)
	}
	for d := range data {
		if d.Volume != 0 {
			t.Logf("%s\r\n", d.String())
		}
	}
}
