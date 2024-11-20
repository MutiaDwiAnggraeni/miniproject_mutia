// controllers/category_controller.go
package controllers

import (
	"net/http"
	"rest_api_muti/config"
	"rest_api_muti/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCategoryHandler(c echo.Context) error {
	category := models.Category{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Input tidak valid", Data: nil})
	}

	result := config.DB.Create(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal menambahkan kategori", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Kategori berhasil ditambahkan", Data: category})
}

func GetCategoryHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "ID tidak valid", Data: nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: false, Message: "Kategori tidak ditemukan", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Kategori ditemukan", Data: category})
}

func GetAllCategoriesHandler(c echo.Context) error {
	var categories []models.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal mengambil data kategori", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Data kategori berhasil diambil", Data: categories})
}

func UpdateCategoryHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "ID tidak valid", Data: nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: false, Message: "Kategori tidak ditemukan", Data: nil})
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Input tidak valid", Data: nil})
	}

	result = config.DB.Save(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal mengupdate kategori", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Kategori berhasil diperbarui", Data: category})
}

func DeleteCategoryHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "ID tidak valid", Data: nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: false, Message: "Kategori tidak ditemukan", Data: nil})
	}

	result = config.DB.Delete(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal menghapus kategori", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Kategori berhasil dihapus", Data: nil})
}
