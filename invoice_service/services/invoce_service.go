package services

import (
	"errors"
	"fmt"

	"billing_service/persistence"
)

// Erros de domínio
var (
	ErrInvoiceNotFound      = errors.New("nota fiscal não encontrada")
	ErrInvoiceNotOpen       = errors.New("a nota fiscal não está com status 'Aberta' e não pode ser impressa")
	ErrInventoryUnavailable = errors.New("serviço de estoque indisponível; tente novamente mais tarde")
	ErrInvalidItems         = errors.New("a nota fiscal deve conter ao menos um produto")
)

// --- DTOs de entrada ---
type InvoiceItemInput struct {
	ProductCode string `json:"product_code" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required,min=1"`
}

type CreateInvoiceInput struct {
	Items []InvoiceItemInput `json:"items" binding:"required,min=1,dive"`
}

// InvoiceService é a estrutura concreta do serviço (sem interface)
type InvoiceService struct {
	repo            *persistence.InvoiceRepository
	inventoryClient InventoryClient
}

// NewInvoiceService cria uma nova instância do serviço
func NewInvoiceService(repo *persistence.InvoiceRepository, client InventoryClient) *InvoiceService {
	return &InvoiceService{
		repo:            repo,
		inventoryClient: client,
	}
}

// CreateInvoice valida a entrada, seta o status inicial como "Aberta" e persiste.
func (s *InvoiceService) CreateInvoice(input CreateInvoiceInput) (*persistence.Invoice, error) {
	if len(input.Items) == 0 {
		return nil, ErrInvalidItems
	}

	invoice := &persistence.Invoice{
		Status: persistence.StatusOpen,
	}

	for _, item := range input.Items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("quantidade inválida (%d) para o produto '%s'", item.Quantity, item.ProductCode)
		}

		invoice.Items = append(invoice.Items, persistence.InvoiceItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	if err := s.repo.Create(invoice); err != nil {
		return nil, err
	}

	return invoice, nil
}

// PrintInvoice executa o fluxo completo de impressão:
// 1. Verifica se a nota está Aberta
// 2. Solicita dedução de estoque ao inventory_service
// 3. Atualiza o status para Fechada apenas se o estoque foi deduzido com sucesso
func (s *InvoiceService) PrintInvoice(id uint) (*persistence.Invoice, error) {
	invoice, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrInvoiceNotFound
	}

	if invoice.Status != persistence.StatusOpen {
		return nil, ErrInvoiceNotOpen
	}

	// Monta payload para dedução de estoque
	deductItems := make([]DeductRequest, 0, len(invoice.Items))
	for _, item := range invoice.Items {
		deductItems = append(deductItems, DeductRequest{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	// Tenta deduzir estoque no serviço de inventário
	if err := s.inventoryClient.DeductStock(deductItems); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInventoryUnavailable, err)
	}

	// Estoque deduzido com sucesso → fecha a nota fiscal
	if err := s.repo.UpdateStatus(id, persistence.StatusClosed); err != nil {
		return nil, err
	}

	invoice.Status = persistence.StatusClosed
	return invoice, nil
}

// GetAllInvoices retorna todas as notas fiscais com seus itens.
func (s *InvoiceService) GetAllInvoices() ([]persistence.Invoice, error) {
	return s.repo.FindAll()
}
