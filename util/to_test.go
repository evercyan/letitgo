package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUint(t *testing.T) {
	assert.Equal(t, uint64(1), ToUint("1"))
	assert.Equal(t, uint64(0), ToUint(""))
	assert.Equal(t, uint64(3), ToUint(3))
}

func TestToBool(t *testing.T) {
	assert.True(t, ToBool(true))
	assert.False(t, ToBool("abc"))
	assert.False(t, ToBool("0"))
	assert.True(t, ToBool("1"))
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
