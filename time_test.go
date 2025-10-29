package cast_test

import (
	"testing"
	"time"

	"go-slim.dev/cast"
)

func TestTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
		hasError bool
	}{
		// RFC3339 format
		{
			name:     "RFC3339",
			input:    "2023-12-25T10:30:45Z",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC3339 with timezone",
			input:    "2023-12-25T10:30:45+08:00",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 0, time.FixedZone("", 8*3600)),
			hasError: false,
		},
		// Custom formats
		{
			name:     "YYYY-MM-DD HH:MM:SS",
			input:    "2023-12-25 10:30:45",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "YYYY-MM-DDTHH:MM:SS",
			input:    "2023-12-25T10:30:45",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "YYYY-MM-DD",
			input:    "2023-12-25",
			expected: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "YYYY/MM/DD HH:MM:SS",
			input:    "2023/12/25 10:30:45",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "YYYY/MM/DD",
			input:    "2023/12/25",
			expected: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "MM/DD/YYYY HH:MM:SS",
			input:    "12/25/2023 10:30:45",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "MM/DD/YYYY",
			input:    "12/25/2023",
			expected: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "Month D, YYYY HH:MM:SS",
			input:    "Dec 25, 2023 10:30:45",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "Month D, YYYY",
			input:    "Dec 25, 2023",
			expected: time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			hasError: false,
		},
		// Nanosecond precision
		{
			name:     "YYYY-MM-DD HH:MM:SS with nanoseconds",
			input:    "2023-12-25 10:30:45.123456789",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 123456789, time.UTC),
			hasError: false,
		},
		{
			name:     "YYYY-MM-DDTHH:MM:SS with nanoseconds",
			input:    "2023-12-25T10:30:45.123456789",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 123456789, time.UTC),
			hasError: false,
		},
		// RFC formats
		{
			name:     "RFC1123",
			input:    "Mon, 25 Dec 2023 10:30:45 GMT",
			expected: time.Date(2023, 12, 25, 10, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "RFC822",
			input:    "25 Dec 23 10:30 GMT",
			expected: time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC),
			hasError: false,
		},
		// Unix timestamp
		{
			name:     "Unix timestamp",
			input:    "1703505045",
			expected: time.Unix(1703505045, 0),
			hasError: false,
		},
		// Stamp formats
		{
			name:     "Stamp",
			input:    "Dec 25 10:30:45",
			expected: time.Date(0, 12, 25, 10, 30, 45, 0, time.UTC),
			hasError: false,
		},
		{
			name:     "StampMilli",
			input:    "Dec 25 10:30:45.123",
			expected: time.Date(0, 12, 25, 10, 30, 45, 123000000, time.UTC),
			hasError: false,
		},
		{
			name:     "StampNano",
			input:    "Dec 25 10:30:45.123456789",
			expected: time.Date(0, 12, 25, 10, 30, 45, 123456789, time.UTC),
			hasError: false,
		},
		// Error cases
		{
			name:     "Invalid format",
			input:    "not a date",
			hasError: true,
		},
		{
			name:     "Empty string",
			input:    "",
			hasError: true,
		},
		{
			name:     "Invalid date components",
			input:    "2023-13-45",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := cast.Time(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Time(%q) expected error but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("Time(%q) unexpected error: %v", tt.input, err)
				return
			}

			if !result.Equal(tt.expected) {
				t.Errorf("Time(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Duration
		hasError bool
	}{
		// Integer input (nanoseconds)
		{
			name:     "Integer nanoseconds",
			input:    "1000000000",
			expected: time.Second,
			hasError: false,
		},
		{
			name:     "Zero nanoseconds",
			input:    "0",
			expected: 0,
			hasError: false,
		},
		// Float input (seconds)
		{
			name:     "Float seconds - integer as float",
			input:    "5.0",
			expected: 5 * time.Second,
			hasError: false,
		},
		{
			name:     "Float seconds - decimal",
			input:    "1.5",
			expected: 1500 * time.Millisecond,
			hasError: false,
		},
		{
			name:     "Float seconds - small decimal",
			input:    "0.001",
			expected: time.Millisecond,
			hasError: false,
		},
		{
			name:     "Float seconds - negative",
			input:    "-2.5",
			expected: -2500 * time.Millisecond,
			hasError: false,
		},
		// Standard Go duration format
		{
			name:     "Hours",
			input:    "2h",
			expected: 2 * time.Hour,
			hasError: false,
		},
		{
			name:     "Minutes",
			input:    "30m",
			expected: 30 * time.Minute,
			hasError: false,
		},
		{
			name:     "Seconds",
			input:    "45s",
			expected: 45 * time.Second,
			hasError: false,
		},
		{
			name:     "Milliseconds",
			input:    "500ms",
			expected: 500 * time.Millisecond,
			hasError: false,
		},
		{
			name:     "Microseconds",
			input:    "250us",
			expected: 250 * time.Microsecond,
			hasError: false,
		},
		{
			name:     "Nanoseconds",
			input:    "100ns",
			expected: 100 * time.Nanosecond,
			hasError: false,
		},
		{
			name:     "Complex duration",
			input:    "1h30m45s",
			expected: 1*time.Hour + 30*time.Minute + 45*time.Second,
			hasError: false,
		},
		{
			name:     "Complex duration with milliseconds",
			input:    "2h15m30s500ms",
			expected: 2*time.Hour + 15*time.Minute + 30*time.Second + 500*time.Millisecond,
			hasError: false,
		},
		{
			name:     "Negative duration",
			input:    "-1h30m",
			expected: -1*time.Hour - 30*time.Minute,
			hasError: false,
		},
		// Error cases
		{
			name:     "Invalid duration",
			input:    "not a duration",
			hasError: true,
		},
		{
			name:     "Empty string",
			input:    "",
			hasError: true,
		},
		{
			name:     "Invalid unit",
			input:    "5x",
			hasError: true,
		},
		{
			name:     "Invalid float",
			input:    "abc.def",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := cast.Duration(tt.input)

			if tt.hasError {
				if err == nil {
					t.Errorf("Duration(%q) expected error but got none", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("Duration(%q) unexpected error: %v", tt.input, err)
				return
			}

			if result != tt.expected {
				t.Errorf("Duration(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkTime(b *testing.B) {
	testCases := []string{
		"2023-12-25T10:30:45Z",
		"2023-12-25 10:30:45",
		"1703505045",
		"Dec 25, 2023 10:30:45",
	}

	for _, tc := range testCases {
		b.Run(tc, func(b *testing.B) {
			for b.Loop() {
				cast.Time(tc)
			}
		})
	}
}

func BenchmarkDuration(b *testing.B) {
	testCases := []string{
		"1h30m45s",
		"5000000000",
		"1.5",
		"2h",
		"-30s",
	}

	for _, tc := range testCases {
		b.Run(tc, func(b *testing.B) {
			for b.Loop() {
				cast.Duration(tc)
			}
		})
	}
}
