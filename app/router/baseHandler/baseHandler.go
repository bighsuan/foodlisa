package baseHandler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Abstract Interface
type IHandler interface {
	Validate()
	Process()
	Run(IHandler)
}

type BaseHandler struct {
	DB       *gorm.DB
	C        *gin.Context
	Validate func()
	Process  func()
}

func (*BaseHandler) Run(i IHandler) {
	i.Validate()
	i.Process()
}