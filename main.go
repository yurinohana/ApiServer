// ルーティング処理

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yurinohana/ApiServer/controllers"
)

func main() {

	r := gin.Default()
	r.GET("/users", controllers.Index)
	r.GET("/user/:id", controllers.Show)
	r.POST("/users", controllers.Create)
	r.DELETE("users/:id", controllers.Delete)
	r.Run(":8080")
}