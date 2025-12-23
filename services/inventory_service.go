package services

import (
	"ecommerce-go-api-gateway/models"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type inventoryService struct {
	baseURL string
	client  *resty.Client
}

func NewInventoryService(baseURL string, client *resty.Client) InventoryService {
	return &inventoryService{baseURL: baseURL, client: client}
}

func (s *inventoryService) UpdateStock(req models.UpdateInventoryRequest) error {
	resp, err := s.client.R().
		SetBody(req).
		Post(s.baseURL + "/inventory/stock")

	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("inventory service error: %s", resp.String())
	}
	return nil
}
