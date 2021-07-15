package util

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"math/rand"
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
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}
	return ""
}

// Guid ...
func Guid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

// Rand ...
func Rand(min, max int) int {
	if min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max+1-min)
}

// Range ...
func Range(min, max int) []int {
	var list []int
	if min > max {
		return list
	}
	for i := min; i <= max; i++ {
		list = append(list, i)
	}
	return list
}

// RandString ...
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
