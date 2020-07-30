package request

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func skip(t *testing.T, err error) {
	if err != nil {
		t.Skip()
	}
}

func TestGet(t *testing.T) {
	resp, err := Get("http://127.0.0.1/letitgo.php?method=get")

	skip(t, err)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"code":0,"method":"get"}`, resp)
}

func TestPost(t *testing.T) {
	param := url.Values{"method": {"post"}}
	resp, err := Post("http://127.0.0.1/letitgo.php", param.Encode())

	skip(t, err)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"code":0,"method":"post"}`, resp)
}

func TestJsonPost(t *testing.T) {
	param := map[string]string{
		"method": "jsonpost",
	}
	bytes, _ := json.Marshal(param)
	resp, err := JsonPost("http://127.0.0.1/letitgo.php", string(bytes))

	skip(t, err)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"code":0,"method":"jsonpost"}`, resp)
}
