package id

import "C"
import (
	"foodlisa/models"
	"foodlisa/router/baseHandler"
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
	var product models.Product
	result := obj.DB.Take(&product, obj.C.Param("id"))

	if result.Error == nil {
		obj.C.JSON(http.StatusOK, gin.H{
			"id": product.ID,
			"storeId": product.StoreID,
			"name": product.Name,
			"price": product.Price,
			"comment": product.Comment,
			"imageUrl": product.ImageUrl,
		})
	} else {
		obj.C.JSON(400, gin.H{
			"message": "Product not found.",
		})
	}



}
