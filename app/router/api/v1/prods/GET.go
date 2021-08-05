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
	products := make([]models.Product, 0)
	//result := obj.DB.Find(&products)

	_store := obj.Ctx.Query("store")
	storeId, _ := strconv.Atoi(_store)

	// 檢查登入id和查詢商品id
	_user := obj.Ctx.MustGet("Id").(string)
	loginUser, _ := strconv.Atoi(_user)

	if loginUser != storeId {
		obj.Ctx.Writer.WriteHeader(403)
		return
	}
	result := obj.DB.Where(&models.Product{UserId: storeId}).Find(&products)

	if result.Error == nil {
		obj.Ctx.JSON(http.StatusOK, products)
	} else {
		obj.Ctx.JSON(400, gin.H{
			"message": "DB query failed.",
		})
	}

}
