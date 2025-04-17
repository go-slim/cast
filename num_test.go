package cast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-slim.dev/cast"
)

func TestInt(t *testing.T) {
	message := "cast: cannot cast `%v` to type `%v`"

	var val any
	var err error

	val, err = cast.Int("1")
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	_, err = cast.Int("str")
	assert.Errorf(t, err, message, "str", "int")

	val, err = cast.Int8("1")
	assert.NoError(t, err)
	assert.Equal(t, int8(1), val)

	_, err = cast.Int8("str")
	assert.Errorf(t, err, message, "str", "int8")

	val, err = cast.Int16("1")
	assert.NoError(t, err)
	assert.Equal(t, int16(1), val)

	_, err = cast.Int16("str")
	assert.Errorf(t, err, message, "str", "int16")

	val, err = cast.Int32("1")
	assert.NoError(t, err)
	assert.Equal(t, int32(1), val)

	_, err = cast.Int32("str")
	assert.Errorf(t, err, message, "str", "int32")

	val, err = cast.Int64("1")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	_, err = cast.Int64("str")
	assert.Errorf(t, err, message, "str", "int64")
}

func TestUint(t *testing.T) {
	message := "cast: cannot cast `%v` to type `%v`"

	var val any
	var err error

	val, err = cast.Uint("1")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), val)

	_, err = cast.Uint("-1")
	assert.Errorf(t, err, message, "-1", "uint")

	val, err = cast.Uint8("1")
	assert.NoError(t, err)
	assert.Equal(t, uint8(1), val)

	_, err = cast.Uint8("-1")
	assert.Errorf(t, err, message, "-1", "uint8")

	val, err = cast.Uint16("1")
	assert.NoError(t, err)
	assert.Equal(t, uint16(1), val)

	_, err = cast.Uint16("-1")
	assert.Errorf(t, err, message, "-1", "uint16")

	val, err = cast.Uint32("1")
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), val)

	_, err = cast.Uint32("-1")
	assert.Errorf(t, err, message, "-1", "uint32")

	val, err = cast.Uint64("1")
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), val)

	_, err = cast.Uint64("-1")
	assert.Errorf(t, err, message, "-1", "uint64")
}

func TestFloat(t *testing.T) {
	message := "cast: cannot cast `%v` to type `%v`"

	var val any
	var err error

	val, err = cast.Float32("3.14")
	assert.NoError(t, err)
	assert.Equal(t, float32(3.14), val)

	_, err = cast.Float32("str")
	assert.Errorf(t, err, message, "str", "float32")

	val, err = cast.Float64("3.14")
	assert.NoError(t, err)
	assert.Equal(t, 3.14, val)

	_, err = cast.Float64("str")
	assert.Errorf(t, err, message, "str", "float64")
}
