package id

import (
	"foodlisa/models"
	"foodlisa/router/baseHandler"
	"github.com/gin-gonic/gin"
)

type PUT struct {
	baseHandler.BaseHandler
}

func (obj *PUT) Validate() {
	// TODO: Validate your input
}

func (obj *PUT) Process() {
	// 檢查商品是否存在
	var product models.Product
	result := obj.DB.Take(&product, obj.Ctx.Param("id"))

	if result.Error != nil {
		obj.Ctx.JSON(400, gin.H{
			"message": "Product is not exist.",
		})
		return
	}

	var json models.Product
	obj.Ctx.BindJSON(&json)

	// 不使用 DB.Save() 因為有些欄位沒有輸入值
	// obj.DB.Model(&product).Update("CategoryID", json.CategoryID)
	obj.DB.Model(&product).Update("Name", json.Name)
	obj.DB.Model(&product).Update("Price", json.Price)
	obj.DB.Model(&product).Update("Comment", json.Comment)
	obj.DB.Model(&product).Update("ImageUrl", json.ImageUrl)

	obj.Ctx.JSON(204, "")
}
