package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasIP(t *testing.T) {
	assert.True(t, HasIP("abc127.0.0.1"))
	assert.False(t, HasIP("abc127.0"))
}
