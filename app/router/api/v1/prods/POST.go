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
	obj.C.BindJSON(&product)

	product.StoreID = 1
	product.CategoryID = 1

	result := obj.DB.Create(&product)
	if result.Error == nil {
		obj.DB.Save(&product)
		obj.C.JSON(204, "")
	} else {
		obj.C.JSON(400, gin.H{
			"message": "Create failed.",
		})
	}
}
