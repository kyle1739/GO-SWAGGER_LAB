package router

import(
	"github.com/gin-gonic/gin"
	"todoapi/controllers"
	"github.com/gin-contrib/cors"

)

func InitRouter() *gin.Engine {
	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
    	AllowHeaders:  []string{"content-type"},
    	ExposeHeaders: []string{"X-Total-Count"},
	}))

	v1 := router.Group("/api/v1/todolist")
	{
		v1.POST("/", controllers.CreateTodo)
		v1.GET("/", controllers.GetAllTodo)
		v1.GET("/:id", controllers.GetSingleTodo)
		v1.PUT("/:id", controllers.UpdateTodo)
		v1.DELETE("/:id", controllers.DeleteTodo)
	}

 	return router
 }