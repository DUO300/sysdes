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
	// Initialize DB connection
	dsn := db.DefaultDSN(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	if err := db.Connect(dsn); err != nil {
		log.Fatal(err)
	}

	// Initialize Gin engine
	engine := gin.Default()
	engine.LoadHTMLGlob("views/*.html")

	// Prepare session
	store := cookie.NewStore([]byte("my-secret"))
	engine.Use(sessions.Sessions("user-session", store))

	// Routing
	engine.Static("/assets", "./assets")
	engine.GET("/", service.Home)
	engine.GET("/list", service.LoginCheck, service.TaskList)

	taskGroup := engine.Group("/task")
	taskGroup.Use(service.LoginCheck)
	{
		// Show task info
		taskGroup.GET("/:id", service.TaskCheck, service.ShowTask) // ":id" is a parameter

		// Register new task
		taskGroup.GET("/new", service.NewTaskForm)
		taskGroup.POST("/new", service.RegisterTask)

		// Edit task
		taskGroup.GET("/edit/:id", service.TaskCheck, service.EditTaskForm)
		taskGroup.POST("/edit/:id", service.TaskCheck, service.UpdateTask)

		// Delete task
		taskGroup.GET("/delete/:id", service.TaskCheck, service.DeleteTask)
	}

	// Register new user
	engine.GET("/user/new", service.NewUserForm)
	engine.POST("/user/new", service.RegisterUser)

	// User login
	engine.GET("/login", service.LoginForm)
	engine.POST("/login", service.Login)

	// User logout
	engine.GET("/logout", service.LoginCheck, service.Logout)

	// Delete user account
	engine.GET("/delete_user", service.LoginCheck, service.DeleteUser)

	// Show user info
	engine.GET("/user/info", service.LoginCheck, service.ShowUser)

	// Edit user account
	engine.GET("/user/edit", service.LoginCheck, service.EditUser)
	engine.POST("/user/edit", service.LoginCheck, service.UpdateUser)

	// Start server
	engine.Run(fmt.Sprintf(":%d", port))
}
