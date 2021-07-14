package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUint(t *testing.T) {
	assert.Equal(t, uint64(1), ToUint("1"))
	assert.Equal(t, uint64(0), ToUint(""))
	assert.Equal(t, uint64(3), ToUint(3))
	var num interface{}
	assert.Equal(t, uint64(0), ToUint(num))
	assert.Equal(t, uint64(1), ToUint(true))
	assert.Equal(t, uint64(0), ToUint(false))
	assert.Equal(t, uint64(1), ToUint(float32(1)))
	assert.Equal(t, uint64(1), ToUint(float64(1)))
	assert.Equal(t, uint64(1), ToUint(int(1)))
	assert.Equal(t, uint64(1), ToUint(int8(1)))
	assert.Equal(t, uint64(1), ToUint(int16(1)))
	assert.Equal(t, uint64(1), ToUint(int32(1)))
	assert.Equal(t, uint64(1), ToUint(int64(1)))
	assert.Equal(t, uint64(1), ToUint(uint(1)))
	assert.Equal(t, uint64(1), ToUint(uint8(1)))
	assert.Equal(t, uint64(1), ToUint(uint16(1)))
	assert.Equal(t, uint64(1), ToUint(uint32(1)))
	assert.Equal(t, uint64(1), ToUint(uint64(1)))
	assert.Equal(t, uint64(0), ToUint("abc"))

	obj := make([]int, 0)
	assert.Equal(t, uint64(0), ToUint(obj))
}

func TestToBool(t *testing.T) {
	assert.True(t, ToBool(true))
	assert.False(t, ToBool("abc"))
	assert.False(t, ToBool("0"))
	assert.True(t, ToBool("1"))

	var num interface{}
	assert.False(t, ToBool(num))
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

type A1 struct {
	Name string
}

type A2 struct {
	Name string
}

func TestToStruct(t *testing.T) {
	a1 := &A1{
		Name: "hello",
	}
	a2 := &A2{}
	err := ToStruct(a1, a2)
	assert.Nil(t, err)
	assert.Equal(t, "hello", a2.Name)
}
