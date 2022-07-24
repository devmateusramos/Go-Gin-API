package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", exibeAlunos)

	r.Run()
}

func exibeAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "1",
		"nome": "Mateus",
	})
}
