package collection

import (
	"encoding/json"
)

type CollectionNumber struct {
	CollectionBase
	value []float64
}

func (c CollectionNumber) Value() interface{} {
	return c.value
}

func (c CollectionNumber) Length() int {
	return c.length
}

func (c CollectionNumber) Json() string {
	str, _ := json.Marshal(c.value)
	return string(str)
}

func (c CollectionNumber) Min() float64 {
	min := c.value[0]
	for i := 1; i < len(c.value); i++ {
		if c.value[i] < min {
			min = c.value[i]
		}
	}
	return min
}

func (c CollectionNumber) Max() float64 {
	max := c.value[0]
	for i := 1; i < len(c.value); i++ {
		if c.value[i] > max {
			max = c.value[i]
		}
	}
	return max
}

func (c CollectionNumber) Contains(value interface{}) bool {
	for i := 0; i < len(c.value); i++ {
		if c.value[i] == value.(float64) {
			return true
		}
	}
	return false
}

func (c CollectionNumber) Unique() Collection {
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

func (c CollectionNumber) DelKey(key int) Collection {
	if key < 0 || key >= len(c.value) {
		return c
	}
	list := []float64{}
	if key == len(c.value)-1 {
		list = c.value[:key]
	} else {
		list = append(c.value[:key], c.value[key+1:]...)
	}
	return Collect(list)
}

func (c CollectionNumber) DelValue(value interface{}) Collection {
	for i := 0; i < len(c.value); i++ {
		if c.value[i] == value.(float64) {
			return c.DelKey(i)
		}
	}
	return c
}

func (c CollectionNumber) Filter(callback FilterCallback) Collection {
	list := []float64{}
	for i := 0; i < len(c.value); i++ {
		if callback(i, c.value[i]) {
			list = append(list, c.value[i])
		}
	}
	return Collect(list)
}
