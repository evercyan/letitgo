package collection

import (
	"encoding/json"
)

type collectionNumber struct {
	collectionBase
	value []float64
}

func (c collectionNumber) Value() interface{} {
	return c.value
}

func (c collectionNumber) Length() int {
	return c.length
}

func (c collectionNumber) Json() string {
	str, _ := json.Marshal(c.value)
	return string(str)
}

func (c collectionNumber) Min() float64 {
	min := c.value[0]
	for i := 1; i < len(c.value); i++ {
		if c.value[i] < min {
			min = c.value[i]
		}
	}
	return min
}

func (c collectionNumber) Max() float64 {
	max := c.value[0]
	for i := 1; i < len(c.value); i++ {
		if c.value[i] > max {
			max = c.value[i]
		}
	}
	return max
}

func (c collectionNumber) Contains(value interface{}) bool {
	for i := 0; i < len(c.value); i++ {
		if c.value[i] == value.(float64) {
			return true
		}
	}
	return false
}

func (c collectionNumber) Unique() collection {
	list := []float64{}
	m := map[float64]bool{}
	for i := 0; i < len(c.value); i++ {
		if _, ok := m[c.value[i]]; !ok {
			list = append(list, c.value[i])
			m[c.value[i]] = true
		}
	}
	return Collect(list)
}

func (c collectionNumber) DelKey(key int) collection {
	if key < 0 || key >= len(c.value) {
		return c
	}
	var list []float64
	if key == len(c.value)-1 {
		list = c.value[:key]
	} else {
		list = append(c.value[:key], c.value[key+1:]...)
	}
	return Collect(list)
}

func (c collectionNumber) DelValue(value interface{}) collection {
	for i := 0; i < len(c.value); i++ {
		if c.value[i] == value.(float64) {
			return c.DelKey(i)
		}
	}
	return c
}

func (c collectionNumber) Filter(callback filterCallback) collection {
	list := []float64{}
	for i := 0; i < len(c.value); i++ {
		if callback(i, c.value[i]) {
			list = append(list, c.value[i])
		}
	}
	return Collect(list)
}
