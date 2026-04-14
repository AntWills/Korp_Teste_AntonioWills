package services

import (
	"errors"
	"inventory_service/persistence"
)

type ProductService struct {
	repo *persistence.ProductRepository
}

func NewProductService(repo *persistence.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(code, description string, balance int) (*persistence.Product, error) {
	if code == "" || description == "" {
		return nil, errors.New("code and description are required")
	}
	
	// Check if already exists
	_, err := s.repo.FindByCode(code)
	if err == nil {
		return nil, errors.New("product with this code already exists")
	}

	product := &persistence.Product{
		Code:        code,
		Description: description,
		Balance:     balance,
	}

	if err := s.repo.Save(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetAllProducts() ([]persistence.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetProductByCode(code string) (*persistence.Product, error) {
	return s.repo.FindByCode(code)
}

func (s *ProductService) UpdateProductBalance(code string, balance int) (*persistence.Product, error) {
	product, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	product.Balance = balance
	if err := s.repo.Update(product); err != nil {
		return nil, err
	}
	
	return product, nil
}
