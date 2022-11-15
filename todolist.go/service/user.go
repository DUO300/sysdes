package service

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	database "todolist.go/db"
)

func NewUserForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "new_user_form.html", gin.H{"Title": "Register user"})
}

func hash(pw string) []byte {
	const salt = "todolist.go#"
	h := sha256.New()
	h.Write([]byte(salt))
	h.Write([]byte(pw))
	return h.Sum(nil)
}

func RegisterUser(ctx *gin.Context) {
	// フォームデータの受け取り
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	password2 := ctx.PostForm("password2")
	switch {
	case username == "":
		ctx.HTML(http.StatusBadRequest, "new_user_form.html", gin.H{"Title": "Register user", "Error": "Usernane is not provided", "Username": username})
		return
	case password == "":
		ctx.HTML(http.StatusBadRequest, "new_user_form.html", gin.H{"Title": "Register user", "Error": "Password is not provided", "Password": password})
		return
	case password2 == "":
		ctx.HTML(http.StatusBadRequest, "new_user_form.html", gin.H{"Title": "Register user", "Error": "Password is not provided", "Password2": password2})
		return
	case password != password2:
		ctx.HTML(http.StatusBadRequest, "new_user_form.html", gin.H{"Title": "Register user", "Error": "Passwords do not match", "Username": username})
		return
	}

	// パスワードのチェック
	if len(password) <= 4 {
		ctx.HTML(http.StatusBadRequest, "new_user_form.html", gin.H{"Title": "Register user", "Error": "Password too short. Password length must be at least 5", "Username": username})
		return
	}
	if re := regexp.MustCompile(`^[0-9]*$`); re.MatchString(password) {
		ctx.HTML(http.StatusBadRequest, "new_user_form.html", gin.H{"Title": "Register user", "Error": "Password must contain at least one alphabet", "Username": username})
		return
	}

	// DB 接続
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// 重複チェック
	var duplicate int
	err = db.Get(&duplicate, "SELECT COUNT(*) FROM users WHERE name=? AND valid=1", username)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	if duplicate > 0 {
		ctx.HTML(http.StatusBadRequest, "new_user_form.html", gin.H{"Title": "Register user", "Error": "Username is already taken", "Username": username, "Password": password, "Password2": password2})
		return
	}

	// DB への保存
	result, err := db.Exec("INSERT INTO users(name, password) VALUES (?, ?)", username, hash(password))
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// 保存状態の確認
	id, _ := result.LastInsertId()
	var user database.User
	err = db.Get(&user, "SELECT id, name, password, updated_at FROM users WHERE id = ? AND valid=1", id)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	//ctx.JSON(http.StatusOK, user)
	ctx.Redirect(http.StatusFound, "/list")
}

func LoginForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{"Title": "User Login"})
}

const userkey = "user"

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// ユーザの取得
	var user database.User
	err = db.Get(&user, "SELECT id, name, password FROM users WHERE name = ? AND valid=1", username)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "login.html", gin.H{"Title": "Login", "Username": username, "Error": "No such user"})
		return
	}

	// パスワードの照合
	if hex.EncodeToString(user.Password) != hex.EncodeToString(hash(password)) {
		ctx.HTML(http.StatusBadRequest, "login.html", gin.H{"Title": "Login", "Username": username, "Error": "Incorrect password"})
		return
	}

	// セッションの保存
	session := sessions.Default(ctx)
	session.Set(userkey, user.ID)
	session.Save()

	ctx.Redirect(http.StatusFound, "/list")
}

func LoginCheck(ctx *gin.Context) {
	if sessions.Default(ctx).Get(userkey) == nil {
		ctx.Redirect(http.StatusFound, "/login")
		ctx.Abort()
	} else {
		userID := sessions.Default(ctx).Get("user")
		// Get DB connection
		db, err := database.GetConnection()
		if err != nil {
			Error(http.StatusInternalServerError, err.Error())(ctx)
			ctx.Abort()
		}
		// Get target user
		var user database.User
		err = db.Get(&user, "SELECT * FROM users WHERE id=? AND valid=1", userID)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/login")
			ctx.Abort()
		}
		ctx.Next()
	}
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	ctx.Redirect(http.StatusFound, "/")
}

func DeleteUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID := session.Get("user")

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Update DB
	_, err = db.Exec("UPDATE users SET valid=0 WHERE id=?", userID)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	ctx.Redirect(http.StatusFound, "/")
}

func ShowUser(ctx *gin.Context) {

	userID := sessions.Default(ctx).Get("user")

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// Get target user
	var user database.User
	err = db.Get(&user, "SELECT * FROM users WHERE id=? AND valid=1", userID) // Use DB#Get for one entry
	if err != nil {
		Error(http.StatusBadRequest, err.Error())(ctx)
		return
	}

	// Render task
	ctx.HTML(http.StatusOK, "user.html", gin.H{"Title": "User", "User": user})
}

func EditUser(ctx *gin.Context) {

	userID := sessions.Default(ctx).Get("user")

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// Get target user
	var user database.User
	err = db.Get(&user, "SELECT * FROM users WHERE id=? AND valid=1", userID)
	if err != nil {
		Error(http.StatusBadRequest, err.Error())(ctx)
		return
	}

	// Render edit form
	ctx.HTML(http.StatusOK, "form_edit_user.html",
		gin.H{"Title": "Edit User", "Username": user.Name})
}

func UpdateUser(ctx *gin.Context) {

	userID := sessions.Default(ctx).Get("user")

	// フォームデータの受け取り
	username := ctx.PostForm("username")
	password_old := ctx.PostForm("password_old")
	password_new1 := ctx.PostForm("password_new1")
	password_new2 := ctx.PostForm("password_new2")
	switch {
	case username == "":
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "Usernane is not provided", "Username": username})
		return
	case password_old == "":
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "Old password is not provided", "Username": username, "Password_old": password_old})
		return
	case password_new1 == "":
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "New Password is not provided", "Username": username, "Password_old": password_old, "Password_new1": password_new1})
		return
	case password_new2 == "":
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "New Password is not provided", "Username": username, "Password_old": password_old, "Password_new2": password_new2})
		return
	case password_new1 != password_new2:
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "New passwords do not match", "Username": username, "Password_old": password_old})
		return
	}

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// ユーザの取得
	var user database.User
	err = db.Get(&user, "SELECT id, name, password FROM users WHERE id=? AND valid=1", userID)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit User", "Error": "No such user"})
		return
	}

	// パスワードの照合
	if hex.EncodeToString(user.Password) != hex.EncodeToString(hash(password_old)) {
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "Incorrect old password", "Username": username})
		return
	}

	// パスワードのチェック
	if len(password_new1) <= 4 {
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "New password too short. Password length must be at least 5", "Username": username, "Password_old": password_old})
		return
	}
	if re := regexp.MustCompile(`^[0-9]*$`); re.MatchString(password_new1) {
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "New password must contain at least one alphabet", "Username": username, "Password_old": password_old})
		return
	}

	// 重複チェック
	var duplicate int
	err = db.Get(&duplicate, "SELECT COUNT(*) FROM users WHERE name=? AND valid=1 AND id!=?", username, userID)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	if duplicate > 0 {
		ctx.HTML(http.StatusBadRequest, "form_edit_user.html", gin.H{"Title": "Edit user", "Error": "Username is already taken", "Username": username, "Password_old": password_old, "Password_new1": password_new1, "Password_new2": password_new2})
		return
	}

	// Update DB
	_, err = db.Exec("UPDATE users SET name=?, password=? WHERE id=?", username, hash(password_new1), userID)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Render status
	ctx.Redirect(http.StatusFound, "/user")
}
