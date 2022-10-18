package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"formapp.go/service"
	"formapp.go/service/stateless"
)

// config
const port = 8000

func main() {
	// initialize Gin engine
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*.html")
	engine.GET("/name-form", service.NameFormHandler)
	// engine.POST("/register-name", registerNameHandler)
	engine.GET("/register-name", service.RegisterNameHandler)

	// routing
	// engine.GET("/", rootHandler)
	engine.GET("/hello", service.HelloHandler)
	engine.GET("/bye", service.ByeHandler)
	engine.GET("/hello.jp", service.HellojpHandler)

	engine.GET("/stateless/start", stateless.Start)
    engine.POST("/stateless/start", stateless.NameForm)
    engine.POST("/stateless/name", stateless.BirthdayForm)
    engine.POST("/stateless/birthday", stateless.MessageForm)
    engine.POST("/stateless/message", stateless.Result)

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}

func notImplemented(ctx *gin.Context) {
    msg := fmt.Sprintf("%s to %s is not implemented yet", ctx.Request.Method, ctx.Request.URL)
    ctx.String(http.StatusNotImplemented, msg)
}

/*
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
*/
