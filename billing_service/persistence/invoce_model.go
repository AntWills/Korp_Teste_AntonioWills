package persistence

import "gorm.io/gorm"

// InvoiceStatus representa o estado possível de uma Nota Fiscal.
type InvoiceStatus string

const (
	StatusOpen   InvoiceStatus = "Aberta"
	StatusClosed InvoiceStatus = "Fechada"
)

type Invoice struct {
	gorm.Model
	Number uint          `gorm:"uniqueIndex;autoIncrement" json:"number"`
	Status InvoiceStatus `gorm:"type:varchar(10);not null;default:'Aberta'"  json:"status"`
	Items  []InvoiceItem `gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE" json:"items"`
}

type InvoiceItem struct {
	gorm.Model
	InvoiceID   uint   `json:"invoice_id"`
	ProductCode string `gorm:"type:varchar(50);not null" json:"product_code"`
	Quantity    int    `gorm:"not null"                  json:"quantity"`
}
