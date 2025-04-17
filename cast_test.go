package cast_test

import (
	"reflect"
	"testing"

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

	// else

	_, err = cast.FromString("0", "invalid")
	assert.Error(t, err, "cast: type %v is not supported", "invalid")

	_, err = cast.FromString("0,1", "[]invalid")
	assert.Error(t, err, "cast: type %v is not supported", "[]invalid")
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
}
