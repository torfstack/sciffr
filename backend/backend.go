package backend

import "github.com/gin-gonic/gin"

type Backend struct {
}

func New() *Backend {
	return &Backend{}
}

func (b *Backend) Register(engine *gin.Engine) {

}
