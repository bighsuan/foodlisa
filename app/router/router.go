package router

import (
	"foodlisa/router/api/v1/prods"
	"foodlisa/router/api/v1/prods/id"
	"foodlisa/router/baseHandler"
	"foodlisa/router/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(DB *gorm.DB) (router *gin.Engine) {
	router = gin.Default()
	// CORS middleware
	router.Use(middleware.CORSMiddleware())
	//router.Use(middleware.Cors()) //开启中间件 允许使用跨域请求

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	
	v1 := router.Group("/api/v1")
	{
		// 商品
		v1.GET("/prods", getHandlerFunc(DB))
		v1.GET("/prods/:id", getHandlerFunc(DB))
		v1.PUT("/prods/:id", getHandlerFunc(DB))
		v1.POST("/prods", getHandlerFunc(DB))
		v1.DELETE("/prods/:id", getHandlerFunc(DB))
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
	return func(c *gin.Context) {
		var handler baseHandler.IHandler

		// 依 router 找對應 handler
		// TODO: 要改個寫法...
		switch c.FullPath() + "|" + c.Request.Method {
		case "/api/v1/prods|GET":        handler = &prods.GET{BaseHandler: baseHandler.BaseHandler{C: c, DB: DB}}
		case "/api/v1/prods/:id|GET":    handler = &id.GET{BaseHandler: baseHandler.BaseHandler{C: c, DB: DB}}
		case "/api/v1/prods/:id|PUT":    handler = &id.PUT{BaseHandler: baseHandler.BaseHandler{C: c, DB: DB}}
		case "/api/v1/prods|POST":       handler = &prods.POST{BaseHandler: baseHandler.BaseHandler{C: c, DB: DB}}
		case "/api/v1/prods/:id|DELETE": handler = &id.DELETE{BaseHandler: baseHandler.BaseHandler{C: c, DB: DB}}
		}

		handler.Run(handler)
	}
}