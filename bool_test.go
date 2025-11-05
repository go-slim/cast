package cast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-slim.dev/cast"
)

func TestBool(t *testing.T) {
	// Test true values
	val, err := cast.Bool("true")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = cast.Bool("1")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = cast.Bool("t")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = cast.Bool("TRUE")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = cast.Bool("True")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	// Test false values
	val, err = cast.Bool("false")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	val, err = cast.Bool("0")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	val, err = cast.Bool("f")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	val, err = cast.Bool("FALSE")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	val, err = cast.Bool("False")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	// Test error cases
	_, err = cast.Bool("invalid")
	assert.Error(t, err)

	_, err = cast.Bool("")
	assert.Error(t, err)

	_, err = cast.Bool("2")
	assert.Error(t, err)

	_, err = cast.Bool("-1")
	assert.Error(t, err)
}

func BenchmarkBool(b *testing.B) {
	testCases := []string{"true", "false", "1", "0", "TRUE", "FALSE"}

	for _, tc := range testCases {
		b.Run(tc, func(b *testing.B) {
			for b.Loop() {
				cast.Bool(tc)
			}
		})
	}
}
