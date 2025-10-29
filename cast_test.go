package cast_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go-slim.dev/cast"
)

func TestFromString(t *testing.T) {
	message := "cast: cannot cast `%v` to type `%v`"

	// string

	val, err := cast.FromString("string", "string")
	assert.NoError(t, err)
	assert.Equal(t, "string", val)

	// int

	val, err = cast.FromString("1", "int")
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	_, err = cast.FromString("str", "int")
	assert.Errorf(t, err, message, "str", "int")

	val, err = cast.FromString("1", "int8")
	assert.NoError(t, err)
	assert.Equal(t, int8(1), val)

	_, err = cast.FromString("str", "int8")
	assert.Errorf(t, err, message, "str", "int8")

	val, err = cast.FromString("1", "int16")
	assert.NoError(t, err)
	assert.Equal(t, int16(1), val)

	_, err = cast.FromString("str", "int16")
	assert.Errorf(t, err, message, "str", "int16")

	val, err = cast.FromString("1", "int32")
	assert.NoError(t, err)
	assert.Equal(t, int32(1), val)

	_, err = cast.FromString("str", "int32")
	assert.Errorf(t, err, message, "str", "int32")

	val, err = cast.FromString("1", "int64")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	_, err = cast.FromString("str", "int64")
	assert.Errorf(t, err, message, "str", "int64")

	// uint

	val, err = cast.FromString("1", "uint")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), val)

	_, err = cast.FromString("-1", "uint")
	assert.Errorf(t, err, message, "-1", "uint")

	val, err = cast.FromString("1", "uint8")
	assert.NoError(t, err)
	assert.Equal(t, uint8(1), val)

	_, err = cast.FromString("-1", "uint8")
	assert.Errorf(t, err, message, "-1", "uint8")

	val, err = cast.FromString("1", "uint16")
	assert.NoError(t, err)
	assert.Equal(t, uint16(1), val)

	_, err = cast.FromString("-1", "uint16")
	assert.Errorf(t, err, message, "-1", "uint16")

	val, err = cast.FromString("1", "uint32")
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), val)

	_, err = cast.FromString("-1", "uint32")
	assert.Errorf(t, err, message, "-1", "uint32")

	val, err = cast.FromString("1", "uint64")
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), val)

	_, err = cast.FromString("-1", "uint64")
	assert.Errorf(t, err, message, "-1", "uint64")

	// float

	val, err = cast.FromString("3.14", "float32")
	assert.NoError(t, err)
	assert.Equal(t, float32(3.14), val)

	_, err = cast.FromString("str", "float32")
	assert.Errorf(t, err, message, "str", "float32")

	val, err = cast.FromString("3.14", "float64")
	assert.NoError(t, err)
	assert.Equal(t, 3.14, val)

	_, err = cast.FromString("str", "float64")
	assert.Errorf(t, err, message, "str", "float64")

	// bool

	val, err = cast.FromString("true", "bool")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	_, err = cast.FromString("1", "bool")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = cast.FromString("false", "bool")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	_, err = cast.FromString("0", "bool")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	_, err = cast.FromString("invalid", "bool")
	assert.Errorf(t, err, message, "invalid", "bool")

	// time.Time
	val, err = cast.FromString("2023-12-25T10:30:45Z", "time.Time")
	assert.NoError(t, err)
	expectedTime, _ := time.Parse(time.RFC3339, "2023-12-25T10:30:45Z")
	assert.Equal(t, expectedTime, val)

	val, err = cast.FromString("1703505045", "time.Time")
	assert.NoError(t, err)
	assert.Equal(t, time.Unix(1703505045, 0), val)

	_, err = cast.FromString("invalid-time", "time.Time")
	assert.Error(t, err)

	// time.Duration
	val, err = cast.FromString("1h30m45s", "time.Duration")
	assert.NoError(t, err)
	assert.Equal(t, 1*time.Hour+30*time.Minute+45*time.Second, val)

	val, err = cast.FromString("5000000000", "time.Duration")
	assert.NoError(t, err)
	assert.Equal(t, 5*time.Second, val)

	val, err = cast.FromString("1.5", "time.Duration")
	assert.NoError(t, err)
	assert.Equal(t, 1500*time.Millisecond, val)

	_, err = cast.FromString("invalid-duration", "time.Duration")
	assert.Error(t, err)

	// else

	_, err = cast.FromString("0", "invalid")
	assert.Error(t, err)

	_, err = cast.FromString("0,1", "[]invalid")
	assert.Error(t, err)
}

func TestFromType(t *testing.T) {
	val, err := cast.FromType("4,-5,6", reflect.TypeOf([]int{}))
	assert.NoError(t, err)
	assert.Equal(t, []int{4, -5, 6}, val)

	val, err = cast.FromType("4,-5,6", reflect.TypeOf([]int8{}))
	assert.NoError(t, err)
	assert.Equal(t, []int8{4, -5, 6}, val)

	val, err = cast.FromType("4,-5, 6", reflect.TypeOf([]int16{}))
	assert.NoError(t, err)
	assert.Equal(t, []int16{4, -5, 6}, val)

	val, err = cast.FromType("4,-5,6", reflect.TypeOf([]int32{}))
	assert.NoError(t, err)
	assert.Equal(t, []int32{4, -5, 6}, val)

	val, err = cast.FromType("4,-5,6", reflect.TypeOf([]int64{}))
	assert.NoError(t, err)
	assert.Equal(t, []int64{4, -5, 6}, val)

	val, err = cast.FromType("4,5,6", reflect.TypeOf([]uint{}))
	assert.NoError(t, err)
	assert.Equal(t, []uint{4, 5, 6}, val)

	val, err = cast.FromType("4,5,6", reflect.TypeOf([]uint8{}))
	assert.NoError(t, err)
	assert.Equal(t, []uint8{4, 5, 6}, val)

	val, err = cast.FromType("4,5,6", reflect.TypeOf([]uint16{}))
	assert.NoError(t, err)
	assert.Equal(t, []uint16{4, 5, 6}, val)

	val, err = cast.FromType("4,5,6", reflect.TypeOf([]uint32{}))
	assert.NoError(t, err)
	assert.Equal(t, []uint32{4, 5, 6}, val)

	val, err = cast.FromType("4,5,6", reflect.TypeOf([]uint64{}))
	assert.NoError(t, err)
	assert.Equal(t, []uint64{4, 5, 6}, val)

	val, err = cast.FromType("3.14,9.8", reflect.TypeOf([]float32{}))
	assert.NoError(t, err)
	assert.Equal(t, []float32{3.14, 9.8}, val)

	val, err = cast.FromType("3.14,9.8", reflect.TypeOf([]float64{}))
	assert.NoError(t, err)
	assert.Equal(t, []float64{3.14, 9.8}, val)

	val, err = cast.FromType("true,false,0,1", reflect.TypeOf([]bool{}))
	assert.NoError(t, err)
	assert.Equal(t, []bool{true, false, false, true}, val)

	val, err = cast.FromType("a, b, c", reflect.TypeOf([]string{}))
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, val)

	val, err = cast.FromType("simple", reflect.TypeOf("string"))
	assert.NoError(t, err)
	assert.Equal(t, "simple", val)

	val, err = cast.FromType("a,b,c", reflect.TypeOf([]int{}))
	assert.Error(t, err)

	// Test time arrays
	val, err = cast.FromType("2023-12-25T10:30:45Z,2023-12-26T11:45:00Z", reflect.TypeOf([]time.Time{}))
	assert.NoError(t, err)
	expectedTimes := []time.Time{}
	t1, _ := time.Parse(time.RFC3339, "2023-12-25T10:30:45Z")
	t2, _ := time.Parse(time.RFC3339, "2023-12-26T11:45:00Z")
	expectedTimes = append(expectedTimes, t1, t2)
	assert.Equal(t, expectedTimes, val)

	// Test duration arrays
	val, err = cast.FromType("1h30m,2h45m", reflect.TypeOf([]time.Duration{}))
	assert.NoError(t, err)
	expectedDurations := []time.Duration{1*time.Hour + 30*time.Minute, 2*time.Hour + 45*time.Minute}
	assert.Equal(t, expectedDurations, val)
}

func TestEdgeCases(t *testing.T) {
	// Test empty arrays
	val, err := cast.FromType("", reflect.TypeOf([]string{}))
	assert.NoError(t, err)
	assert.Equal(t, []string{""}, val)

	// Test arrays with extra spaces
	val, err = cast.FromType("  a ,  b  , c  ", reflect.TypeOf([]string{}))
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, val)

	// Test arrays with newline characters (should be treated as single element since only comma splits)
	val, err = cast.FromType("a\nb\nc", reflect.TypeOf([]string{}))
	assert.NoError(t, err)
	assert.Equal(t, []string{"a\nb\nc"}, val)

	// Test arrays with newline and comma combination
	val, err = cast.FromType("a,b\nc", reflect.TypeOf([]string{}))
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b\nc"}, val)

	// Test arrays with various whitespace
	val, err = cast.FromType("a\t,\nb\r,\tc", reflect.TypeOf([]string{}))
	assert.NoError(t, err)
	assert.Equal(t, []string{"a\t", "b", "\tc"}, val)
}
