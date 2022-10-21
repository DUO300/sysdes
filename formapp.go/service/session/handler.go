package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "start.html", gin.H{"Target": "/session/start"})
}

func NameForm(ctx *gin.Context) {
	session, err := NewSession()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Fail to create a new session")
		return
	}
	ctx.SetCookie("userid", session.ID(), 600, "/session/", "localhost:8000", false, false)
	ctx.HTML(http.StatusOK, "name-form.html", gin.H{"Target": "/session/name", "Target2": "/session/start"})
}

func BirthdayForm(ctx *gin.Context) {
	id, err := ctx.Cookie("userid")
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid access")
		return
	}
	session := Session{id}
	name, exist := ctx.GetPostForm("name")
	state, _ := session.GetState()
	/*
		if !exist {
			ctx.String(http.StatusBadRequest, "parameter 'name' is not provided")
			return
		}*/
	if exist {
		state.Name = name
		session.SetState(state)
	}
	ctx.HTML(http.StatusOK, "session-birthday-form.html", nil)
}

func MessageForm(ctx *gin.Context) {
	id, err := ctx.Cookie("userid")
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid access")
		return
	}
	session := Session{id}
	birthday, exist := ctx.GetPostForm("birthday")
	state, _ := session.GetState()
	/*
		if !exist {
			ctx.String(http.StatusBadRequest, "parameter 'birthday' is not provided")
			return
		}*/
	if exist {
		state.Birthday = birthday
		session.SetState(state)
	}
	ctx.HTML(http.StatusOK, "session-message-form.html", nil)
}

func SubmitForm(ctx *gin.Context) {
	id, err := ctx.Cookie("userid")
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid access")
		return
	}
	session := Session{id}
	message, exist := ctx.GetPostForm("message")
	state, _ := session.GetState()
	/*
		if !exist {
			ctx.String(http.StatusBadRequest, "parameter 'message' is not provided")
			return
		} */
	if exist {
		state.Message = message
		session.SetState(state)
	}

	ctx.HTML(http.StatusOK, "session-submit.html", gin.H{"Name": state.Name, "Birthday": state.Birthday, "Message": state.Message})
}
