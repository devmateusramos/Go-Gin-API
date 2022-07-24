package routes

import (
	"gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)

	r.Run()
}
