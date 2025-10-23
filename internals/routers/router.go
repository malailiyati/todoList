package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/malailiyati/todoList/internals/handlers"
	"github.com/malailiyati/todoList/internals/repositories"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize repository & handler
	categoryRepo := repositories.NewCategoryRepository(db)
	todoRepo := repositories.NewTodoRepository(db)
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)
	todoHandler := handlers.NewTodoHandler(todoRepo)

	// Routes
	api := r.Group("/api")
	{
		api.GET("/categories", categoryHandler.GetAll)
		api.POST("/categories", categoryHandler.Create)
		api.PATCH("/categories/:id", categoryHandler.Update)
		api.DELETE("/categories/:id", categoryHandler.Delete)

		api.GET("/todos", todoHandler.GetAll)
		api.POST("/todos", todoHandler.Create)
		api.PATCH("/todos/:id", todoHandler.Update)
		api.DELETE("/todos/:id", todoHandler.Delete)
		api.PATCH("/todos/:id/complete", todoHandler.ToggleComplete)
	}

	return r
}
