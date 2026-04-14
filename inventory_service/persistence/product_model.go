package persistence

import "gorm.io/gorm"

// Product representa a entidade Produto no banco de dados
type Product struct {
	gorm.Model
	Code        string `gorm:"uniqueIndex;not null" json:"code"`
	Description string `gorm:"not null" json:"description"`
	Balance     int    `gorm:"not null;default:0" json:"balance"`
}
