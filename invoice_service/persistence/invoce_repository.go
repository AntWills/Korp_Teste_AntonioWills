package persistence

import (
	"fmt"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Create(invoice *Invoice) error {
	if err := r.db.Create(invoice).Error; err != nil {
		return fmt.Errorf("erro ao criar nota fiscal: %w", err)
	}
	return nil
}

func (r *InvoiceRepository) FindAll() ([]Invoice, error) {
	var invoices []Invoice
	if err := r.db.Preload("Items").Order("number asc").Find(&invoices).Error; err != nil {
		return nil, fmt.Errorf("erro ao listar notas fiscais: %w", err)
	}
	return invoices, nil
}

func (r *InvoiceRepository) FindByID(id uint) (*Invoice, error) {
	var invoice Invoice
	if err := r.db.Preload("Items").First(&invoice, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("nota fiscal %d não encontrada", id)
		}
		return nil, fmt.Errorf("erro ao buscar nota fiscal %d: %w", id, err)
	}
	return &invoice, nil
}

func (r *InvoiceRepository) UpdateStatus(id uint, status InvoiceStatus) error {
	result := r.db.Model(&Invoice{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return fmt.Errorf("erro ao atualizar status da nota fiscal %d: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("nota fiscal %d não encontrada para atualização", id)
	}
	return nil
}

func (r *InvoiceRepository) Save(invoice *Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *InvoiceRepository) DB() *gorm.DB {
	return r.db
}
