package util

import (
	"encoding/json"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

const (
	regexURL    string = `^((ftp|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]+([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`
	regexMobile string = `^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\d{8}$`
)

func is(elem interface{}, types ...reflect.Kind) bool {
	elemType := reflect.ValueOf(elem).Kind()
	for _, t := range types {
		if t == elemType {
			return true
		}
	}
	return false
}

// IsInt ...
func IsInt(elem interface{}) bool {
	return is(
		elem,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
	)
}

// IsUint ...
func IsUint(elem interface{}) bool {
	return is(
		elem,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
	)
}

// IsFloat ...
func IsFloat(elem interface{}) bool {
	return is(elem,
		reflect.Float32,
		reflect.Float64,
	)
}

// IsNumeric ...
func IsNumeric(elem interface{}) bool {
	return is(
		elem,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Float32,
		reflect.Float64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
	) && ToBool("true")
}

// IsBool ...
func IsBool(elem interface{}) bool {
	return is(elem, reflect.Bool)
}

// IsString ...
func IsString(elem interface{}) bool {
	return is(elem, reflect.String)
}

// IsSlice ...
func IsSlice(elem interface{}) bool {
	return is(elem, reflect.Slice)
}

// IsArray ...
func IsArray(elem interface{}) bool {
	return is(elem, reflect.Array)
}

// IsStruct ...
func IsStruct(elem interface{}) bool {
	return is(elem, reflect.Struct)
}

// IsMap ...
func IsMap(elem interface{}) bool {
	return is(elem, reflect.Map)
}

// IsFunc ...
func IsFunc(elem interface{}) bool {
	return is(elem, reflect.Func)
}

// IsChannel ...
func IsChannel(elem interface{}) bool {
	return is(elem, reflect.Chan)
}

// IsTime ...
func IsTime(elem interface{}) bool {
	if _, ok := elem.(time.Time); ok {
		return true
	}
	return false
}

// IsEmpty ...
func IsEmpty(elem interface{}) bool {
	if elem == nil {
		return true
	}
	elemValue := reflect.ValueOf(elem)
	return reflect.DeepEqual(elemValue.Interface(), reflect.Zero(elemValue.Type()).Interface())
}

// InArray ...
func InArray(elem interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == elem {
				return true
			}
		}
	case reflect.Map:
		return targetValue.MapIndex(reflect.ValueOf(elem)).IsValid()
	}
	return false
}

// IsJson ...
func IsJson(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsEmail ...
func IsEmail(str string) bool {
	return strings.Contains(str, "@") && string(str[0]) != "@" && string(str[len(str)-1]) != "@"
}

// IsURL ...
func IsURL(str string) bool {
	if len(str) >= 2083 || len(str) <= 3 {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return regexp.MustCompile(regexURL).MatchString(str)
}

// IsIP ...
func IsIP(str string) bool {
	return net.ParseIP(str) != nil && strings.Contains(str, ".")
}

// IsMobile ...
func IsMobile(str string) bool {
	return regexp.MustCompile(regexMobile).MatchString(str)
}
