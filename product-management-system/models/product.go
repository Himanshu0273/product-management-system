package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    UserID                 uint
    ProductName            string
    ProductDescription     string
    ProductPrice           float64
    ProductImages          []string `gorm:"type:text[]"`
    CompressedProductImages []string `gorm:"type:text[]"`
}
