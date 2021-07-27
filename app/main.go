package main

import (
	"fmt"
	. "foodlisa/config"
	"foodlisa/orm"
	"foodlisa/router"
)



func main() {
	/* GORM */
	dsn := fmt.Sprint(USERNAME,":",PASSWORD,"@tcp(",HOST,":",PORT,")/",DATABASE,"?charset=utf8mb4&parseTime=True&loc=Local")
	var DB = orm.InitDatabase(dsn)

	/* GIN framework*/
	var router = router.InitRouter(DB)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}