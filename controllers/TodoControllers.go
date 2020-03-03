package controllers

import(
	 "net/http"
	 "strconv"

	 Gdb "todoapi/databases"

	 "github.com/gin-gonic/gin"
	 "github.com/jinzhu/gorm"
)

 type(
	 todoModel struct{
		 gorm.Model
		 Title string `json:"title"`
		 Completed int `json:"completed"`
	 }
)
// A PetBodyParam
//
// This is used for operations that want an Order as body of the request
// swagger:parameters ID IID IDI
 type(
	 transformedTodo struct {
	// 成功
 	//
	// required: true
	// In: path
		 ID uint `json:"id"`
		 Title string `json:"title"`
		 Completed bool `json:"completed"`
	 }
 )
// swagger:route POST /todolist/ Todo Todo
//
// 創建新待辦事項1
//
// 創新新待辦事項2
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: Successful 
//		 404: ErrorMessage
func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{Title: c.PostForm("title"), Completed: completed}
	Gdb.Db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created succesfully!", "resourceId": todo.ID})
}

// swagger:route GET /todolist/ Todo Todo
//
// 全待辦事項1
//
// 全待辦事項2
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: Successful 
//		 404: ErrorMessage
func GetAllTodo(c *gin.Context){
	var todolist []todoModel
	var _todolist []transformedTodo
	
	Gdb.Db.Find(&todolist)	
	
	if len(todolist) <= 0 {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
	return
}
	for _, item := range todolist {
		completed := false
		if item.Completed == 1 {
			completed = true
	} 	else {
			completed = false
	}
	_todolist = append(_todolist, transformedTodo{ID: item.ID, Title: item.Title, Completed:completed})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todolist})
	
}

// swagger:route GET /todolist/{id} Todo IDI
//
// 單待辦事項1
//
// 單待辦事項2
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: Successful 
//		 404: ErrorMessage
func GetSingleTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")
	
	Gdb.Db.First(&todo, todoID)
	
	if todo.ID == 0 {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, })
	return	
}
	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}
	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// swagger:route PUT /todolist/{id} Todo IID
//
// 編輯待辦事項1
//
// 編輯待辦事項2
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: Successful 
//		 404: ErrorMessage
func UpdateTodo(c *gin.Context){
	var todo todoModel
	todoID := c.Param("id")

	Gdb.Db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	Gdb.Db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	Gdb.Db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// swagger:route DELETE /todolist/{id} Todo ID
//
// 刪除待辦事項1
//
// 刪除待辦事項2
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: Successful 
//       404: ErrorMessage
func DeleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	Gdb.Db.First(&todo, todoID)

	if todo.ID == 0 {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
	return
}	
	Gdb.Db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}