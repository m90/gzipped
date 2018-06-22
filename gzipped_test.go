package gzipped

import (
	"bytes"
	"testing"
)

func TestHumanize(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected string
	}{
		{"bytes", 88, "88B"},
		{"kilobytes", 12346, "12.1K"},
		{"megabytes", 2434622, "2.3M"},
		{"gigabytes", 54346229102, "50.6G"},
		{"terabytes", 243462291021234, "221.4T"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := humanize(test.input)
			if test.expected != result {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name             string
		input            []byte
		expectError      bool
		expectedOutBytes uint64
	}{
		{
			"default",
			[]byte("this is a string that tests whatever \n happens when we gzip its contents"),
			false,
			87,
		},
		{
			"empty",
			[]byte(""),
			true,
			0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBuffer(test.input)
			result, err := Compare(buf)
			if test.expectError != (err != nil) {
				t.Fatalf("Unexpected error %v", err)
			}
			if test.expectedOutBytes != 0 && test.expectedOutBytes != result.OutBytes {
				t.Errorf("Expected %v out bytes, got %v", test.expectedOutBytes, result.OutBytes)
			}
		})
	}
}
