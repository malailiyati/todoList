package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/malailiyati/todoList/internals/models"
	"github.com/malailiyati/todoList/internals/services"
	"github.com/malailiyati/todoList/internals/utils"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, http.StatusOK, "Categories fetched successfully", data)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := h.service.Create(&category); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusCreated, "Category created successfully", category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	updated, err := h.service.Update(uint(id), &category)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Category updated successfully", updated)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.Delete(uint(id)); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Category deleted successfully", nil)
}
