package gconfig

import (
	"testing"
)

func TestLoadJsonFile(t *testing.T) {
	testFile := "testdata/config.json"

	_, err := LoadJsonFile(testFile)
	if err != nil {
		t.Errorf("LoadJsonFile '%s' failed: %s", testFile, err.Error())
	}
}
