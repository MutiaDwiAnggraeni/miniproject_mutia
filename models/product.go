package models

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id_product"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"type:decimal(10,2)" json:"price"`
	CategoryID  uint      `gorm:"not null" json:"id_category"`
	Category    string    `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt   time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:datetime" json:"updated_at"`
	Popularity  int       `gorm:"default:0" json:"popularity"`
	Rating      float64   `gorm:"default:0" json:"rating"`
}
