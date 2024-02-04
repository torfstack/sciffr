package backend

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"sciffr/backend/cryptservice"
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
	e.POST("/encrypt", b.encrypt)
	e.POST("/decrypt", b.decrypt)
}

func (b *Backend) encrypt(c *gin.Context) {
	var i CryptInput
	err := c.BindJSON(&i)
	if err != nil {
		// TODO: handle error
	}
	output := base64Encode(b.crypt.Encrypt(base64Decode(i.Value)))
	c.JSON(200, CryptOutput{
		Value: output,
	})
}

func (b *Backend) decrypt(c *gin.Context) {
	var i CryptInput
	err := c.BindJSON(&i)
	if err != nil {
		// TODO: handle error
	}
	output := base64Encode(b.crypt.Decrypt(base64Decode(i.Value)))
	c.JSON(200, CryptOutput{
		Value: output,
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

type CryptInput struct {
	Value string `json:"value"`
}

type CryptOutput struct {
	Value string `json:"value"`
}
