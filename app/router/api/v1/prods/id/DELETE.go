package id

import (
	"foodlisa/models"
	"foodlisa/router/baseHandler"
	"github.com/gin-gonic/gin"
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
	result := obj.DB.Take(&product, obj.C.Param("id"))

	if result.Error != nil {
		obj.C.JSON(400, gin.H{
			"message": "Product is not exist.",
		})
		return
	}

	// 刪除
	result = obj.DB.Delete(&models.Product{}, obj.C.Param("id"))

	if result.Error == nil {
		obj.C.JSON(204, "")
	} else {
		obj.C.JSON(400, gin.H{
			"message": "Delete failed.",
		})
	}
}
