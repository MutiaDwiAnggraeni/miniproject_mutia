// helper/query_helper.go
package helper

import (
	"fmt"
	"rest_api_muti/models"
)

// GenerateProductQuery membuat query deskripsi produk berdasarkan informasi dari model ProductRecommen.
func GenerateProductQuery(product models.ProductRecommen) string {
	var query string

	if product.Category == "Organik" {
		query = fmt.Sprintf(
			"Produk '%s' adalah kategori '%s', memiliki rating %.1f, dan harga %.2f. Karena termasuk kategori organik, "+
				"jelaskan dampak positif produk ini terhadap lingkungan, rekomendasi pengurangan limbah, serta manfaat "+
				"penggunaannya. Deskripsikan juga bagaimana produk ini dapat mendukung gaya hidup yang lebih ramah lingkungan.",
			product.Name, product.Category, product.Rating, product.Price,
		)
	} else {
		query = fmt.Sprintf(
			"Produk '%s' adalah kategori '%s', memiliki rating %.1f, dan harga %.2f. Jelaskan dampak lingkungannya, "+
				"penggunaannya, dan manfaat produk ini. Sertakan rekomendasi untuk pengguna produk ini agar lebih efektif dan efisien.",
			product.Name, product.Category, product.Rating, product.Price,
		)
	}

	return query
}
