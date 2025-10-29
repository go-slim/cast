package cast

import (
	"fmt"
	"reflect"
	"strings"
)

// FromType casts a string value to the given reflected type.
// For slice types (e.g., []int, []string), the input string is split by commas,
// and each item is trimmed and converted to the element type.
//
// Example:
//
//	t := reflect.TypeOf([]int(nil))
//	result, err := FromType("1,2,3", t)
//	// result == []int{1, 2, 3}
func FromType(s string, targetType reflect.Type) (any, error) {
	var typeName = targetType.String()

	if strings.HasPrefix(typeName, "[]") {
		itemType := typeName[2:]
		array := reflect.New(targetType).Elem()

		for v := range strings.SplitSeq(s, ",") {
			if item, err := FromString(strings.Trim(v, " \n\r"), itemType); err != nil {
				return array.Interface(), err
			} else {
				array = reflect.Append(array, reflect.ValueOf(item))
			}
		}

		return array.Interface(), nil
	}

	return FromString(s, typeName)
}

// FromString casts a string value to the given type name.
// Supported type names include all Go numeric types (int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64, float32, float64), bool, string,
// and time types (time.Time, time.Duration).
//
// Returns an error if the type name is not supported or if parsing fails.
//
// Example:
//
//	v, err := FromString("42", "int")
//	v, err := FromString("true", "bool")
//	v, err := FromString("2023-12-25T10:30:45Z", "time.Time")
func FromString(s string, targetType string) (any, error) {
	switch targetType {
	case "int":
		return Int(s)
	case "int8":
		return Int8(s)
	case "int16":
		return Int16(s)
	case "int32":
		return Int32(s)
	case "int64":
		return Int64(s)
	case "uint":
		return Uint(s)
	case "uint8":
		return Uint8(s)
	case "uint16":
		return Uint16(s)
	case "uint32":
		return Uint32(s)
	case "uint64":
		return Uint64(s)
	case "bool":
		return Bool(s)
	case "float32":
		return Float32(s)
	case "float64":
		return Float64(s)
	case "string":
		return s, nil
	case "time.Time":
		return Time(s)
	case "time.Duration":
		return Duration(s)
	default:
		return nil, fmt.Errorf("cast: type %v is not supported", targetType)
	}
}
