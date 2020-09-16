package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	mrand "math/rand"
	"net"
	"time"
)

// Md5 ...
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GetClientIp ...
func GetClientIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return ""
}

// Guid ...
func Guid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(string(base64.URLEncoding.EncodeToString(b)))
}

// Rand ...
func Rand(min, max int) int {
	if min > max {
		return 0
	}
	mrand.Seed(time.Now().UnixNano())
	return min + mrand.Intn(max+1-min)
}

// Range ...
func Range(min, max int) []int {
	list := []int{}
	if min > max {
		return list
	}
	for i := min; i <= max; i++ {
		list = append(list, i)
	}
	return list
}
