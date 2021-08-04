package prods

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
	products := make([]models.Product, 0)
	result := obj.DB.Find(&products)

	if result.Error == nil {
		obj.Ctx.JSON(http.StatusOK, products)
	} else {
		obj.Ctx.JSON(400, gin.H{
			"message": "DB query failed.",
		})
	}

}
