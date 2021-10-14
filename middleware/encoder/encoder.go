package encoder

import (
	"crypto/md5"
	"fmt"
	"io"
)

func EncodeHash (str string) (string, error) {
	h := md5.New()
	if _, err := io.WriteString(h, str); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%x", h.Sum(nil)), nil
	}
}
