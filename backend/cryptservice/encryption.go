package cryptservice

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

type CryptService interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

type cryptService struct {
	key []byte
}

func New() CryptService {
	return &cryptService{
		key: generateKey(),
	}
}

func (c *cryptService) Encrypt(bytes []byte) []byte {
	nonce := c.nonce()
	ciphertext := c.cipher().Seal(nonce, nonce, bytes, nil)
	return ciphertext
}

func (c *cryptService) Decrypt(bytes []byte) []byte {
	nonceSize := c.cipher().NonceSize()
	nonce, toDecrypt := bytes[:nonceSize], bytes[nonceSize:]
	plaintext, err := c.cipher().Open(nil, nonce, toDecrypt, nil)
	if err != nil {
		// handle error dunno
	}
	return plaintext
}

func (c *cryptService) nonce() []byte {
	n := make([]byte, c.cipher().NonceSize())
	read, err := rand.Read(n)
	if err != nil || read != c.cipher().NonceSize() {
		// handle error dunno
	}
	return n
}

func (c *cryptService) cipher() cipher.AEAD {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		// handle error dunno
	}
	cip, err := cipher.NewGCM(block)
	if err != nil {
		// handle error dunno
	}
	return cip
}

func generateKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		// handle error dunno
	}
	return key
}
