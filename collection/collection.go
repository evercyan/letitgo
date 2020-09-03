/**
 * 基于对 https://github.com/chenhg5/collection 的学习和自己使用的简化
 *
 * 定义 Collection 接口, []string, []number, []map 均实现该接口
 * 公共 func, 继承自 CollectionBase
 *
 * number 类型统一存储为 float64
 *
 * []string, []int, []map[string]interface{}, []interface{}
 */

package collection

type Collection interface {
	Value() interface{}
	Length() int
	Json() string

	Join(string) string

	Min() float64
	Max() float64

	Contains(interface{}) bool
	Unique() Collection
	DelKey(int) Collection
	DelValue(interface{}) Collection

	Pluck(string) Collection
	DelKeyValue(string, interface{}) Collection

	Filter(FilterCallback) Collection
	GroupBy(string) map[string]interface{}
}

type FilterCallback func(key, value interface{}) bool

func Collect(elem interface{}) Collection {
	switch elem.(type) {
	case []string:
		var c CollectionString
		c.value = elem.([]string)
		c.length = len(elem.([]string))
		return c
	case []int:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]int)))
		for k, v := range elem.([]int) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []int8:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]int8)))
		for k, v := range elem.([]int8) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []int16:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]int16)))
		for k, v := range elem.([]int16) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []int32:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]int32)))
		for k, v := range elem.([]int32) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []int64:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]int64)))
		for k, v := range elem.([]int64) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]uint)))
		for k, v := range elem.([]uint) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint8:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]uint8)))
		for k, v := range elem.([]uint8) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint16:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]uint16)))
		for k, v := range elem.([]uint16) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint32:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]uint32)))
		for k, v := range elem.([]uint32) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint64:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]uint64)))
		for k, v := range elem.([]uint64) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []float32:
		var c CollectionNumber
		var value = make([]float64, len(elem.([]float32)))
		for k, v := range elem.([]float32) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []float64:
		var c CollectionNumber
		c.value = elem.([]float64)
		c.length = len(elem.([]float64))
		return c
	case []map[string]interface{}:
		var m CollectionMap
		m.value = elem.([]map[string]interface{})
		m.length = len(elem.([]map[string]interface{}))
		return m
	case []interface{}:
		length := len(elem.([]interface{}))
		if length == 0 {
			return CollectionBase{}
		}
		switch elem.([]interface{})[0].(type) {
		case string:
			var c CollectionString
			var value = make([]string, length)
			for k, v := range elem.([]interface{}) {
				value[k] = v.(string)
			}
			c.value = value
			c.length = length
			return c
		case int:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int))
			}
			c.value = value
			c.length = length
			return c
		case int8:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int8))
			}
			c.value = value
			c.length = length
			return c
		case int16:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int16))
			}
			c.value = value
			c.length = length
			return c
		case int32:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int32))
			}
			c.value = value
			c.length = length
			return c
		case int64:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int64))
			}
			c.value = value
			c.length = length
			return c
		case uint:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint))
			}
			c.value = value
			c.length = length
			return c
		case uint8:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint8))
			}
			c.value = value
			c.length = length
			return c
		case uint16:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint16))
			}
			c.value = value
			c.length = length
			return c
		case uint32:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint32))
			}
			c.value = value
			c.length = length
			return c
		case uint64:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint64))
			}
			c.value = value
			c.length = length
			return c
		case float32:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(float32))
			}
			c.value = value
			c.length = length
			return c
		case float64:
			var c CollectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(float64))
			}
			c.value = value
			c.length = length
			return c
		case map[string]interface{}:
			var c CollectionMap
			var value = make([]map[string]interface{}, length)
			for k, v := range elem.([]interface{}) {
				value[k] = v.(map[string]interface{})
			}
			c.value = value
			c.length = length
			return c
		default:
			return CollectionBase{}
		}
	default:
		return CollectionBase{}
	}
}
