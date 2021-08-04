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

// custom claims
type Claims struct {
	Id string      `json:"id"`
	jwt.StandardClaims
}

var body struct{
	Phone string
	Password string
}

// jwt secret key
//var jwtSecret = []byte(config.JWTKEY)
var jwtSecret = "secret"

func (obj *POST) Validate() {
	// TODO: Validate your input
}

func (obj *POST) Process() {
	c := obj.Ctx

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	result := obj.DB.Where(&models.User{Phone: body.Phone}).First(&user)

	if result.Error != nil {
		obj.Ctx.JSON(400, gin.H{
			"message": "用戶尚未註冊",
		})
		return
	}

	// 密碼解密
	encryPassword, err := base64.StdEncoding.DecodeString(body.Password)
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
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(decryPassword))
	if err != nil {
		obj.Ctx.JSON(400, gin.H{
			"message": "密碼錯誤",
		})
		return
	} else {
		// 登入成功, 產token
		now := time.Now()

		// set claims and sign
		claims := Claims{
			Id: strconv.Itoa(int(user.ID)),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: now.Add(24 * time.Hour).Unix(),
			},
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, err := tokenClaims.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "伺服器錯誤",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
