package main

// https://qiita.com/takehanKosuke/items/bbeb7581330910e72bb2

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// docsのディレクトリを指定
	// ←追記

	ginSwagger "github.com/swaggo/gin-swagger"   // ←追記
	"github.com/swaggo/gin-swagger/swaggerFiles" // ←追記
)

// @title go-swagger_test
// @version 1.0
// @license.name Naoki Maruyama
// @description このswaggerはswaggerの見本apiです

func main() {
	fmt.Println("Hello gin")

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
