package cast_test

import (
	"testing"

	"github.com/shopspring/decimal"
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
	var val any
	var err error

	val, err = cast.Float32("3.14")
	assert.NoError(t, err)
	assert.Equal(t, float32(3.14), val)

	_, err = cast.Float32("str")
	assert.Error(t, err)

	val, err = cast.Float64("3.14")
	assert.NoError(t, err)
	assert.Equal(t, 3.14, val)

	_, err = cast.Float64("str")
	assert.Error(t, err)
}

func TestDecimal(t *testing.T) {
	var val decimal.Decimal
	var err error

	// Basic decimal tests
	val, err = cast.Decimal("123.45")
	assert.NoError(t, err)
	assert.Equal(t, "123.45", val.String())

	val, err = cast.Decimal("-456.789")
	assert.NoError(t, err)
	assert.Equal(t, "-456.789", val.String())

	val, err = cast.Decimal("0")
	assert.NoError(t, err)
	assert.Equal(t, "0", val.String())

	// Scientific notation
	val, err = cast.Decimal("1.23e4")
	assert.NoError(t, err)
	assert.Equal(t, "12300", val.String())

	// Error cases
	_, err = cast.Decimal("invalid")
	assert.Error(t, err)

	_, err = cast.Decimal("")
	assert.Error(t, err)
}

func TestNumberBoundaries(t *testing.T) {
	t.Run("int8 boundaries", func(t *testing.T) {
		// Max int8
		result, err := cast.Int8("127")
		assert.NoError(t, err)
		assert.Equal(t, int8(127), result)

		// Min int8
		result, err = cast.Int8("-128")
		assert.NoError(t, err)
		assert.Equal(t, int8(-128), result)

		// Overflow int8
		_, err = cast.Int8("128")
		assert.Error(t, err)
	})

	t.Run("uint8 boundaries", func(t *testing.T) {
		// Max uint8
		result, err := cast.Uint8("255")
		assert.NoError(t, err)
		assert.Equal(t, uint8(255), result)

		// Overflow uint8
		_, err = cast.Uint8("256")
		assert.Error(t, err)

		// Negative uint8
		_, err = cast.Uint8("-1")
		assert.Error(t, err)
	})

	t.Run("float32 boundaries", func(t *testing.T) {
		// Large float32
		result, err := cast.Float32("3.4028235e38")
		assert.NoError(t, err)
		assert.Equal(t, float32(3.4028235e38), result)

		// Very large number (will become infinity)
		result, err = cast.Float32("1e39")
		assert.NoError(t, err)
		assert.True(t, result > float32(3.4028235e38)) // Will be +inf
	})
}

func TestNumberBases(t *testing.T) {
	t.Run("hexadecimal", func(t *testing.T) {
		// Hex int
		result, err := cast.Int("0xFF")
		assert.NoError(t, err)
		assert.Equal(t, 255, result)

		// Hex int8
		result8, err := cast.Int8("0x7F")
		assert.NoError(t, err)
		assert.Equal(t, int8(127), result8)

		// Hex int16
		result16, err := cast.Int16("0xFF")
		assert.NoError(t, err)
		assert.Equal(t, int16(255), result16)
	})

	t.Run("octal", func(t *testing.T) {
		// Octal int
		result, err := cast.Int("0755")
		assert.NoError(t, err)
		assert.Equal(t, 493, result)

		// Octal int8
		result8, err := cast.Int8("077")
		assert.NoError(t, err)
		assert.Equal(t, int8(63), result8)
	})

	t.Run("binary", func(t *testing.T) {
		// Binary int
		result, err := cast.Int("0b1010")
		assert.NoError(t, err)
		assert.Equal(t, 10, result)

		// Binary int8
		result8, err := cast.Int8("0b101")
		assert.NoError(t, err)
		assert.Equal(t, int8(5), result8)
	})

	t.Run("decimal with prefix", func(t *testing.T) {
		// Regular decimal
		result, err := cast.Int("123")
		assert.NoError(t, err)
		assert.Equal(t, 123, result)
	})
}

// Benchmark tests
func BenchmarkInt(b *testing.B) {
	for b.Loop() {
		cast.Int("123")
	}
}

func BenchmarkFloat(b *testing.B) {
	for b.Loop() {
		cast.Float64("3.14")
	}
}

func BenchmarkDecimal(b *testing.B) {
	for b.Loop() {
		cast.Decimal("123.45")
	}
}
