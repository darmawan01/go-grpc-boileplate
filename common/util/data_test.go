package util

import "testing"

func TestDefaultValueString(t *testing.T) {
	tests := []struct {
		name        string
		expected    string
		defaultData string
		data        string
	}{
		{"no data", "localhost", "localhost", ""},
		{"with data", "data 1", "localhost", "data 1"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := DefaultValueString(tt.defaultData, tt.data)
			if actual != tt.expected {
				t.Errorf("(%s, %s): expected %s, actual %s", tt.defaultData, tt.data, tt.expected, actual)
			}
		})
	}
}

func TestDefaultValueInt(t *testing.T) {
	tests := []struct {
		name        string
		expected    int
		defaultData int
		data        string
	}{
		{"no data", 3, 3, ""},
		{"wrong data", 3, 3, "f"},
		{"with data", 3, 15, "3"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := DefaultValueInt(tt.defaultData, tt.data)
			if actual != tt.expected {
				t.Errorf("(%d, %s): expected %d, actual %d", tt.defaultData, tt.data, tt.expected, actual)
			}
		})
	}
}

func TestDefaultValueBool(t *testing.T) {
	tests := []struct {
		name        string
		expected    bool
		defaultData bool
		data        string
	}{
		{"no data", false, false, ""},
		{"wrong data", false, false, "asdf"},
		{"with data 1", true, false, "true"},
		{"with data 2", false, true, "false"},
		{"with data 3", true, false, "t"},
		{"with data 4", false, true, "f"},
		{"with data 5", true, false, "TRUE"},
		{"with data 6", false, true, "FALSE"},
		{"with data 7", true, false, "T"},
		{"with data 8", false, true, "F"},
		{"with data 9", true, false, "True"},
		{"with data 10", false, true, "False"},
		{"with data 11", true, false, "1"},
		{"with data 12", false, true, "0"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := DefaultValueBool(tt.defaultData, tt.data)
			if actual != tt.expected {
				t.Errorf("(%v, %s): expected %v, actual %v", tt.defaultData, tt.data, tt.expected, actual)
			}
		})
	}
}
