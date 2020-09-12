package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	assert.True(t, IsInt(1))
	assert.True(t, IsUint(uint32(1)))
	assert.True(t, IsFloat(float32(1)))

	assert.True(t, IsNumeric(1))
	assert.True(t, IsNumeric(uint32(1)))
	assert.True(t, IsNumeric(float32(1)))

	assert.True(t, IsBool(true))
	assert.False(t, IsBool("1"))

	assert.True(t, IsString("hello"))
	assert.False(t, IsString(1))

	assert.True(t, IsSlice([]int{1, 2, 3}))
	assert.True(t, IsArray([3]int{1, 2, 3}))
	assert.True(t, IsMap(make(map[string]string)))
	assert.True(t, IsChannel(make(chan string)))

	assert.True(t, IsTime(time.Now()))
	assert.False(t, IsTime("1"))

	elemStruct := struct {
		Name string
	}{}
	assert.True(t, IsStruct(elemStruct))

	elemFunc := func() {}
	assert.True(t, IsFunc(elemFunc))
}

func TestIsEmpty(t *testing.T) {
	assert.True(t, IsEmpty(""))
	assert.True(t, IsEmpty(0))
	assert.True(t, IsEmpty(nil))
	assert.True(t, IsEmpty(false))
	assert.False(t, IsEmpty("0"))
}

func TestIsEmail(t *testing.T) {
	assert.True(t, IsEmail("evercyan@qq.com"))
	assert.False(t, IsEmail("evercyan"))
}

func TestIsURL(t *testing.T) {
	assert.True(t, IsURL("http://t.cn"))
	assert.False(t, IsURL(""))
	assert.False(t, IsURL("h"))
	assert.False(t, IsURL("http"))
	assert.False(t, IsURL("http://.t.cn"))
	assert.False(t, IsURL(".http://.t.cn"))
}

func TestIsJson(t *testing.T) {
	assert.True(t, IsJson("{\"Title\":\"AAA\",\"Text\":\"BBB\"}"))
	assert.True(t, IsJson("[1, 2, 3]"))
	assert.False(t, IsJson("["))
}

func TestIsIP(t *testing.T) {
	assert.True(t, IsIP("127.0.0.1"))
	assert.False(t, IsIP("127.0.0.256"))
}

func TestIsMobile(t *testing.T) {
	assert.True(t, IsMobile("18500000000"))
	assert.False(t, IsMobile("134"))
}

func TestInArray(t *testing.T) {
	assert.True(t, InArray(1, []int{1, 2, 3}))
	assert.False(t, InArray(4, []int{1, 2, 3}))

	assert.True(t, InArray(1, map[int]int{1: 1, 2: 2}))
	assert.False(t, InArray(3, map[int]int{1: 1, 2: 2}))
}
