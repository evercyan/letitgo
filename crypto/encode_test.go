package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Encode(t *testing.T) {
	assert.Equal(t, "MTExMTEx", Base64Encode("111111"))
}
func TestBase64Decode(t *testing.T) {
	assert.Equal(t, "111111", Base64Decode("MTExMTEx"))
	assert.Empty(t, Base64Decode("1"))
}
func TestUrlEncode(t *testing.T) {
	assert.Equal(t, "http%3A%2F%2Ft.cn", UrlEncode("http://t.cn"))
}
func TestUrlDecode(t *testing.T) {
	assert.Equal(t, "http://t.cn", UrlDecode("http%3A%2F%2Ft.cn"))
	assert.Empty(t, UrlDecode("%"))
}

func TestJsonEncode(t *testing.T) {
	s := struct {
		Title string
		Text  string
	}{
		Title: "AAA",
		Text:  "BBB",
	}
	assert.JSONEq(t, `{"Text":"BBB","Title":"AAA"}`, JsonEncode(s))
}
