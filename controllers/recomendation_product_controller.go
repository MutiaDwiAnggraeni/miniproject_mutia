package controllers

import (
	"context"
	"net/http"
	"rest_api_muti/config"
	"rest_api_muti/helper"
	"rest_api_muti/models"

	"github.com/labstack/echo/v4"
)

// RecommendationRequest menerima input dari user
type RecommendationRequest struct {
	Category string  `json:"category"`
	MinPrice float64 `json:"min_price"`
	MaxPrice float64 `json:"max_price"`
	Limit    int     `json:"limit"`
}

// GetRecommendationsWithAI mengelola rekomendasi produk berdasarkan input user dan AI
func GetRecommendationsWithAI(c echo.Context) error {
	var products []models.ProductRecommen
	var request RecommendationRequest

	// Bind input JSON ke struct RecommendationRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Format input tidak valid",
		})
	}

	// Query produk dari database
	query := config.DB

	// Tambahkan filter berdasarkan kategori
	if request.Category != "" {
		query = query.Where("category = ?", request.Category)
	}

	// Tambahkan filter berdasarkan harga
	if request.MinPrice > 0 {
		query = query.Where("price >= ?", request.MinPrice)
	}
	if request.MaxPrice > 0 {
		query = query.Where("price <= ?", request.MaxPrice)
	}

	// Urutkan produk berdasarkan popularitas dan rating
	err := query.Order("popularity DESC, rating DESC").
		Limit(request.Limit).
		Find(&products).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal mendapatkan rekomendasi produk",
		})
	}

	// Proses deskripsi menggunakan AI
	ctx := context.Background()
	if err := helper.GenerateDescriptions(ctx, products); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal menghasilkan deskripsi produk menggunakan AI",
		})
	}

	// Kembalikan produk dengan deskripsi ke pengguna
	return c.JSON(http.StatusOK, products)
}
