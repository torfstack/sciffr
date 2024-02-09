package cryptservice

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

const (
	NonceSize = 12
)

type Encrypted struct {
	Ciphertext []byte
	Key        []byte
}

type CryptService interface {
	Encrypt([]byte) Encrypted
	Decrypt(encrypted Encrypted) []byte
}

type AesGcmCryptService struct {
	masterCipher cipher.AEAD
}

func New() *AesGcmCryptService {
	master := generateKey()
	return &AesGcmCryptService{
		masterCipher: initCipher(master),
	}
}

func (c *AesGcmCryptService) Encrypt(bytes []byte) Encrypted {
	nonce := c.nonce()
	key := generateKey()
	sealedKey := c.masterCipher.Seal(nonce, nonce, key, nil)
	cip := initCipher(key)
	ciphertext := cip.Seal(nonce, nonce, bytes, nil)
	return Encrypted{
		Ciphertext: ciphertext,
		Key:        sealedKey,
	}
}

func (c *AesGcmCryptService) Decrypt(encrypted Encrypted) []byte {
	bytes, key := encrypted.Ciphertext, encrypted.Key
	keyNonce, keyToDecrypt := key[:NonceSize], key[NonceSize:]
	bytesNonce, bytesToDecrypt := bytes[:NonceSize], bytes[NonceSize:]
	keyDecrypted, err := c.masterCipher.Open(nil, keyNonce, keyToDecrypt, nil)
	if err != nil {
		// handle error dunno
	}
	cip := initCipher(keyDecrypted)
	plaintext, err := cip.Open(nil, bytesNonce, bytesToDecrypt, nil)
	if err != nil {
		// handle error dunno
	}
	return plaintext
}

func (c *AesGcmCryptService) nonce() []byte {
	n := make([]byte, NonceSize)
	read, err := rand.Read(n)
	if err != nil || read != NonceSize {
		// handle error dunno
	}
	return n
}

func initCipher(key []byte) cipher.AEAD {
	block, err := aes.NewCipher(key)
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
