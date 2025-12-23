package services

import (
	"ecommerce-go-api-gateway/models"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type orderService struct {
	baseURL string
	client  *resty.Client
}

func NewOrderService(baseURL string, client *resty.Client) OrderService {
	return &orderService{baseURL: baseURL, client: client}
}

func (s *orderService) CreateOrder(req models.CreateOrderRequest) (*models.Order, error) {
	resp, err := s.client.R().
		SetBody(req).
		Post(s.baseURL + "/orders")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("order service error: %s", resp.String())
	}

	var order models.Order
	if err := json.Unmarshal(resp.Body(), &order); err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *orderService) GetOrder(id uint) (*models.Order, error) {
	resp, err := s.client.R().
		Get(fmt.Sprintf("%s/orders/%d", s.baseURL, id))

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("order service error: %s", resp.String())
	}

	var order models.Order
	if err := json.Unmarshal(resp.Body(), &order); err != nil {
		return nil, err
	}
	return &order, nil
}
