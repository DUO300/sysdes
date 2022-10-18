package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func NameFormHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "name_form.html", nil)
}

func RegisterNameHandler(ctx *gin.Context) {
    name, _ := ctx.GetPostForm("name")
    ctx.HTML(http.StatusOK, "result.html", gin.H{"Name": name})
}

func RootHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world.")
}

func HelloHandler(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "Hello world.")
	ctx.HTML(http.StatusOK, "hello.html", nil)
}

func ByeHandler(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "Good bye.")
	ctx.HTML(http.StatusOK, "bye.html", nil)
}

func HellojpHandler(ctx *gin.Context) {
	// ctx.String(http.StatusOK, "こんにちは")
	ctx.HTML(http.StatusOK, "hellojp.html", nil)
}