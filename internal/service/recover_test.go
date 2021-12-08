package service

import "testing"

func TestCollectMetadataFromLog(t *testing.T) {
	data, err := CollectMetadataFromLog("../../tests/log/data.log")
	if err != nil {
		t.Fatal(err)
	}
	for d := range data {
		if d.Volume != 0 {
			t.Logf("%s\r\n", d.String())
		}
	}
}
