package cast

import (
	"errors"
	"strconv"
	"time"
)

// timeFormats defines the supported time formats for parsing.
// The formats are tried in order until a successful parse is found.
// Includes RFC standards, common date/time formats, and timestamp formats.
var timeFormats = [...]string{
	time.RFC3339,                    // 2006-01-02T15:04:05Z07:00
	"2006-01-02 15:04:05",           // YYYY-MM-DD HH:MM:SS
	"2006-01-02T15:04:05",           // YYYY-MM-DDTHH:MM:SS
	"2006-01-02",                    // YYYY-MM-DD
	"2006/01/02 15:04:05",           // YYYY/MM/DD HH:MM:SS
	"2006/01/02",                    // YYYY/MM/DD
	"01/02/2006 15:04:05",           // MM/DD/YYYY HH:MM:SS
	"01/02/2006",                    // MM/DD/YYYY
	"Jan 2, 2006 15:04:05",          // Month D, YYYY HH:MM:SS
	"Jan 2, 2006",                   // Month D, YYYY
	"2006-01-02 15:04:05.999999999", // YYYY-MM-DD HH:MM:SS with nanoseconds
	"2006-01-02T15:04:05.999999999", // YYYY-MM-DDTHH:MM:SS with nanoseconds
	time.RFC3339Nano,                // RFC3339 with nanoseconds
	time.RFC1123,                    // Mon, 02 Jan 2006 15:04:05 MST
	time.RFC1123Z,                   // Mon, 02 Jan 2006 15:04:05 -0700
	time.RFC822,                     // 02 Jan 06 15:04 MST
	time.RFC822Z,                    // 02 Jan 06 15:04 -0700
	time.Stamp,                      // Jan _2 15:04:05
	time.StampMilli,                 // Jan _2 15:04:05.000
	time.StampNano,                  // Jan _2 15:04:05.000000000
	time.Layout,                     // 01/02 03:04:05PM '06 -0700
}

// Time parses a string into a time.Time value.
// It tries multiple common time formats including RFC3339, various date/time formats,
// and Unix timestamps. The function returns an error if none of the formats match.
//
// Supported formats:
//   - RFC standards (RFC3339, RFC1123, RFC822, etc.)
//   - Common date formats (YYYY-MM-DD, MM/DD/YYYY, etc.)
//   - Date-time combinations with various separators
//   - Unix timestamps (seconds since epoch)
//   - Nanosecond precision support
//
// Example:
//
//	t, err := Time("2023-12-25T10:30:45Z")
//	t, err := Time("2023-12-25")
//	t, err := Time("1703505045") // Unix timestamp
func Time(s string) (time.Time, error) {
	// Try each predefined format
	for _, format := range timeFormats {
		if t, err := time.Parse(format, s); err == nil {
			return t, nil
		}
	}

	// Try parsing as Unix timestamp (seconds since epoch)
	if unix, err := strconv.ParseInt(s, 10, 64); err == nil {
		return time.Unix(unix, 0), nil
	}

	return time.Time{}, errors.New("cast: unable to parse time string: " + s)
}

// Duration parses a string into a time.Duration value.
// It supports multiple input formats:
//   - Standard Go duration format (e.g., "1h30m45s", "500ms", "2h")
//   - Integer values are treated as nanoseconds
//   - Floating-point values are treated as seconds
//
// The function tries to parse in the following order:
//  1. Integer (interpreted as nanoseconds)
//  2. Float (interpreted as seconds)
//  3. Standard Go duration format
//
// Example:
//
//	d, err := Duration("1h30m45s")  // 1 hour, 30 minutes, 45 seconds
//	d, err := Duration("1.5")       // 1.5 seconds
//	d, err := Duration("1000000000") // 1 second (1 billion nanoseconds)
func Duration(s string) (time.Duration, error) {
	// Try parsing as integer (treat as nanoseconds)
	if n, err := strconv.Atoi(s); err == nil {
		return time.Duration(n), nil
	}

	// Try parsing as float (treat as seconds)
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return time.Duration(f * float64(time.Second)), nil
	}

	// Try standard Go duration format (e.g., "1h30m", "500ms")
	if d, err := time.ParseDuration(s); err == nil {
		return d, nil
	}

	return 0, errors.New("cast: unable to parse duration string: " + s)
}
