package file

import (
	"bufio"
	"image"
	_ "image/gif"  // ...
	_ "image/jpeg" // ...
	_ "image/png"  // ...
	"os"
)

// IsExist ...
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

// IsFile ...
func IsFile(filePath string) bool {
	file, err := os.Stat(filePath)
	return err == nil && !file.IsDir()
}

// IsDir ...
func IsDir(filePath string) bool {
	file, err := os.Stat(filePath)
	return err == nil && file.IsDir()
}

// IsImage ...
func IsImage(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	_, _, err = image.Decode(bufio.NewReader(file))
	return err == nil
}
