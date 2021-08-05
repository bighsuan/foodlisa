package prods

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
	// loginUser
	_user := obj.Ctx.MustGet("Id").(string)
	loginUser, _ := strconv.Atoi(_user)

	products := make([]models.Product, 0)
	result := obj.DB.Where(&models.Product{UserId: loginUser}).Find(&products)

	if result.Error == nil {
		obj.Ctx.JSON(http.StatusOK, products)
	} else {
		obj.Ctx.JSON(400, gin.H{
			"message": "DB query failed.",
		})
	}

}
