package server

import (
	"sha256"
)

func accept(base64string string) string {
	return sha256.Sum256([]byte(base64string))
}
