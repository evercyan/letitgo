package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getMap() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"a": "a1",
			"b": 1,
			"c": int64(11),
		},
		{
			"a": "a2",
			"b": 2,
			"c": int64(22),
		},
	}
}

func TestCollectionMapBase(t *testing.T) {
	assert.Equal(t, 2, Collect(getMap()).Length())
	assert.Equal(t, getMap(), Collect(getMap()).Value().([]map[string]interface{}))
	assert.Equal(t, "[{\"a\":\"a1\",\"b\":1,\"c\":11},{\"a\":\"a2\",\"b\":2,\"c\":22}]", Collect(getMap()).Json())
}

func TestCollectionMapDel(t *testing.T) {
	expected := []map[string]interface{}{
		{
			"a": "a1",
			"b": 1,
			"c": int64(11),
		},
	}
	assert.Equal(t, expected, Collect(getMap()).DelKey(1).Value().([]map[string]interface{}))
	assert.Equal(t, expected, Collect(getMap()).DelKeyValue("a", "a2").Value().([]map[string]interface{}))
}

func TestCollectionMapPluck(t *testing.T) {
	assert.Equal(t, []string{"a1", "a2"}, Collect(getMap()).Pluck("a").Value().([]string))
	assert.Equal(t, []float64{1, 2}, Collect(getMap()).Pluck("b").Value().([]float64))
	assert.Equal(t, []float64{11, 22}, Collect(getMap()).Pluck("c").Value().([]float64))
}

func TestCollectionMapFilter(t *testing.T) {
	expected1 := []map[string]interface{}{
		{
			"a": "a1",
			"b": 1,
			"c": int64(11),
		},
	}
	assert.Equal(t, expected1, Collect(getMap()).Filter(func(key, value interface{}) bool {
		return value.(map[string]interface{})["b"] == 1
	}).Value().([]map[string]interface{}))
}

func TestCollectionMapGroupBy(t *testing.T) {
	expected := map[string]interface{}{
		"a1": map[string]interface{}{
			"a": "a1",
			"b": 1,
			"c": int64(11),
		},
		"a2": map[string]interface{}{
			"a": "a2",
			"b": 2,
			"c": int64(22),
		},
	}
	assert.Equal(t, expected, Collect(getMap()).GroupBy("a"))
	assert.Empty(t, Collect(getMap()).GroupBy("b"))
}

func TestCollectionMapCoverage(t *testing.T) {
	assert.Equal(t, 2, Collect(getMap()).DelKey(2).Length())
	assert.Equal(t, 1, Collect(getMap()).DelKey(0).Length())
	assert.Equal(t, 2, Collect(getMap()).DelKeyValue("a", "a2222").Length())
	assert.Equal(t, map[string]interface{}{}, Collect([]map[string]interface{}{}).GroupBy("a"))
	assert.Equal(t, map[string]interface{}{}, Collect(getMap()).GroupBy("aaaaa"))
	assert.Equal(t, map[string]interface{}{}, Collect(getMap()).GroupBy("b"))
}
