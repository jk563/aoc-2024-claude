package main

import (
	"path/filepath"
	"testing"
)

func TestValidateDay(t *testing.T) {
	tests := []struct {
		name    string
		day     int
		wantErr bool
	}{
		{"Valid day 1", 1, false},
		{"Valid day 25", 25, false},
		{"Valid day 15", 15, false},
		{"Invalid day 0", 0, true},
		{"Invalid day -1", -1, true},
		{"Invalid day 26", 26, true},
		{"Invalid day 100", 100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDay(tt.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDay() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetInputFilePath(t *testing.T) {
	tests := []struct {
		name string
		day  int
		want string
	}{
		{"Day 1", 1, filepath.Join("internal", "day01", "puzzle-input.txt")},
		{"Day 5", 5, filepath.Join("internal", "day05", "puzzle-input.txt")},
		{"Day 10", 10, filepath.Join("internal", "day10", "puzzle-input.txt")},
		{"Day 25", 25, filepath.Join("internal", "day25", "puzzle-input.txt")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getInputFilePath(tt.day)
			if got != tt.want {
				t.Errorf("getInputFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateArgs(t *testing.T) {
	tests := []struct {
		name    string
		day     int
		part    int
		wantErr bool
		errMsg  string
	}{
		{"Valid: all days", 0, 0, false, ""},
		{"Valid: specific day", 5, 0, false, ""},
		{"Valid: specific day and part", 5, 1, false, ""},
		{"Valid: day 1 part 2", 1, 2, false, ""},
		{"Valid: day 25 part 1", 25, 1, false, ""},
		{"Invalid: part without day", 0, 1, true, "cannot specify part without day"},
		{"Invalid: part 2 without day", 0, 2, true, "cannot specify part without day"},
		{"Invalid: day too low", -1, 0, true, "day must be between 1 and 25"},
		{"Invalid: day too high", 26, 0, true, "day must be between 1 and 25"},
		{"Invalid: part too low", 5, -1, true, "part must be 1 or 2"},
		{"Invalid: part too high", 5, 3, true, "part must be 1 or 2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateArgs(tt.day, tt.part)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil {
				if err.Error() != tt.errMsg {
					t.Errorf("validateArgs() error message = %v, want %v", err.Error(), tt.errMsg)
				}
			}
		})
	}
}
