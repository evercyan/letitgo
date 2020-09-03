package collection

import (
	"encoding/json"
	"reflect"
)

type CollectionMap struct {
	CollectionBase
	value []map[string]interface{}
}

func (c CollectionMap) Value() interface{} {
	return c.value
}

func (c CollectionMap) Length() int {
	return c.length
}

func (c CollectionMap) Json() string {
	str, _ := json.Marshal(c.value)
	return string(str)
}

func (c CollectionMap) DelKey(key int) Collection {
	if key < 0 || key >= len(c.value) {
		return c
	}
	list := []map[string]interface{}{}
	if key == len(c.value)-1 {
		list = c.value[:key]
	} else {
		list = append(c.value[:key], c.value[key+1:]...)
	}
	return Collect(list)
}

func (c CollectionMap) DelKeyValue(key string, value interface{}) Collection {
	for i := 0; i < len(c.value); i++ {
		if c.value[i][key] == value {
			return c.DelKey(i)
		}
	}
	return c
}

func (c CollectionMap) Pluck(key string) Collection {
	list := make([]interface{}, 0)
	for i := 0; i < len(c.value); i++ {
		list = append(list, c.value[i][key])
	}
	return Collect(list)
}

func (c CollectionMap) Filter(callback FilterCallback) Collection {
	list := []map[string]interface{}{}
	for i := 0; i < len(c.value); i++ {
		if callback(i, c.value[i]) {
			list = append(list, c.value[i])
		}
	}
	return Collect(list)
}

func (c CollectionMap) GroupBy(key string) map[string]interface{} {
	m := map[string]interface{}{}
	if c.length == 0 {
		return m
	}
	value, ok := c.value[0][key]
	if !ok {
		return m
	}
	if reflect.ValueOf(value).Kind() != reflect.String {
		return m
	}
	for i := 0; i < len(c.value); i++ {
		m[c.value[i][key].(string)] = c.value[i]
	}
	return m
}
