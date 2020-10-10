package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/evercyan/letitgo/util"
)

// Size ...
func Size(filePath string) int64 {
	file, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return file.Size()
}

// SizeText ...
func SizeText(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.2fKB", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2fMB", float64(size)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2fGB", float64(size)/(1024*1024*1024))
	}
}

// Read ...
func Read(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()
	resp, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(resp)
}

// Write ...
func Write(filePath, str string) error {
	return ioutil.WriteFile(filePath, []byte(str), 0755)
}

// Ext ...
func Ext(file string) string {
	return strings.ToLower(path.Ext(file))
}

// LineCount ...
func LineCount(filePath string) int {
	count := 0
	file, err := os.Open(filePath)
	if err != nil {
		return count
	}
	defer file.Close()

	fr := bufio.NewReader(file)
	buf := make([]byte, 32*1024)
	separator := []byte("\n")
	for {
		b, err := fr.Read(buf)
		count += bytes.Count(buf[:b], separator)
		// io.EOF 或异常都直接返回
		if err != nil {
			return count
		}
	}
}

// LineContent ...
func LineContent(filePath string, numbers ...int) map[int]string {
	result := map[int]string{}
	file, err := os.Open(filePath)
	if err != nil {
		return result
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	number := 0
	for fileScanner.Scan() {
		number++
		// 如果 numbers 为空, 则取所有行
		if len(numbers) == 0 || util.InArray(number, numbers) {
			result[number] = fileScanner.Text()
		}
	}
	return result
}
