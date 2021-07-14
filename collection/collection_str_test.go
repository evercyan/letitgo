package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getString() []string {
	return []string{"a", "b", "c", "d"}
}

func TestCollectionStringBase(t *testing.T) {
	assert.Equal(t, 4, Collect(getString()).Length())
	assert.Equal(t, getString(), Collect(getString()).Value().([]string))
	assert.Equal(t, "[\"a\",\"b\",\"c\",\"d\"]", Collect(getString()).Json())
}

func TestCollectionStringJoin(t *testing.T) {
	assert.Equal(t, "a-b-c-d", Collect(getString()).Join("-"))
}

func TestCollectionStringContains(t *testing.T) {
	assert.True(t, Collect(getString()).Contains("a"))
	assert.False(t, Collect(getString()).Contains("e"))
}

func TestCollectionStringUnique(t *testing.T) {
	list := []string{"a", "a", "b", "c", "d"}
	expected := []string{"a", "b", "c", "d"}
	assert.Equal(t, expected, Collect(list).Unique().Value().([]string))
}

func TestCollectionStringDel(t *testing.T) {
	expected := []string{"a", "b", "d"}
	assert.Equal(t, expected, Collect(getString()).DelKey(2).Value().([]string))
	assert.Equal(t, expected, Collect(getString()).DelValue("c").Value().([]string))
}

func TestCollectionStringFilter(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, Collect(getString()).Filter(func(key, value interface{}) bool {
		return value.(string) != "d"
	}).Value().([]string))
}

func TestCollectionStringCoverage(t *testing.T) {
	assert.Equal(t, 4, Collect(getString()).DelKey(4).Length())
	assert.Equal(t, 3, Collect(getString()).DelKey(3).Length())
	assert.Equal(t, 4, Collect(getString()).DelValue("f").Length())
}
