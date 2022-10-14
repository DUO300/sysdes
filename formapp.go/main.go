package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// config
const port = 8000

func main() {
	// initialize Gin engine
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")
	engine.GET("/name-form", nameFormHandler)
	// engine.POST("/register-name", registerNameHandler)
	engine.GET("/register-name", registerNameHandler)

	// routing
	// engine.GET("/", rootHandler)
	engine.GET("/hello", helloHandler)
	engine.GET("/bye", byeHandler)
	engine.GET("/hello.jp", hellojpHandler)

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}

func nameFormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "name_form.html", nil)
}

func registerNameHandler(ctx *gin.Context) {
    name, _ := ctx.GetPostForm("name")
    ctx.HTML(http.StatusOK, "result.html", gin.H{"Name": name})
}

func rootHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world.")
}

func helloHandler(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "Hello world.")
	ctx.HTML(http.StatusOK, "hello.html", nil)
}

func byeHandler(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "Good bye.")
	ctx.HTML(http.StatusOK, "bye.html", nil)
}

func hellojpHandler(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "こんにちは")
	ctx.HTML(http.StatusOK, "hellojp.html", nil)
}
