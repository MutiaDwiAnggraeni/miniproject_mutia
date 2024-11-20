// controllers/category_controller.go
package controllers

import (
	"net/http"
	"rest_api_muti/config"
	"rest_api_muti/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductHandler(c echo.Context) error {
	product := models.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Input tidak valid", Data: nil})
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal menambahkan product", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "product berhasil ditambahkan", Data: product})
}

func GetProductHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "ID tidak valid", Data: nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: false, Message: "product tidak ditemukan", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "product ditemukan", Data: category})
}

func GetAllProductsHandler(c echo.Context) error {
	var categories []models.Product
	result := config.DB.Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal mengambil data Product", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Data Product berhasil diambil", Data: categories})
}

func UpdateProductsHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "ID tidak valid", Data: nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: false, Message: "Product tidak ditemukan", Data: nil})
	}

	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "Input tidak valid", Data: nil})
	}

	result = config.DB.Save(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal mengupdate Product", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Product berhasil diperbarui", Data: category})
}

func DeleteProductHandler(c echo.Context) error {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{Status: false, Message: "ID tidak valid", Data: nil})
	}

	category := models.Category{}
	result := config.DB.First(&category, idNumber)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{Status: false, Message: "Product tidak ditemukan", Data: nil})
	}

	result = config.DB.Delete(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: "Gagal menghapus Product", Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Product berhasil dihapus", Data: nil})
}
