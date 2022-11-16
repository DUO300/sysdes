package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"todolist.go/db"
	"todolist.go/service"
)

const port = 8000

func main() {
	// initialize DB connection
	dsn := db.DefaultDSN(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	if err := db.Connect(dsn); err != nil {
		log.Fatal(err)
	}

	// initialize Gin engine
	engine := gin.Default()
	engine.LoadHTMLGlob("views/*.html")

	// prepare session
	store := cookie.NewStore([]byte("my-secret"))
	engine.Use(sessions.Sessions("user-session", store))

	// routing
	engine.Static("/assets", "./assets")
	engine.GET("/", service.Home)
	engine.GET("/list", service.LoginCheck, service.TaskList)

	taskGroup := engine.Group("/task")
	taskGroup.Use(service.LoginCheck)
	{
		taskGroup.GET("/:id", service.TaskCheck, service.ShowTask) // ":id" is a parameter

		// タスクの新規登録
		taskGroup.GET("/new", service.NewTaskForm)
		taskGroup.POST("/new", service.RegisterTask)

		// 既存タスクの編集
		taskGroup.GET("/edit/:id", service.TaskCheck, service.EditTaskForm)
		taskGroup.POST("/edit/:id", service.TaskCheck, service.UpdateTask)

		// 既存タスクの削除
		taskGroup.GET("/delete/:id", service.TaskCheck, service.DeleteTask)
	}

	// ユーザ登録
	engine.GET("/user/new", service.NewUserForm)
	engine.POST("/user/new", service.RegisterUser)

	// ユーザログイン
	engine.GET("/login", service.LoginForm)
	engine.POST("/login", service.Login)
	engine.GET("/logout", service.LoginCheck, service.Logout)
	engine.GET("/delete_user", service.LoginCheck, service.DeleteUser)
	engine.GET("/user/info", service.LoginCheck, service.ShowUser)
	engine.GET("/user/edit", service.LoginCheck, service.EditUser)
	engine.POST("/user/edit", service.LoginCheck, service.UpdateUser)

	// start server
	engine.Run(fmt.Sprintf(":%d", port))
}
