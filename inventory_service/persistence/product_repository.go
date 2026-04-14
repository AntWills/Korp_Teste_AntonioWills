package persistence

import "gorm.io/gorm"

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Save(product *Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) FindByCode(code string) (*Product, error) {
	var product Product
	err := r.db.Where("code = ?", code).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepository) Update(product *Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) DeleteByCode(code string) error {
	return r.db.Where("code = ?", code).Delete(&Product{}).Error
}
