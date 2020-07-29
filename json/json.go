package json

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type js struct {
	value interface{}
}

func Json(text string) *js {
	j := new(js)
	var v interface{}
	if err := json.Unmarshal([]byte(text), &v); err != nil {
		return j
	}
	j.value = v
	return j
}

func (j *js) Key(key string) *js {
	m, ok := (j.value).(map[string]interface{})
	if !ok {
		j.value = nil
		return j
	}
	v, ok := m[key]
	if !ok {
		j.value = nil
		return j
	}
	j.value = v
	return j
}

func (j *js) Index(index int) *js {
	m, ok := (j.value).([]interface{})
	if !ok {
		j.value = nil
		return j
	}
	if index > len(m)-1 {
		j.value = nil
		return j
	}
	j.value = m[index]
	return j
}

func (j *js) Value() interface{} {
	return j.value
}

func (j *js) ToString() string {
	if j.value == nil {
		return ""
	}
	return string(fmt.Sprintf("%v", j.value))
}

func (j *js) ToInt64() int64 {
	resp, err := strconv.ParseInt(j.ToString(), 0, 64)
	if err != nil {
		return 0
	}
	return resp
}

func (j *js) ToJson() string {
	if j.value == nil {
		return ""
	}
	bytes, err := json.Marshal(j.value)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (j *js) ToArray() interface{} {
	switch (j.value).(type) {
	case []interface{}:
		return (j.value).([]interface{})
	case map[string]interface{}:
		return (j.value).(map[string]interface{})
	default:
		return nil
	}
}
