package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt64(t *testing.T) {
	assert.Equal(t, int64(123), ToInt64("123"))
	assert.Equal(t, int64(0), ToInt64("abc"))
}

func TestToBool(t *testing.T) {
	assert.True(t, ToBool("1"))
	assert.False(t, ToBool("abc"))
	assert.False(t, ToBool("0"))
}

func TestToString(t *testing.T) {
	assert.Equal(t, "1234", ToString(1234))
	assert.Equal(t, "[1 2 3 4]", ToString([]int{1, 2, 3, 4}))
}

func TestToCamelCase(t *testing.T) {
	assert.Equal(t, "UserName", ToCamelCase("user_name"))
}

func TestToSnakeCase(t *testing.T) {
	assert.Equal(t, "user_name", ToSnakeCase("userName"))
	assert.Equal(t, "user_name", ToSnakeCase("UserName"))
}
