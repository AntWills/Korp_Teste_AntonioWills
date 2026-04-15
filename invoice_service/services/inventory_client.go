package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// DeductRequest é o payload enviado ao inventory_service para dedução de estoque.
type DeductRequest struct {
	ProductCode string `json:"product_code"`
	Quantity    int    `json:"quantity"`
}

// InventoryClient define o contrato de comunicação com o serviço de estoque.
type InventoryClient interface {
	DeductStock(items []DeductRequest) error
}

type inventoryClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewInventoryClient() InventoryClient {
	baseURL := os.Getenv("INVENTORY_SERVICE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return &inventoryClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *inventoryClient) DeductStock(items []DeductRequest) error {
	payload, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("falha ao serializar requisição de dedução: %w", err)
	}

	url := fmt.Sprintf("%s/api/inventory/deduct", c.baseURL)
	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		// Erro de rede: serviço fora do ar, timeout, recusa de conexão etc.
		return fmt.Errorf("serviço de estoque inacessível (%s): %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("serviço de estoque retornou erro HTTP %d", resp.StatusCode)
	}

	return nil
}
