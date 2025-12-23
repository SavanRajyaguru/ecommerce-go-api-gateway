package services

import (
	"ecommerce-go-api-gateway/models"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type paymentService struct {
	baseURL string
	client  *resty.Client
}

func NewPaymentService(baseURL string, client *resty.Client) PaymentService {
	return &paymentService{baseURL: baseURL, client: client}
}

func (s *paymentService) ProcessPayment(req models.CreatePaymentRequest) (*models.Payment, error) {
	resp, err := s.client.R().
		SetBody(req).
		Post(s.baseURL + "/payments")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("payment service error: %s", resp.String())
	}

	var payment models.Payment
	if err := json.Unmarshal(resp.Body(), &payment); err != nil {
		return nil, err
	}
	return &payment, nil
}
