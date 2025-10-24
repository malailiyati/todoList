package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/malailiyati/todoList/internals/models"
	"github.com/malailiyati/todoList/internals/services"
	"github.com/malailiyati/todoList/internals/utils"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service}
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	search := c.Query("search")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10

	data, err := h.service.GetAll(search, page, limit)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Todos fetched successfully", data)
}

func (h *TodoHandler) Create(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := h.service.Create(&todo); err != nil {
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

	updated, err := h.service.Update(uint(id), &todo)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Todo updated successfully", updated)
}

func (h *TodoHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(uint(id)); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "Todo deleted successfully", nil)
}

func (h *TodoHandler) ToggleComplete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	updated, err := h.service.ToggleComplete(uint(id))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	status := "marked as incomplete"
	if updated.Completed {
		status = "marked as complete"
	}
	utils.Success(c, http.StatusOK, "Todo "+status, updated)
}

func (h *TodoHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.service.GetByID(uint(id))
	if err != nil {
		utils.Error(c, http.StatusNotFound, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "Todo fetched successfully", todo)
}
