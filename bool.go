package cast

import "strconv"

// Bool returns the boolean value represented by the string.
func Bool(s string) (bool, error) {
	return strconv.ParseBool(s)
}
