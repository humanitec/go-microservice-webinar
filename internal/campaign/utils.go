package campaign

import (
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)

func GetStringFromBuffer(body io.ReadCloser) (string, error) {
	b := new(strings.Builder)
	_, err := io.Copy(b, body)
	if err != nil {
		return "", fmt.Errorf("error handling the response body: %w", err)
	}
	return b.String(), nil
}

// Hash ...
func Hash(o interface{}) string {
	var h = sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
