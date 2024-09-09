package main

import (
	"testing"
)

func TestExtractNumbers(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedLeft  int
		expectedRight int
		wantErr       bool
	}{
		{
			name:          "Valid number",
			input:         "111-11",
			expectedLeft:  111,
			expectedRight: 11,
			wantErr:       false,
		},
		{
			name:          "Empty left",
			input:         "1111-0",
			expectedLeft:  1111,
			expectedRight: 0,
			wantErr:       false,
		},
		{
			name:          "Empty right",
			input:         "0-1111",
			expectedLeft:  0,
			expectedRight: 1111,
			wantErr:       false,
		},
		{
			name:          "invalid no dash",
			input:         "1111",
			expectedLeft:  0,
			expectedRight: 0,
			wantErr:       true,
		},
		{
			name:          "Empty invalid missing number",
			input:         "-1111",
			expectedLeft:  0,
			expectedRight: 0,
			wantErr:       true,
		},
		{
			name:          "Empty number",
			input:         "",
			expectedLeft:  0,
			expectedRight: 0,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultLeft, resultRight, err := extractNumbers(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resultLeft != tt.expectedLeft {
				t.Errorf("extractNumbers() = %v, want %v", resultLeft, tt.expectedLeft)
			}
			if resultRight != tt.expectedRight {
				t.Errorf("extractNumbers() = %v, want %v", resultRight, tt.expectedRight)
			}
		})
	}
}

func TestConvertPental(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		wantErr  bool
	}{
		{
			name:     "Valid number 1",
			input:    1,
			expected: 1,
			wantErr:  false,
		},
		{
			name:     "Valid number 6",
			input:    11,
			expected: 6,
			wantErr:  false,
		},
		{
			name:     "Invalid number beyond 4",
			input:    15,
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertPental(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertPental() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("convertPental() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		wantErr  bool
	}{
		{
			name:     "Valid number",
			input:    "111",
			expected: 111,
			wantErr:  false,
		},
		{
			name:     "Empty number",
			input:    "",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "Invalid number with letters",
			input:    "123abc",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseNumber(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("parseNumber() = %v, want %v", result, tt.expected)
			}
		})
	}
}
