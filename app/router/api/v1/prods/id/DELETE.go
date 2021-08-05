package id

import (
	"foodlisa/models"
	"foodlisa/router/baseHandler"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DELETE struct {
	baseHandler.BaseHandler
}

func (obj *DELETE) Validate() {
	// TODO: Validate your input
}

func (obj *DELETE) Process() {
	// 檢查商品是否存在
	var product models.Product
	result := obj.DB.Take(&product, obj.Ctx.Param("id"))

	if result.Error != nil {
		obj.Ctx.JSON(400, gin.H{
			"message": "Product is not exist.",
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

	// 刪除
	result = obj.DB.Delete(&models.Product{}, obj.Ctx.Param("id"))

	if result.Error == nil {
		obj.Ctx.Writer.WriteHeader(204)
	} else {
		obj.Ctx.JSON(400, gin.H{
			"message": "Delete failed.",
		})
	}
}
