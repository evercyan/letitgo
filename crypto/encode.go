package crypto

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

// Base64Encode ...
func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// Base64Decode ...
func Base64Decode(text string) string {
	resp, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}
	return string(resp)
}

// UrlEncode ...
func UrlEncode(text string) string {
	return url.QueryEscape(text)
}

// UrlDecode ...
func UrlDecode(text string) string {
	resp, err := url.QueryUnescape(text)
	if err != nil {
		return ""
	}
	return resp
}

// JsonEncode ...
func JsonEncode(elem interface{}) string {
	bytes, _ := json.Marshal(elem)
	return string(bytes)
}
