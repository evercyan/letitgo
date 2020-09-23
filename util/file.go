package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// IsExist ...
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsFile ...
func IsFile(path string) bool {
	file, err := os.Stat(path)
	return err == nil && !file.IsDir()
}

// IsDir ...
func IsDir(path string) bool {
	file, err := os.Stat(path)
	return err == nil && file.IsDir()
}

// GetSize ...
func GetSize(path string) int64 {
	file, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return file.Size()
}

// ReadFile ...
func ReadFile(path string) string {
	file, err := os.Open(path)
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

// WriteFile ...
func WriteFile(path, str string) error {
	return ioutil.WriteFile(path, []byte(str), 0755)
}

// GetSizeText ...
func GetSizeText(size int64) string {
	if size < 1024 {
		return ToString(size) + "B"
	}
	if size < 1024*1024 {
		return fmt.Sprintf("%.2fKB", float64(size)/1024)
	}
	if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2fMB", float64(size)/(1024*1024))
	}
	return fmt.Sprintf("%.2fGB", float64(size)/(1024*1024*1024))
}

// GetLineCount ...
func GetLineCount(path string) int {
	count := 0
	file, err := os.Open(path)
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

// GetLineContent ...
func GetLineContent(path string, numbers ...int) map[int]string {
	result := map[int]string{}
	file, err := os.Open(path)
	if err != nil {
		return result
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	number := 0
	for fileScanner.Scan() {
		number++
		// 如果 numbers 为空, 则取所有行
		if len(numbers) == 0 || InArray(number, numbers) {
			result[number] = fileScanner.Text()
		}
	}
	return result
}
