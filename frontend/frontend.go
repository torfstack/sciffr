package frontend

import "github.com/gin-gonic/gin"

type Frontend struct {
}

func New() *Frontend {
	return &Frontend{}
}

func (f *Frontend) Register(engine *gin.Engine) {

}
