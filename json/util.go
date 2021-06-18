package json

import (
	"bytes"
	"encoding/json"
)

// Pretty ...
func Pretty(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
