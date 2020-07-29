package util

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

const (
	REGEX_CAMEL string = `[\p{L}\p{N}]+`
)

func ToInt64(str string) int64 {
	resp, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return 0
	}
	return resp
}

func ToString(elem interface{}) string {
	return string(fmt.Sprintf("%v", elem))
}

func ToBool(str string) bool {
	resp, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return resp
}

func ToCamelCase(str string) string {
	regexCameling := regexp.MustCompile(REGEX_CAMEL)
	chunks := regexCameling.FindAll([]byte(str), -1)
	for k, v := range chunks {
		chunks[k] = bytes.Title(v)
	}
	return string(bytes.Join(chunks, nil))
}

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
