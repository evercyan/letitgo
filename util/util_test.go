package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "96e79218965eb72c92a549dd5a330112", Md5("111111"))
}

func TestGetClientIp(t *testing.T) {
	assert.NotEmpty(t, GetClientIp())
}

func TestGuid(t *testing.T) {
	assert.NotEmpty(t, Guid())
}

func TestRand(t *testing.T) {
	assert.Equal(t, 0, Rand(2, 1))
	assert.LessOrEqual(t, 1, Rand(1, 10))
	assert.GreaterOrEqual(t, 10, Rand(1, 10))

	assert.LessOrEqual(t, -10, Rand(-10, 10))
	assert.GreaterOrEqual(t, 10, Rand(-10, 10))
}

func TestRanger(t *testing.T) {
	assert.Equal(t, []int(nil), Range(2, 1))
	assert.Equal(t, []int{1, 2}, Range(1, 2))
}
