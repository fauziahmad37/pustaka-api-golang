package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", booksHandler)

	router.Run()
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Ahmad Fauzi",
		"bio":  "a software enginering",
	})
}

func helloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Ahmad Fauzi",
		"subtitle": "belajar golang bersama ahmad fauzi",
	})
}

func booksHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
