package cast

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// Uint parses a string into an uint value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Uint(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

// Uint8 parses a string into a uint8 value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Uint8(s string) (uint8, error) {
	v, err := strconv.ParseUint(s, 0, 8)
	if err != nil {
		return 0, err
	}
	return uint8(v), nil
}

// Uint16 parses a string into a uint16 value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Uint16(s string) (uint16, error) {
	v, err := strconv.ParseUint(s, 0, 16)
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

// Uint32 parses a string into a uint32 value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Uint32(s string) (uint32, error) {
	v, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

// Uint64 parses a string into a uint64 value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Uint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 0, 64)
}

// Int parses a string into an int value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
// Returns a custom error message on parse failure.
func Int(s string) (int, error) {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return 0, fmt.Errorf("cast: cannot cast `%v` to type `%v`", s, "int")
	}
	return int(v), nil
}

// Int8 parses a string into an int8 value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Int8(s string) (int8, error) {
	v, err := strconv.ParseInt(s, 0, 8)
	if err != nil {
		return 0, err
	}
	return int8(v), nil
}

// Int16 parses a string into an int16 value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Int16(s string) (int16, error) {
	v, err := strconv.ParseInt(s, 0, 16)
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

// Int32 parses a string into an int32 value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Int32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

// Int64 parses a string into an int64 value.
// Supports decimal, hexadecimal (0x prefix), octal (0 prefix), and binary (0b prefix) formats.
func Int64(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 64)
}

// Float32 parses a string into a float32 value.
func Float32(s string) (float32, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

// Float64 parses a string into a float64 value.
func Float64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Decimal parses a string into a decimal.Decimal value for high-precision arithmetic.
// Uses the github.com/shopspring/decimal package.
func Decimal(s string) (decimal.Decimal, error) {
	return decimal.NewFromString(s)
}
