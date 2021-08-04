package users

import (
	"encoding/base64"
	"fmt"
	"foodlisa/models"
	"foodlisa/router/baseHandler"
	"foodlisa/tools/rsa"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type POST struct {
	baseHandler.BaseHandler
}


func (obj *POST) Validate() {
	// TODO: Validate your input
}

func (obj *POST) Process() {
	c := obj.Ctx

	var post models.User

	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	// 檢查是否已註冊過
	var user models.User
	result := obj.DB.Where(&models.User{Phone: post.Phone}).First(&user)

	if result.Error == nil {
		obj.Ctx.JSON(400, gin.H{
			"message": "此號碼已被註冊",
		})
		return
	}

	post.Username = post.Phone

	//密碼解密
	encryPassword, err := base64.StdEncoding.DecodeString(post.Password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	decryPassword,err:=rsa.RsaDecrypt(encryPassword)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}

	//decryPassword := "123"

	// hash and salt
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(decryPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePW := string(hashPassword)  // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	post.Password = encodePW

	result = obj.DB.Create(&post)
	if result.Error == nil {
		obj.DB.Save(&user)
		obj.Ctx.JSON(204, "")
	} else {
		obj.Ctx.JSON(400, gin.H{
			"message": "註冊失敗",
		})
	}
}
