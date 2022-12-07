package github

import (
	"testing"
)

func TestStr(t *testing.T) {
	ss := []string{
		"1234567890",
		"12345678901",
		"1234567890123",
		"1234567",
		"1",
	}
	for _, s := range ss {
		id := ToConvertAreaRegionID(s)
		if id == "" {
			t.Error("Expected id", s)
		}
		t.Log(id)
	}
}
