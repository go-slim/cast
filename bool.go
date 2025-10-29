package cast

import "strconv"

// Bool parses a string into a boolean value.
// It accepts "1", "t", "T", "true", "TRUE", "True" for true,
// and "0", "f", "F", "false", "FALSE", "False" for false.
// Any other value returns an error.
func Bool(s string) (bool, error) {
	return strconv.ParseBool(s)
}
