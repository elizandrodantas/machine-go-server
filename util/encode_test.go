package util

import (
	"testing"
)

func TestEncode64(t *testing.T) {
	const result = "dGhpcyBpcyBhbiB0ZXN0IGVuY29kZSBiYXNlNjQ="

	encode64 := Encode64([]byte("this is an test encode base64"))

	if encode64 != result {
		t.Errorf("base64 not encoded correctly, expected `%s` and got `%s`", result, encode64)
	}
}

func TestDecode64(t *testing.T) {
	const result = "this is an test decode base64"

	decode64 := Decode64("dGhpcyBpcyBhbiB0ZXN0IGRlY29kZSBiYXNlNjQ=")

	if decode64 == nil {
		t.Error("decode64 returned null value")
	}

	if string(decode64) != result {
		t.Errorf("base64 was not decoded correctly, expected `%s` and received `%s`", result, string(decode64))
	}
}

func TestByteToString(t *testing.T) {
	const result = "hello test"

	str := ByteToString([]byte(result))

	if str != result {
		t.Errorf("byte conversion expects a '%s' and received '%s'", result, str)
	}
}

func TestEncodeHex(t *testing.T) {
	const result = "68656c6c6f207465737420686578"

	hex := EncodeHex([]byte("hello test hex"))

	if hex != result {
		t.Errorf("hex conversion should be '%s' and get '%s'", result, hex)
	}
}
