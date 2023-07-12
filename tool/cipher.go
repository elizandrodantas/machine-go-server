package tool

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

type Cipher struct {
	Key []byte
}

func RandomByte(len int) []byte {
	randomBytes := make([]byte, len)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return []byte{0}
	}

	return randomBytes
}

func CipherTool(key []byte) *Cipher {
	if key == nil {
		key = RandomByte(16)
	}

	return &Cipher{key}
}

func (y *Cipher) Encrypt(plaintext, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(y.Key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

func (y *Cipher) Decrypt(ciphertext, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(y.Key)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)

	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
