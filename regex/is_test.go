package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIP(t *testing.T) {
	assert.True(t, IsIPV4("127.0.0.1"))
	assert.True(t, IsIPV4("255.255.255.255"))
	assert.False(t, IsIPV4("256.255.255.255"))

	assert.False(t, IsIPV6("127.0.0.1"))
	assert.True(t, IsIPV6("2001:0db8:3c4d:0015:0000:0000:1a2f:1a2b"))

	assert.True(t, IsIP("127.0.0.1"))
}

func TestIsMacAddress(t *testing.T) {
	assert.True(t, IsMacAddress("6a:00:02:9c:1f:30"))
	assert.False(t, IsMacAddress(""))
}

func TestIsPhone(t *testing.T) {
	assert.True(t, IsPhone("18500000000"))
	assert.False(t, IsPhone("134"))
}
