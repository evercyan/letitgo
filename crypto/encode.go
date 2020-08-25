package crypto

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

func Base64Encode(text string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(text)))
}

func Base64Decode(text string) string {
	resp, err := base64.StdEncoding.DecodeString(text)
	if nil != err {
		return ""
	}
	return string(resp)
}

func UrlEncode(text string) string {
	return url.QueryEscape(text)
}

func UrlDecode(text string) string {
	resp, err := url.QueryUnescape(text)
	if nil != err {
		return ""
	}
	return string(resp)
}

func JsonEncode(elem interface{}) string {
	bytes, _ := json.Marshal(elem)
	return string(bytes)
}
