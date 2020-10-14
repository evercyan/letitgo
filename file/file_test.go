package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	assert.Equal(t, int64(0), Size("./file_t.go"))
	assert.LessOrEqual(t, int64(0), Size("./file.go"))
}

func TestSizeText(t *testing.T) {
	assert.Equal(t, "0.00B", SizeText(0))
	assert.Equal(t, "1023.00B", SizeText(1023))
	assert.Equal(t, "1.00KB", SizeText(1024))
	assert.Equal(t, "1.65KB", SizeText(1024+666))
	assert.Equal(t, "1.65MB", SizeText((1024+666)*1024))
	assert.Equal(t, "1.65GB", SizeText((1024+666)*1024*1024))
	assert.Equal(t, "1.00TB", SizeText(1*1024*1024*1024*1024))
}

func TestRead(t *testing.T) {
	assert.NotEmpty(t, Read("./file.go"))
	assert.Empty(t, Read("./file_t.go"))
}

func TestWrite(t *testing.T) {
	file := "/tmp/letitgo.tmp"
	content := "letitgo"
	assert.Nil(t, Write(file, content))
	assert.Equal(t, content, Read(file))
}

func TestExt(t *testing.T) {
	assert.Equal(t, ".go", Ext("./file.go"))
	assert.Empty(t, Ext("./file_t"))
}

func TestLine(t *testing.T) {
	assert.Equal(t, 0, LineCount("../LICENSES"))
	assert.Equal(t, map[int]string{}, LineContent("../LICENSES"))

	assert.Equal(t, 21, LineCount("../LICENSE"))
	assert.Equal(t, map[int]string{
		1: "MIT License",
		2: "",
		3: "Copyright (c) 2020 严宇川",
	}, LineContent("../LICENSE", 1, 2, 3))
}
