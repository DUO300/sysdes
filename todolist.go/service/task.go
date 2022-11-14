package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "todolist.go/db"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// TaskList renders list of tasks in DB
func TaskList(ctx *gin.Context) {
	// define pagesize
	const PAGESIZE = 5

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// Get query parameter
	kw := ctx.Query("kw")
	statusstr := ctx.Query("status")
	kw_h := ctx.Query("kw_h")
	status_h := ctx.Query("status_h")
	pagenum_str := ctx.Query("pagenum")
	search := ctx.Query("search")
	movpage := ctx.Query("movpage")

	// Get current page number
	var pagenum int
	if pagenum_str == "" {
		pagenum = 1
	} else {
		pagenum, err = strconv.Atoi(pagenum_str)
		if err != nil {
			Error(http.StatusInternalServerError, err.Error())(ctx)
			return
		}
	}

	// If "search" button is not pressed, then replace the search query to the old one
	// to avoid searching with different query when "<" or ">" is pressed
	if search == "" {
		kw = kw_h
		statusstr = status_h
	}
	if movpage == "<" {
		pagenum--
	} else if movpage == ">" {
		pagenum++
	}

	// Change statusstr (string) to 0, 1 or % (= 0 or 1)
	var status string
	switch statusstr {
	case "complete":
		status = "1"
	case "incomplete":
		status = "0"
	default:
		status = "%"
	}

	// Get tasks in DB
	var tasks []database.Task
	switch {
	case kw != "":
		err = db.Select(&tasks, "SELECT * FROM tasks WHERE title LIKE ? AND is_done LIKE ?", "%"+kw+"%", status)
	default:
		err = db.Select(&tasks, "SELECT * FROM tasks WHERE is_done LIKE ?", status)
	}
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// Check whether the page is the last or not to disable the ">" button
	is_lastpage := false
	if pagenum * PAGESIZE >= len(tasks) {
		is_lastpage = true
	}

	// Render tasks
	ctx.HTML(http.StatusOK, "task_list.html", gin.H{"Title": "Task list", "Tasks": tasks[(pagenum - 1) * PAGESIZE : min(pagenum * PAGESIZE, len(tasks))], "Kw": kw, "Status": statusstr, "Pagenum": pagenum, "Is_lastpage": is_lastpage})
}

// ShowTask renders a task with given ID
func ShowTask(ctx *gin.Context) {
	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// parse ID given as a parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		Error(http.StatusBadRequest, err.Error())(ctx)
		return
	}

	// Get a task with given ID
	var task database.Task
	err = db.Get(&task, "SELECT * FROM tasks WHERE id=?", id) // Use DB#Get for one entry
	if err != nil {
		Error(http.StatusBadRequest, err.Error())(ctx)
		return
	}

	// Render task
	// ctx.String(http.StatusOK, task.Title)
	ctx.HTML(http.StatusOK, "task.html", task)
}

func NewTaskForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form_new_task.html", gin.H{"Title": "Task registration"})
}

func RegisterTask(ctx *gin.Context) {
	// Get task title
	title, exist := ctx.GetPostForm("title")
	if !exist {
		Error(http.StatusBadRequest, "No title is given")(ctx)
		return
	}
	description, exist := ctx.GetPostForm("description")
	if !exist {
		Error(http.StatusBadRequest, "No description is given")(ctx)
		return
	}
	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Create new data with given title and description on DB
	result, err := db.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", title, description)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Render status
	path := "/list" // デフォルトではタスク一覧ページへ戻る
	if id, err := result.LastInsertId(); err == nil {
		path = fmt.Sprintf("/task/%d", id) // 正常にIDを取得できた場合は /task/<id> へ戻る
	}
	ctx.Redirect(http.StatusFound, path)
}

func EditTaskForm(ctx *gin.Context) {
	// ID の取得
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		Error(http.StatusBadRequest, err.Error())(ctx)
		return
	}
	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Get target task
	var task database.Task
	err = db.Get(&task, "SELECT * FROM tasks WHERE id=?", id)
	if err != nil {
		Error(http.StatusBadRequest, err.Error())(ctx)
		return
	}
	// Render edit form
	ctx.HTML(http.StatusOK, "form_edit_task.html",
		gin.H{"Title": fmt.Sprintf("Edit task %d", task.ID), "Task": task})
}

func UpdateTask(ctx *gin.Context) {
	// Get task id
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		Error(http.StatusBadRequest, err.Error())(ctx)
		return
	}
	// Get task title
	title, exist := ctx.GetPostForm("title")
	if !exist {
		Error(http.StatusBadRequest, "No title is given")(ctx)
		return
	}
	// Get task description
	description, exist := ctx.GetPostForm("description")
	if !exist {
		Error(http.StatusBadRequest, "No description is given")(ctx)
		return
	}
	// Get task is_done
	is_done, exist := ctx.GetPostForm("is_done")
	if !exist {
		Error(http.StatusBadRequest, "No checkmark is given")(ctx)
		return
	}
	is_done_bool, err := strconv.ParseBool(is_done)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Update DB
	_, err = db.Exec("UPDATE tasks SET title=?, is_done=?, description=? WHERE id=?", title, is_done_bool, description, id)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Render status
	path := fmt.Sprintf("/task/%d", id)
	ctx.Redirect(http.StatusFound, path)
}

func DeleteTask(ctx *gin.Context) {
	// ID の取得
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		Error(http.StatusBadRequest, err.Error())(ctx)
		return
	}
	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Delete the task from DB
	_, err = db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		Error(http.StatusInternalServerError, err.Error())(ctx)
		return
	}
	// Redirect to /list
	ctx.Redirect(http.StatusFound, "/list")
}
