package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionCoverage1(t *testing.T) {
	assert.Equal(t, 1, Collect([]int{1}).Length())
	assert.Equal(t, 1, Collect([]int8{1}).Length())
	assert.Equal(t, 1, Collect([]int16{1}).Length())
	assert.Equal(t, 1, Collect([]int32{1}).Length())
	assert.Equal(t, 1, Collect([]int64{1}).Length())
	assert.Equal(t, 1, Collect([]uint{1}).Length())
	assert.Equal(t, 1, Collect([]uint8{1}).Length())
	assert.Equal(t, 1, Collect([]uint16{1}).Length())
	assert.Equal(t, 1, Collect([]uint32{1}).Length())
	assert.Equal(t, 1, Collect([]uint64{1}).Length())
	assert.Equal(t, 1, Collect([]float32{1}).Length())
	assert.Equal(t, 1, Collect([]float64{1}).Length())

	m1 := []map[string]interface{}{
		{
			"a1":  int(1),
			"a2":  int8(1),
			"a3":  int16(1),
			"a4":  int32(1),
			"a5":  int64(1),
			"a6":  uint(1),
			"a7":  uint8(1),
			"a8":  uint16(1),
			"a9":  uint32(1),
			"a10": uint64(1),
			"a11": float32(1),
			"a12": float64(1),
		},
	}
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a1").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a2").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a3").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a4").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a5").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a6").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a7").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a8").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a9").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a10").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a11").Value().([]float64))
	assert.Equal(t, []float64{1}, Collect(m1).Pluck("a12").Value().([]float64))

	m2 := []interface{}{
		map[string]interface{}{
			"a": "b",
		},
	}
	assert.Equal(t, 1, Collect(m2).Length())

	m3 := []interface{}{
		[]string{"a"},
	}
	assert.Equal(t, 0, Collect(m3).Length())
}

func TestCollectionCoverage2(t *testing.T) {
	assert.Equal(t, "", Collect(getNumber()).Join("-"))
	assert.Equal(t, float64(0), Collect(getString()).Min())
	assert.Equal(t, float64(0), Collect(getString()).Max())
	assert.Equal(t, false, Collect(getMap()).Contains("a"))
	assert.Equal(t, nil, Collect(getMap()).Unique())
	assert.Equal(t, nil, Collect(getMap()).DelValue("a"))
	assert.Equal(t, nil, Collect(getString()).Pluck("a"))
	assert.Equal(t, nil, Collect(getString()).DelKeyValue("a", "b"))
	assert.Equal(t, map[string]interface{}{}, Collect(getString()).GroupBy("a"))
}

func TestCollectionCoverage3(t *testing.T) {
	assert.Equal(t, nil, Collect(1).Value())
	assert.Equal(t, 0, Collect(1).Length())
	assert.Equal(t, "", Collect(1).Json())
	assert.Equal(t, nil, Collect(1).DelKey(1))
	assert.Equal(t, nil, Collect(1).Filter(func(key, value interface{}) bool {
		return value.(float64) != 4
	}))
	assert.Equal(t, 0, Collect([]interface{}{}).Length())
}
