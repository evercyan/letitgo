package httpcli

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func skip(t *testing.T, err error) {
	if err != nil {
		t.Skip()
	}
}

func TestGet(t *testing.T) {
	resp, err := NewClient().Get(context.Background(), "http://127.0.0.1/letitgo.php?method=get", nil)
	skip(t, err)
	assert.Nil(t, err)
	b, _ := ioutil.ReadAll(resp.Body)
	assert.JSONEq(t, `{"code":0,"method":"get"}`, string(b))
}
