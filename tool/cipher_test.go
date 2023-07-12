package tool

import "testing"

func TestCipher(t *testing.T) {
	key := RandomByte(16)
	iv := RandomByte(16)
	text := "hello test"
	cipher := CipherTool(key)

	enc, err := cipher.Encrypt([]byte(text), iv)

	if err != nil {
		t.Errorf("error encrypting in cipher: %s", err)
	}

	dec, err := cipher.Decrypt(enc, iv)

	if err != nil {
		t.Errorf("error decrypting cipher: %s", err)
	}

	if string(dec) != text {
		t.Errorf("invalid decipher values, expected '%s' and received '%s'", text, dec)
	}
}
