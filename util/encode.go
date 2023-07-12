package util

import (
	"encoding/base64"
	"encoding/hex"
)

func Encode64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode64(s string) []byte {
	output, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		return []byte{}
	}

	return output
}

func ByteToString(b []byte) string {
	return string(b)
}

func EncodeHex(h []byte) string {
	return hex.EncodeToString(h)
}
