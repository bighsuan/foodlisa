package id

import "C"
import (
	"foodlisa/models"
	"foodlisa/router/baseHandler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GET struct {
	baseHandler.BaseHandler
}

func (obj *GET) Validate() {
	// TODO: Validate your input
}

func (obj *GET) Process() {
	// 用id查商品
	var product models.Product
	result := obj.DB.Take(&product, obj.Ctx.Param("id"))

	if result.Error != nil {
		obj.Ctx.JSON(400, gin.H{
			"message": "Product not found.",
		})
		return
	}

	// 檢查是否是該用戶的商品
	_user := obj.Ctx.MustGet("Id").(string)
	loginUser, _ := strconv.Atoi(_user)
	if loginUser != product.UserId {
		obj.Ctx.Writer.WriteHeader(403)
		return
	}

	obj.Ctx.JSON(http.StatusOK, product)
}
