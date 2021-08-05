package router

import (
	"foodlisa/router/api/v1/prods"
	"foodlisa/router/api/v1/prods/id"
	"foodlisa/router/api/v1/publickey"
	"foodlisa/router/api/v1/sessions"
	"foodlisa/router/api/v1/users"
	"foodlisa/router/baseHandler"
	"foodlisa/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"path/filepath"
	"strings"
)

func InitRouter(DB *gorm.DB) (router *gin.Engine) {
	router = gin.Default()
	// CORS middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())


	// ssl 驗證用
	const DOWNLOADS_PATH = "downloads/"

	router.GET("/.test/pki/B78E08FB0DA3F3DA482B9E4FF5FC3D9B.txt", func (ctx *gin.Context) {
		const DOWNLOADS_PATH = ".test/pki"

		fileName := "B78E08FB0DA3F3DA482B9E4FF5FC3D9B.txt"
		targetPath := filepath.Join(DOWNLOADS_PATH, fileName)
		//This ckeck is for example, I not sure is it can prevent all possible filename attacks - will be much better if real filename will not come from user side. I not even tryed this code
		if !strings.HasPrefix(filepath.Clean(targetPath), DOWNLOADS_PATH) {
			ctx.String(403, "Look like you attacking me")
			return
		}
		//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+fileName )
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(targetPath)
	})


	v1 := router.Group("/api/v1")

	// 登入, 註冊
	v1.GET("/publickey", getHandlerFunc(DB))
	v1.POST("/sessions", getHandlerFunc(DB))
	v1.POST("/users", getHandlerFunc(DB)) // 註冊

	v1.Use(middleware.JwtMiddleware)
	{
		// 登出
		//v1.DELETE("/sessions", getHandlerFunc(DB))

		// 商品
		v1.GET("/prods", getHandlerFunc(DB))
		v1.GET("/prods/:id", getHandlerFunc(DB))
		v1.PUT("/prods/:id", getHandlerFunc(DB))
		v1.POST("/prods", getHandlerFunc(DB))
		v1.DELETE("/prods/:id", getHandlerFunc(DB))

		//// member
		//v1.GET("/users/:id", getHandlerFunc(DB))
		//v1.PUT("/users/:id", getHandlerFunc(DB))
		//v1.PUT("/users/:id/password", getHandlerFunc(DB)) //修改密碼
		//
		//// stores
		//v1.GET("/users/:id/stores", getHandlerFunc(DB))
		//v1.GET("/stores/1", getHandlerFunc(DB)) // 查某商店的資訊
		//v1.PUT("/stores/1", getHandlerFunc(DB)) // 改某商店的資訊
		//v1.DELETE("/stores/1", getHandlerFunc(DB)) // 刪除某商店

	}

	// 測試
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to FoodLisa.",
		})
	})

	return
}

var typeRegistry = make(map[string]baseHandler.IHandler)


func getHandlerFunc(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var handler baseHandler.IHandler

		// 依 router 找對應 handler
		// TODO: 要改個寫法...
		switch ctx.FullPath() + "|" + ctx.Request.Method {
		case "/api/v1/prods|GET":        handler = &prods.GET{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}
		case "/api/v1/prods/:id|GET":    handler = &id.GET{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}
		case "/api/v1/prods/:id|PUT":    handler = &id.PUT{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}
		case "/api/v1/prods|POST":       handler = &prods.POST{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}
		case "/api/v1/prods/:id|DELETE": handler = &id.DELETE{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}

		case "/api/v1/publickey|GET": handler = &publickey.GET{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}
		case "/api/v1/sessions|POST": handler = &sessions.POST{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}
		//case "/api/v1/sessions|DELETE": handler = &sessions.DELETE{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}

		case "/api/v1/users|POST": handler = &users.POST{BaseHandler: baseHandler.BaseHandler{Ctx: ctx, DB: DB}}

		}

		handler.Run(handler)
	}
}
