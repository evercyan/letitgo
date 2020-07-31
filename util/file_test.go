package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExist(t *testing.T) {
	assert.True(t, IsExist("./is.go"))
	assert.False(t, IsExist("./is_t.go"))
}

func TestIsFile(t *testing.T) {
	assert.True(t, IsFile("./is.go"))
	assert.False(t, IsFile("../util"))
}

func TestIsDir(t *testing.T) {
	assert.True(t, IsDir("../util"))
	assert.False(t, IsDir("./is.go"))
}

func TestGetSize(t *testing.T) {
	assert.Equal(t, int64(0), GetSize("./is_t.go"))
	assert.LessOrEqual(t, int64(0), GetSize("./is.go"))
}

func TestReadFile(t *testing.T) {
	assert.NotEmpty(t, ReadFile("./is.go"))
}

func TestWriteFile(t *testing.T) {
	file := "/tmp/letitgo.tmp"
	content := "letitgo"
	assert.Nil(t, WriteFile(file, content))
	assert.Equal(t, content, ReadFile(file))
}
