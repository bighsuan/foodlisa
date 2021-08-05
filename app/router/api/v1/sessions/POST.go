package sessions

import (
	"encoding/base64"
	"fmt"
	"foodlisa/models"
	"foodlisa/router/baseHandler"
	"foodlisa/tools/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
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

	var userInDb models.User
	result := obj.DB.Where(&models.User{Phone: post.Phone}).First(&userInDb)

	if result.Error != nil {
		obj.Ctx.JSON(400, gin.H{
			"message": "用戶尚未註冊",
		})
		return
	}

	// 密碼解密
	encryPassword, err := base64.StdEncoding.DecodeString(post.Password)
	if err != nil {
		fmt.Println(err.Error())
		obj.Ctx.JSON(400, gin.H{
			"message": "密碼加密錯誤（base64)",
		})
		return
	}

	decryPassword,err:=rsa.RsaDecrypt(encryPassword)
	if err!=nil{
		fmt.Println(err.Error())
		obj.Ctx.JSON(400, gin.H{
			"message": "密碼加密錯誤(rsa)",
		})
		return
	}

	//decryPassword := "123"
	err = bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(decryPassword))
	if err != nil {
		obj.Ctx.JSON(400, gin.H{
			"message": "密碼錯誤",
		})
		return
	} else {
		// 登入成功, 產token
		now := time.Now()

		// custom claims
		type Claims struct {
			UserId string      `json:"id"`
			jwt.StandardClaims
		}

		// jwt secret key
		//var jwtSecret = []byte(config.JWTKEY)
		var jwtSecret = []byte("secret")

		// set claims and sign
		claims := Claims{
			UserId: strconv.Itoa(int(userInDb.ID)),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: now.Add(24 * time.Hour).Unix(),
			},
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "伺服器錯誤",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id": userInDb.ID,
			"name": userInDb.FirstName,
			"token": token,
		})
	}
}
