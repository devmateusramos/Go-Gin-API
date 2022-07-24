package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExibeAlunos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "1",
		"nome": "Mateus",
	})
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"api diz": "Olá " + nome + ", tudo beleza?",
	})
}
