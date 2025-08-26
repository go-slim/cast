# cast

A minimal, dependency-free casting utility for Go strings to native types.

Module path: `go-slim.dev/cast`

- Convert string input to Go numeric/boolean/string types.
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

    cast "go-slim.dev/cast"
)

func main() {
    v, _ := cast.FromString("42", "int")
    fmt.Printf("%T %v\n", v, v) // int 42

    v, _ = cast.FromString("true", "bool")
    fmt.Printf("%T %v\n", v, v) // bool true

    // Using reflect.Type, including slices (comma-separated)
    t := reflect.TypeOf([]int(nil)) // []int
    v, _ = cast.FromType("1,2,3", t)
    fmt.Printf("%T %v\n", v, v) // []int [1 2 3]
}
```

## API

- Core entry points
  - `FromString(s string, targetType string) (any, error)`
    - Supported type names: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `bool`, `float32`, `float64`, `string`.
  - `FromType(s string, targetType reflect.Type) (any, error)`
    - If `targetType` is a slice (e.g., `[]int`), input is split by comma `,`, items are trimmed, and each is cast via `FromString` to the slice element type.

- Numeric helpers (string → typed value)
  - `Int`, `Int8`, `Int16`, `Int32`, `Int64`
  - `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
  - `Float32`, `Float64`
- Boolean helper
  - `Bool`

See implementations in:
- `cast/cast.go` (dispatchers `FromString`, `FromType`)
- `cast/num.go` (numeric/boolean parsing)

## Error Behavior

- On unsupported type names, `FromString` returns:
  - `fmt.Errorf("cast: type %v is not supported", targetType)`
- On parse failures, helpers return parse errors from `strconv` (except `Int`, which formats a message like `cast: cannot cast \\"%v\\" to type \\"%v\\"`).
- `FromType` with slice types stops and returns the first encountered item error.

## Examples

```go
v, err := cast.FromString("1", "uint16")
// v == uint16(1), err == nil

v, err = cast.FromString("str", "int")
// err != nil (cannot cast)

// Slices via FromType
t := reflect.TypeOf([]uint8(nil))
v, err = cast.FromType("1, 2, 255", t)
// v == []uint8{1,2,255}, err == nil
```

## Tests

Run tests:

```bash
go test -v
```

The suite includes coverage for:
- `FromString` happy paths and failures
- All numeric widths and signed/unsigned boundaries
- Boolean parsing
- `FromType` with slice targets

## Versioning & Compatibility

- Go 1.20+ recommended (tested with Go 1.24)
- API is small and stable; changes will aim to be backward compatible

## License

MIT
