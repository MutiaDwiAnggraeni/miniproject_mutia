package models

import "gorm.io/gorm"

type ProductRecommen struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255)" json:"name"`
	Category    string  `gorm:"type:varchar(255)" json:"category"`
	Popularity  int     `gorm:"default:0" json:"popularity"` // Jumlah pembelian atau dilihat
	Rating      float64 `gorm:"default:0" json:"rating"`     // Skor rata-rata dari pengguna
	Price       float64 `gorm:"type:decimal(10,2)" json:"price"`
	Description string  `gorm:"type:text" json:"description"`
	Rekomendasi string  `json:"rekomendasi"` // AI-generated description
}
