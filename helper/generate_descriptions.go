package helper

import (
	"context"
	"fmt"
	"log"
	"rest_api_muti/config"
	"rest_api_muti/models"
)

// GenerateDescriptions iterasi pada entri ProductRecommen dan tambahkan deskripsi yang dihasilkan AI
func GenerateDescriptions(ctx context.Context, products []models.ProductRecommen) error {
	for i := range products {
		// Buat query untuk AI berdasarkan data produk
		query := GenerateProductQuery(products[i])

		// Dapatkan deskripsi rekomendasi dari Gemini API
		rekomendasi, err := ResponseAI(ctx, query)
		if err != nil {
			log.Printf("Gagal mendapatkan rekomendasi AI untuk produk %s: %v", products[i].Name, err)
			continue
		}

		// Tambahkan rekomendasi ke deskripsi produk
		products[i].Description = fmt.Sprintf(
			"Deskripsi Produk: %s\n\nRekomendasi AI: %s",
			products[i].Description,
			rekomendasi,
		)

		fmt.Printf("Produk %s:\n%s\n", products[i].Name, products[i].Description)

		// Simpan produk ke database
		if err := config.DB.Save(&products[i]).Error; err != nil {
			log.Printf("Gagal menyimpan deskripsi untuk produk %s: %v", products[i].Name, err)
		}
	}
	return nil
}
