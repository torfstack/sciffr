package backend

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"sciffr/backend/cryptservice"
	"sciffr/backend/db"

	"github.com/gin-gonic/gin"
)

type Backend struct {
	crypt cryptservice.CryptService
}

func New() *Backend {
	return &Backend{
		crypt: cryptservice.New(),
	}
}

func (b *Backend) Register(e *gin.Engine) {
	envelope := e.Group("v1/envelope")
	envelope.POST("/encrypt", b.encrypt)
	envelope.POST("/decrypt", b.decrypt)

	err := db.NewDatabasePostgres().Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.NewDatabasePostgres().AddKey("test")
	if err != nil {
		return
	}
}

func (b *Backend) encrypt(c *gin.Context) {
	var i Plaintext
	err := c.BindJSON(&i)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("invalid request"))
		return
	}
	toEncrypt := base64Decode(i.Plaintext)
	encrypted := b.crypt.Encrypt(toEncrypt)
	c.JSON(200, Ciphertext{
		Ciphertext: base64Encode(encrypted.Ciphertext),
		Key:        base64Encode(encrypted.Key),
	})
}

func (b *Backend) decrypt(c *gin.Context) {
	var i Ciphertext
	err := c.BindJSON(&i)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("invalid request"))
		return
	}
	toDecrypt := base64Decode(i.Ciphertext)
	key := base64Decode(i.Key)
	decrypted := b.crypt.Decrypt(cryptservice.Encrypted{
		Ciphertext: toDecrypt,
		Key:        key,
	})
	c.JSON(200, Plaintext{
		Plaintext: base64Encode(decrypted),
	})
}

func base64Decode(b string) []byte {
	dst := make([]byte, encoding().DecodedLen(len(b)))
	_, err := encoding().Decode(dst, []byte(b))
	if err != nil {
		// TODO: handle error
	}
	return dst
}

func base64Encode(b []byte) string {
	dst := make([]byte, encoding().EncodedLen(len(b)))
	encoding().Encode(dst, b)
	return string(dst)
}

func encoding() *base64.Encoding {
	return base64.RawURLEncoding
}

type Plaintext struct {
	Plaintext string `json:"plaintext"`
}

type Ciphertext struct {
	Ciphertext string `json:"ciphertext"`
	Key        string `json:"key"`
}
