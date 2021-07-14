package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getNumber() []int {
	return []int{1, 2, 3, 4}
}

func TestCollectionNumberBase(t *testing.T) {
	assert.Equal(t, 4, Collect(getNumber()).Length())

	expected := []float64{1, 2, 3, 4}
	assert.Equal(t, expected, Collect(getNumber()).Value().([]float64))

	assert.Equal(t, "[1,2,3,4]", Collect(getNumber()).Json())
}

func TestCollectionNumberMinMax(t *testing.T) {
	assert.Equal(t, float64(1), Collect(getNumber()).Min())
	assert.Equal(t, float64(4), Collect(getNumber()).Max())
}

func TestCollectionNumberContains(t *testing.T) {
	assert.True(t, Collect(getNumber()).Contains(float64(1)))
	assert.False(t, Collect(getNumber()).Contains(float64(5)))
}

func TestCollectionNumberUnique(t *testing.T) {
	list := []int{1, 1, 2, 3, 4}
	expected := []float64{1, 2, 3, 4}
	assert.Equal(t, expected, Collect(list).Unique().Value().([]float64))
}

func TestCollectionNumberDel(t *testing.T) {
	expected := []float64{1, 2, 4}
	assert.Equal(t, expected, Collect(getNumber()).DelKey(2).Value().([]float64))
	assert.Equal(t, expected, Collect(getNumber()).DelValue(float64(3)).Value().([]float64))
}

func TestCollectionNumberFilter(t *testing.T) {
	assert.Equal(t, []float64{1, 2, 3}, Collect(getNumber()).Filter(func(key, value interface{}) bool {
		return value.(float64) != 4
	}).Value().([]float64))

	assert.Equal(t, []float64{1, 2}, Collect(getNumber()).Filter(func(key, value interface{}) bool {
		return value.(float64) <= 2
	}).Value().([]float64))
}

func TestCollectionNumberCoverage(t *testing.T) {
	assert.Equal(t, float64(1), Collect([]int{5, 1, 10, 2}).Min())
	assert.Equal(t, 4, Collect(getNumber()).DelKey(4).Length())
	assert.Equal(t, 3, Collect(getNumber()).DelKey(3).Length())
	assert.Equal(t, 4, Collect(getNumber()).DelValue(float64(100)).Length())
}
