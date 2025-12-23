package services

import (
	"ecommerce-go-api-gateway/models"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type productService struct {
	baseURL string
	client  *resty.Client
}

func NewProductService(baseURL string, client *resty.Client) ProductService {
	return &productService{baseURL: baseURL, client: client}
}

func (s *productService) GetProduct(id uint) (*models.Product, error) {
	resp, err := s.client.R().
		Get(fmt.Sprintf("%s/products/%d", s.baseURL, id))

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("product service error: %s", resp.String())
	}

	var product models.Product
	if err := json.Unmarshal(resp.Body(), &product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *productService) ListProducts() ([]models.Product, error) {
	resp, err := s.client.R().
		Get(s.baseURL + "/products")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("product service error: %s", resp.String())
	}

	var products []models.Product
	if err := json.Unmarshal(resp.Body(), &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) CreateProduct(req models.CreateProductRequest) (*models.Product, error) {
	resp, err := s.client.R().
		SetBody(req).
		Post(s.baseURL + "/products")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("product service error: %s", resp.String())
	}

	var product models.Product
	if err := json.Unmarshal(resp.Body(), &product); err != nil {
		return nil, err
	}
	return &product, nil
}
