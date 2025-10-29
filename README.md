# cast

[English](README.md) | [简体中文](README.zh-CN.md)

A minimal, dependency-free casting utility for Go strings to native types.

Module path: `go-slim.dev/cast`

- Convert string input to Go numeric/boolean/string types.
- Support time and duration type conversions.
- Reflect-friendly helper to cast to a target `reflect.Type`.
- Simple slice support via `FromType` (comma-separated values).

## Install

```bash
go get go-slim.dev/cast
```

## Quick Start

```go
package main

import (
    "fmt"
    "reflect"
    "time"

    cast "go-slim.dev/cast"
)

func main() {
    // Basic type conversion
    v, _ := cast.FromString("42", "int")
    fmt.Printf("%T %v\n", v, v) // int 42

    v, _ = cast.FromString("true", "bool")
    fmt.Printf("%T %v\n", v, v) // bool true

    // Time conversion
    v, _ = cast.FromString("2023-12-25T10:30:45Z", "time.Time")
    fmt.Printf("%T %v\n", v, v) // time.Time 2023-12-25 10:30:45 +0000 UTC

    // Duration conversion
    v, _ = cast.FromString("1h30m", "time.Duration")
    fmt.Printf("%T %v\n", v, v) // time.Duration 1h30m0s

    // Using reflect.Type, including slices (comma-separated)
    t := reflect.TypeOf([]int(nil)) // []int
    v, _ = cast.FromType("1,2,3", t)
    fmt.Printf("%T %v\n", v, v) // []int [1 2 3]
}
```

## API

### Core Entry Points

- **`FromString(s string, targetType string) (any, error)`**
  - Supported type names:
    - Integers: `int`, `int8`, `int16`, `int32`, `int64`
    - Unsigned integers: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
    - Floats: `float32`, `float64`
    - Boolean: `bool`
    - String: `string`
    - Time: `time.Time`, `time.Duration`

- **`FromType(s string, targetType reflect.Type) (any, error)`**
  - If `targetType` is a slice (e.g., `[]int`), input is split by comma `,`, items are trimmed, and each is cast via `FromString` to the slice element type.

### Numeric Helpers (string → typed value)

- **Integer conversions**: `Int`, `Int8`, `Int16`, `Int32`, `Int64`
- **Unsigned integer conversions**: `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
- **Float conversions**: `Float32`, `Float64`
- **Decimal conversion**: `Decimal` (uses `github.com/shopspring/decimal`)

### Boolean Helper

- **`Bool(s string) (bool, error)`**
  - Supports: `true`, `false`, `1`, `0`, `t`, `f`, `TRUE`, `FALSE`, etc.

### Time Helpers

- **`Time(s string) (time.Time, error)`**
  - Supports multiple time formats:
    - RFC3339: `2023-12-25T10:30:45Z`
    - Date-time: `2023-12-25 10:30:45`, `2023/12/25 10:30:45`
    - Dates: `2023-12-25`, `12/25/2023`
    - RFC formats: RFC1123, RFC822
    - Unix timestamp: `1703505045`
    - Nanosecond precision support

- **`Duration(s string) (time.Duration, error)`**
  - Supports multiple formats:
    - Standard Go format: `1h30m45s`, `500ms`, `2h`
    - Integer (nanoseconds): `1000000000`
    - Float (seconds): `1.5`, `0.001`

### Implementation Files

- `cast/cast.go` - Dispatchers `FromString`, `FromType`
- `cast/num.go` - Numeric/boolean parsing
- `cast/time.go` - Time/duration parsing

## Error Behavior

- On unsupported type names, `FromString` returns:
  - `fmt.Errorf("cast: type %v is not supported", targetType)`
- On parse failures, helpers return parse errors from `strconv` (except `Int`, which formats a message like `cast: cannot cast \\"%v\\" to type \\"%v\\"`).
- `FromType` with slice types stops and returns the first encountered item error.

## Examples

### Basic Type Conversion

```go
// Unsigned integer
v, err := cast.FromString("1", "uint16")
// v == uint16(1), err == nil

// Parse error
v, err = cast.FromString("str", "int")
// err != nil (cannot cast)

// Slices via FromType
t := reflect.TypeOf([]uint8(nil))
v, err = cast.FromType("1, 2, 255", t)
// v == []uint8{1,2,255}, err == nil
```

### Time Conversion

```go
// RFC3339 format
t, err := cast.Time("2023-12-25T10:30:45Z")
// t == 2023-12-25 10:30:45 +0000 UTC

// Common date format
t, err = cast.Time("2023-12-25")
// t == 2023-12-25 00:00:00 +0000 UTC

// Unix timestamp
t, err = cast.Time("1703505045")
// t == 2023-12-25 10:30:45 +0000 UTC
```

### Duration Conversion

```go
// Standard format
d, err := cast.Duration("1h30m45s")
// d == 1*time.Hour + 30*time.Minute + 45*time.Second

// Float seconds
d, err = cast.Duration("1.5")
// d == 1500 * time.Millisecond

// Integer nanoseconds
d, err = cast.Duration("1000000000")
// d == time.Second
```

### Number Base Support

```go
// Hexadecimal
v, err := cast.Int("0xFF")
// v == 255

// Octal
v, err = cast.Int("0755")
// v == 493

// Binary
v, err = cast.Int("0b1010")
// v == 10
```

## Tests

Run tests:

```bash
go test -v
```

Run benchmarks:

```bash
go test -bench . -benchmem
```

The suite includes coverage for:

- `FromString` happy paths and failures
- All numeric widths and signed/unsigned boundaries
- Boolean parsing (multiple formats)
- Time parsing (20+ formats)
- Duration parsing (multiple input formats)
- `FromType` with slice targets
- Edge cases and error scenarios
- Performance benchmarks

## Versioning & Compatibility

- Go 1.20+ recommended (tested with Go 1.24)
- API is small and stable; changes will aim to be backward compatible

## License

MIT
