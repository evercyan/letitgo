/**
 * 定义 collection 接口, []string, []number, []map 均实现该接口
 * 公共 func, 继承自 collectionBase
 *
 * number 类型统一存储为 float64
 *
 * []string, []int, []map[string]interface{}, []interface{}
 */

package collection

type collection interface {
	Value() interface{}
	Length() int
	Json() string

	Join(string) string

	Min() float64
	Max() float64

	Contains(interface{}) bool
	Unique() collection
	DelKey(int) collection
	DelValue(interface{}) collection

	Pluck(string) collection
	DelKeyValue(string, interface{}) collection

	Filter(filterCallback) collection
	GroupBy(string) map[string]interface{}
}

type filterCallback func(key, value interface{}) bool

// Collect ...
func Collect(elem interface{}) collection {
	switch elem.(type) {
	case []string:
		var c collectionString
		c.value = elem.([]string)
		c.length = len(elem.([]string))
		return c
	case []int:
		var c collectionNumber
		var value = make([]float64, len(elem.([]int)))
		for k, v := range elem.([]int) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []int8:
		var c collectionNumber
		var value = make([]float64, len(elem.([]int8)))
		for k, v := range elem.([]int8) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []int16:
		var c collectionNumber
		var value = make([]float64, len(elem.([]int16)))
		for k, v := range elem.([]int16) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []int32:
		var c collectionNumber
		var value = make([]float64, len(elem.([]int32)))
		for k, v := range elem.([]int32) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []int64:
		var c collectionNumber
		var value = make([]float64, len(elem.([]int64)))
		for k, v := range elem.([]int64) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint:
		var c collectionNumber
		var value = make([]float64, len(elem.([]uint)))
		for k, v := range elem.([]uint) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint8:
		var c collectionNumber
		var value = make([]float64, len(elem.([]uint8)))
		for k, v := range elem.([]uint8) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint16:
		var c collectionNumber
		var value = make([]float64, len(elem.([]uint16)))
		for k, v := range elem.([]uint16) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint32:
		var c collectionNumber
		var value = make([]float64, len(elem.([]uint32)))
		for k, v := range elem.([]uint32) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []uint64:
		var c collectionNumber
		var value = make([]float64, len(elem.([]uint64)))
		for k, v := range elem.([]uint64) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []float32:
		var c collectionNumber
		var value = make([]float64, len(elem.([]float32)))
		for k, v := range elem.([]float32) {
			value[k] = float64(v)
		}
		c.value = value
		c.length = len(value)
		return c
	case []float64:
		var c collectionNumber
		c.value = elem.([]float64)
		c.length = len(elem.([]float64))
		return c
	case []map[string]interface{}:
		var m collectionMap
		m.value = elem.([]map[string]interface{})
		m.length = len(elem.([]map[string]interface{}))
		return m
	case []interface{}:
		length := len(elem.([]interface{}))
		if length == 0 {
			return collectionBase{}
		}
		switch elem.([]interface{})[0].(type) {
		case string:
			var c collectionString
			var value = make([]string, length)
			for k, v := range elem.([]interface{}) {
				value[k] = v.(string)
			}
			c.value = value
			c.length = length
			return c
		case int:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int))
			}
			c.value = value
			c.length = length
			return c
		case int8:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int8))
			}
			c.value = value
			c.length = length
			return c
		case int16:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int16))
			}
			c.value = value
			c.length = length
			return c
		case int32:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int32))
			}
			c.value = value
			c.length = length
			return c
		case int64:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(int64))
			}
			c.value = value
			c.length = length
			return c
		case uint:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint))
			}
			c.value = value
			c.length = length
			return c
		case uint8:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint8))
			}
			c.value = value
			c.length = length
			return c
		case uint16:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint16))
			}
			c.value = value
			c.length = length
			return c
		case uint32:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint32))
			}
			c.value = value
			c.length = length
			return c
		case uint64:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(uint64))
			}
			c.value = value
			c.length = length
			return c
		case float32:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(float32))
			}
			c.value = value
			c.length = length
			return c
		case float64:
			var c collectionNumber
			var value = make([]float64, length)
			for k, v := range elem.([]interface{}) {
				value[k] = float64(v.(float64))
			}
			c.value = value
			c.length = length
			return c
		case map[string]interface{}:
			var c collectionMap
			var value = make([]map[string]interface{}, length)
			for k, v := range elem.([]interface{}) {
				value[k] = v.(map[string]interface{})
			}
			c.value = value
			c.length = length
			return c
		default:
			return collectionBase{}
		}
	default:
		return collectionBase{}
	}
}
