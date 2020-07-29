package crypto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	key := "sdflIerl34i^flkj"
	expire := 24 * 3600
	jwt := NewJWT(key, expire)

	payload := map[string]interface{}{
		"name": "hello",
	}

	token, err := jwt.GetToken(payload)
	assert.NotEmpty(t, token)
	assert.Nil(t, err)

	resp, err := jwt.ParseToken(token)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
}

func TestJWTCoverage(t *testing.T) {
	key := "sdflIerl34i^flkj"
	jwt := NewJWT(key, 100)

	payload := map[string]interface{}{
		"name": "hello",
		"nbf":  time.Now().Add(time.Duration(1000) * time.Second).Unix(),
	}

	token, err := jwt.GetToken(payload)
	assert.Nil(t, err)

	resp, err := jwt.ParseToken(token)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}
