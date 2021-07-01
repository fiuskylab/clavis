package clavis

import (
	"crypto/sha1"
	"fmt"
	"io"
)

// Encrypt key
func Sha1Encrypt(key string) string {
	h := sha1.New()

	io.WriteString(h, key)

	return fmt.Sprintf("%x", h.Sum(nil))
}
