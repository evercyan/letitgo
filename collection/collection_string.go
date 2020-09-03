package collection

import (
	"encoding/json"
)

type CollectionString struct {
	CollectionBase
	value []string
}

func (c CollectionString) Value() interface{} {
	return c.value
}

func (c CollectionString) Length() int {
	return c.length
}

func (c CollectionString) Json() string {
	str, _ := json.Marshal(c.value)
	return string(str)
}

func (c CollectionString) Join(delimiter string) string {
	str := ""
	for i := 0; i < len(c.value); i++ {
		if i != len(c.value)-1 {
			str += c.value[i] + delimiter
		} else {
			str += c.value[i]
		}
	}
	return str
}

func (c CollectionString) Contains(value interface{}) bool {
	for i := 0; i < len(c.value); i++ {
		if c.value[i] == value.(string) {
			return true
		}
	}
	return false
}

func (c CollectionString) Unique() Collection {
	list := []string{}
	m := map[string]bool{}
	for i := 0; i < len(c.value); i++ {
		if _, ok := m[c.value[i]]; !ok {
			list = append(list, c.value[i])
			m[c.value[i]] = true
		}
	}
	return Collect(list)
}

func (c CollectionString) DelKey(key int) Collection {
	if key < 0 || key >= len(c.value) {
		return c
	}
	list := []string{}
	if key == len(c.value)-1 {
		list = c.value[:key]
	} else {
		list = append(c.value[:key], c.value[key+1:]...)
	}
	return Collect(list)
}

func (c CollectionString) DelValue(value interface{}) Collection {
	for i := 0; i < len(c.value); i++ {
		if c.value[i] == value.(string) {
			return c.DelKey(i)
		}
	}
	return c
}

func (c CollectionString) Filter(callback FilterCallback) Collection {
	list := []string{}
	for i := 0; i < len(c.value); i++ {
		if callback(i, c.value[i]) {
			list = append(list, c.value[i])
		}
	}
	return Collect(list)
}
