package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// docsのディレクトリを指定
	// ←追記

	ginSwagger "github.com/swaggo/gin-swagger"   // ←追記
	"github.com/swaggo/gin-swagger/swaggerFiles" // ←追記
)

func main() {
	fmt.Println("Hello gin")

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
