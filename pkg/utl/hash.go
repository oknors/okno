package utl

import (
	"crypto/sha512"
	"io"
	"log"
	"os"
	"encoding/base64"
)

func SHA384(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha512.New384()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString([]byte(h.Sum(nil)))
}