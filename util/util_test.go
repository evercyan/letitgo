package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "96e79218965eb72c92a549dd5a330112", Md5("111111"))
}

func TestTimestamp(t *testing.T) {
	assert.NotEqual(t, 0, Timestamp())
}
func TestGetClientIp(t *testing.T) {
	assert.NotEmpty(t, GetClientIp())
}

func TestGuid(t *testing.T) {
	assert.NotEmpty(t, Guid())
}

func TestRand(t *testing.T) {
	assert.Equal(t, Rand(10, 10), int64(10))
	assert.LessOrEqual(t, Rand(10, 20), int64(20))
	assert.GreaterOrEqual(t, Rand(10, 20), int64(10))
}
