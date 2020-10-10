package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageType(t *testing.T) {
	assert.Equal(t, "png", ImageType("../letitgo.png"))
	assert.Empty(t, ImageType("../letitgo"))
	assert.Empty(t, ImageType("../README.md"))
}
