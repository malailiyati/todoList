package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/malailiyati/todoList/internals/models"
	"github.com/malailiyati/todoList/internals/repositories"
	"github.com/malailiyati/todoList/internals/utils"
)

type CategoryHandler struct {
	repo *repositories.CategoryRepository
}

func NewCategoryHandler(repo *repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	data, err := h.repo.GetAll()
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

	// Cek apakah nama kategori sudah ada
	exists, err := h.repo.CheckNameExists(category.Name)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to check category name")
		return
	}
	if exists {
		utils.Error(c, http.StatusBadRequest, "Category name already exists")
		return
	}

	// Simpan ke DB
	if err := h.repo.Create(&category); err != nil {
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

	// Cek dulu apakah category ada
	if _, err := h.repo.FindByID(uint(id)); err != nil {
		utils.Error(c, http.StatusNotFound, "Category not found")
		return
	}

	updated, err := h.repo.Update(uint(id), &category)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Category updated successfully", updated)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if _, err := h.repo.FindByID(uint(id)); err != nil {
		utils.Error(c, http.StatusNotFound, "Category not found")
		return
	}

	if err := h.repo.Delete(uint(id)); err != nil {
		utils.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(c, http.StatusOK, "Category deleted successfully", nil)
}
