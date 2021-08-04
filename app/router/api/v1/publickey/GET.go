package publickey

import (
	"foodlisa/router/baseHandler"
	myrsa "foodlisa/tools/rsa"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GET struct {
	baseHandler.BaseHandler
}

func (obj *GET) Validate() {
	// TODO: Validate your input
}

func (obj *GET) Process() {
	if err := myrsa.GenRsaKey(1024); err != nil {
		obj.Ctx.JSON(500,gin.H{
			"message": "server error",
		})
		return
	}
	obj.Ctx.JSON(http.StatusOK, myrsa.PublicKeyStr)
}
