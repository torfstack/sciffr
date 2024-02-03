package backend

import "github.com/gin-gonic/gin"

type Backend struct {
}

func New() *Backend {
	return &Backend{}
}

func (b *Backend) Register(e *gin.Engine) {
	e.GET("/encrypt", encrypt)
}

func encrypt(c *gin.Context) {
	c.JSON(200, "ENCRYPTED")
}
