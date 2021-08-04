package prods

import (
	"foodlisa/models"
	"foodlisa/router/baseHandler"
	"github.com/gin-gonic/gin"
)

type POST struct {
	baseHandler.BaseHandler
}

var product models.Product

func (obj *POST) Validate() {
	// TODO: Validate your input
}

func (obj *POST) Process() {
	var product models.Product
	obj.Ctx.BindJSON(&product)

	// 給預設值
	if product.StoreID == 0 {
		product.StoreID = 1
	}

	product.CategoryID = 1

	result := obj.DB.Create(&product)
	if result.Error == nil {
		obj.DB.Save(&product)
		obj.Ctx.JSON(204, "")
	} else {
		obj.Ctx.JSON(400, gin.H{
			"message": "Create failed.",
		})
	}
}
