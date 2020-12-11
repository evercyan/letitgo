package util

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

const (
	regexCamel string = `[\p{L}\p{N}]+`
)

// ToUint ...
func ToUint(elem interface{}) uint64 {
	if elem == nil {
		return 0
	}
	switch v := elem.(type) {
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case string:
		if v == "" {
			return 0
		}
		n, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0
		}
		return n
	default:
		return 0
	}
}

// ToString ...
func ToString(elem interface{}) string {
	return fmt.Sprintf("%v", elem)
}

// ToBool ...
func ToBool(elem interface{}) bool {
	if elem == nil {
		return false
	}
	if v, ok := elem.(bool); ok {
		return v
	}
	v, err := strconv.ParseBool(ToString(elem))
	return err == nil && v
}

// ToCamelCase ...
func ToCamelCase(str string) string {
	chunks := regexp.MustCompile(regexCamel).FindAll([]byte(str), -1)
	for k, v := range chunks {
		chunks[k] = bytes.Title(v)
	}
	return string(bytes.Join(chunks, nil))
}

// ToSnakeCase ...
func ToSnakeCase(str string) string {
	str = ToCamelCase(str)
	runes := []rune(str)
	length := len(runes)
	var resp []rune
	for i := 0; i < length; i++ {
		resp = append(resp, unicode.ToLower(runes[i]))
		if i+1 < length && (unicode.IsUpper(runes[i+1]) && unicode.IsLower(runes[i])) {
			resp = append(resp, '_')
		}
	}
	return string(resp)
}
