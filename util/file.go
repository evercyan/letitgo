package util

import (
	"io/ioutil"
	"os"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsFile(path string) bool {
	file, err := os.Stat(path)
	return err == nil && !file.IsDir()
}

func IsDir(path string) bool {
	file, err := os.Stat(path)
	return err == nil && file.IsDir()
}

func GetSize(path string) int64 {
	file, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return file.Size()
}

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

func WriteFile(path, str string) error {
	return ioutil.WriteFile(path, []byte(str), 0755)
}
