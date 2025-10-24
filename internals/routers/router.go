package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/malailiyati/todoList/internals/handlers"
	"github.com/malailiyati/todoList/internals/repositories"
	"github.com/malailiyati/todoList/internals/services"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	categoryRepo := repositories.NewCategoryRepository(db)
	todoRepo := repositories.NewTodoRepository(db)

	categoryService := services.NewCategoryService(categoryRepo)
	todoService := services.NewTodoService(todoRepo, categoryRepo)

	categoryHandler := handlers.NewCategoryHandler(categoryService)
	todoHandler := handlers.NewTodoHandler(todoService)

	api := r.Group("/api")
	{
		api.GET("/categories", categoryHandler.GetAll)
		api.POST("/categories", categoryHandler.Create)
		api.PATCH("/categories/:id", categoryHandler.Update)
		api.DELETE("/categories/:id", categoryHandler.Delete)

		api.GET("/todos", todoHandler.GetAll)
		api.POST("/todos", todoHandler.Create)
		api.GET("/todos/:id", todoHandler.GetByID)
		api.PATCH("/todos/:id", todoHandler.Update)
		api.DELETE("/todos/:id", todoHandler.Delete)
		api.PATCH("/todos/:id/complete", todoHandler.ToggleComplete)
	}

	return r
}
