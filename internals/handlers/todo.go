package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/malailiyati/todoList/internals/models"
	"github.com/malailiyati/todoList/internals/repositories"
	"github.com/malailiyati/todoList/internals/utils"
)

type TodoHandler struct {
	repo *repositories.TodoRepository
}

func NewTodoHandler(repo *repositories.TodoRepository) *TodoHandler {
	return &TodoHandler{repo}
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	search := c.Query("search")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10
	offset := (page - 1) * limit

	todos, total, err := h.repo.GetAll(search, limit, offset)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	pagination := utils.NewPagination(page, limit, int(total))
	utils.Success(c, http.StatusOK, "Todos fetched successfully", gin.H{
		"todos":      todos,
		"pagination": pagination,
	})
}

func (h *TodoHandler) Create(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validate category_id if provided
	if todo.CategoryID != 0 {
		category, err := h.repo.CheckCategoryExists(todo.CategoryID)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to check category")
			return
		}
		if !category {
			utils.Error(c, http.StatusBadRequest, "Invalid category_id")
			return
		}
	}

	// Validate priority value
	if todo.Priority != "" &&
		todo.Priority != "high" &&
		todo.Priority != "medium" &&
		todo.Priority != "low" {
		utils.Error(c, http.StatusBadRequest, "Invalid priority (must be high, medium, or low)")
		return
	}

	// Save to DB
	if err := h.repo.Create(&todo); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusCreated, "Todo created successfully", todo)
}

func (h *TodoHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validate category_id if provided
	if todo.CategoryID != 0 {
		category, err := h.repo.CheckCategoryExists(todo.CategoryID)
		if err != nil {
			utils.Error(c, http.StatusInternalServerError, "Failed to check category")
			return
		}
		if !category {
			utils.Error(c, http.StatusBadRequest, "Invalid category_id")
			return
		}
	}

	// Validate priority if provided
	if todo.Priority != "" &&
		todo.Priority != "high" &&
		todo.Priority != "medium" &&
		todo.Priority != "low" {
		utils.Error(c, http.StatusBadRequest, "Invalid priority (must be high, medium, or low)")
		return
	}

	// Check if todo exists
	existing, err := h.repo.FindByID(uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Todo not found")
		return
	}

	// Preserve ID
	todo.ID = existing.ID

	updatedTodo, err := h.repo.Update(uint(id), &todo)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Todo updated successfully", updatedTodo)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if _, err := h.repo.FindByID(uint(id)); err != nil {
		utils.Error(c, http.StatusNotFound, "Todo not found")
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Todo deleted successfully", nil)
}

func (h *TodoHandler) ToggleComplete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := h.repo.FindByID(uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, "Todo not found")
		return
	}

	// Toggle status
	todo.Completed = !todo.Completed

	updated, err := h.repo.Update(uint(id), todo)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to update todo status")
		return
	}

	status := "marked as incomplete"
	if updated.Completed {
		status = "marked as complete"
	}
	utils.Success(c, http.StatusOK, "Todo "+status, updated)
}
