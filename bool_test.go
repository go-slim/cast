package cast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-slim.dev/cast"
)

func TestBool(t *testing.T) {
	val, err := cast.Bool("true")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	_, err = cast.Bool("1")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = cast.Bool("false")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	_, err = cast.Bool("0")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	_, err = cast.Bool("invalid")
	assert.Errorf(t, err, "cast: cannot cast `%v` to type `bool`", "invalid")
}
