package collection

import (
	"encoding/json"
)

type collectionString struct {
	collectionBase
	value []string
}

func (c collectionString) Value() interface{} {
	return c.value
}

func (c collectionString) Length() int {
	return c.length
}

func (c collectionString) Json() string {
	str, _ := json.Marshal(c.value)
	return string(str)
}

func (c collectionString) Join(delimiter string) string {
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

func (c collectionString) Contains(value interface{}) bool {
	for i := 0; i < len(c.value); i++ {
		if c.value[i] == value.(string) {
			return true
		}
	}
	return false
}

func (c collectionString) Unique() collection {
	list := make([]string, 0)
	m := map[string]bool{}
	for i := 0; i < len(c.value); i++ {
		if _, ok := m[c.value[i]]; !ok {
			list = append(list, c.value[i])
			m[c.value[i]] = true
		}
	}
	return Collect(list)
}

func (c collectionString) DelKey(key int) collection {
	if key < 0 || key >= len(c.value) {
		return c
	}
	var list []string
	if key == len(c.value)-1 {
		list = c.value[:key]
	} else {
		list = append(c.value[:key], c.value[key+1:]...)
	}
	return Collect(list)
}

func (c collectionString) DelValue(value interface{}) collection {
	for i := 0; i < len(c.value); i++ {
		if c.value[i] == value.(string) {
			return c.DelKey(i)
		}
	}
	return c
}

func (c collectionString) Filter(callback filterCallback) collection {
	list := make([]string, 0)
	for i := 0; i < len(c.value); i++ {
		if callback(i, c.value[i]) {
			list = append(list, c.value[i])
		}
	}
	return Collect(list)
}
