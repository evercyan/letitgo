package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	key := "adfIIe4la4i^f9kk"
	text := "123456"

	encrypt, err := AesEncrypt(text, key)
	assert.Equal(t, "sXxgEij7p+oTTqvOVoKbGQ==", encrypt)
	assert.Nil(t, err)

	decrypt, err := AesDecrypt(encrypt, key)
	assert.Equal(t, text, decrypt)
	assert.Nil(t, err)
}
