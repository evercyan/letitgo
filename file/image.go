package file

import (
	"bufio"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// ImageType ...
func ImageType(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	_, t, err := image.Decode(bufio.NewReader(file))
	if err != nil {
		return ""
	}
	return t
}
