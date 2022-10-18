package stateless
 
import (
    "net/http"
    "github.com/gin-gonic/gin"
)
 
func Start(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "start.html", gin.H{ "Target": "/stateless/start" })
}

func NameForm(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "name-form.html", gin.H{ "Target": "/stateless/name" })
}

func BirthdayForm(ctx *gin.Context) {
    name, exist := ctx.GetPostForm("name")
    if !exist {
        ctx.String(http.StatusBadRequest, "parameter 'name' is not provided")
    }
    ctx.HTML(http.StatusOK, "stateless-birthday-form.html", gin.H{ "Name": name })
}

func MessageForm(ctx *gin.Context) {
    name, exist := ctx.GetPostForm("name")
	if !exist {
        ctx.String(http.StatusBadRequest, "parameter 'name' is not provided")
    }
	birthday, exist := ctx.GetPostForm("birthday")
    if !exist {
        ctx.String(http.StatusBadRequest, "parameter 'birthday' is not provided")
    }
    ctx.HTML(http.StatusOK, "stateless-message-form.html", gin.H{ "Name": name , "Birthday": birthday })
}

func Result(ctx *gin.Context) {
	name, exist := ctx.GetPostForm("name")
	if !exist {
        ctx.String(http.StatusBadRequest, "parameter 'name' is not provided")
    }
	birthday, exist := ctx.GetPostForm("birthday")
    if !exist {
        ctx.String(http.StatusBadRequest, "parameter 'birthday' is not provided")
    }
	message, exist := ctx.GetPostForm("message")
    if !exist {
        ctx.String(http.StatusBadRequest, "parameter 'message' is not provided")
    }
    ctx.HTML(http.StatusOK, "stateless-submit.html", gin.H{ "Name": name , "Birthday": birthday , "Message": message })
}