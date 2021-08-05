package prods

import (
	"foodlisa/models"
	"foodlisa/router/baseHandler"
	"github.com/gin-gonic/gin"
	"strconv"
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

	// loginUser
	_user := obj.Ctx.MustGet("Id").(string)
	loginUser, _ := strconv.Atoi(_user)
	product.UserId = loginUser

	result := obj.DB.Create(&product)
	if result.Error == nil {
		obj.DB.Save(&product)
		obj.Ctx.Writer.WriteHeader(204)
	} else {
		obj.Ctx.JSON(400, gin.H{
			"message": "Create failed.",
		})
	}
}
