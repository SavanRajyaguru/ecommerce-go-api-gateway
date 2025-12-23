package services

import (
	"ecommerce-go-api-gateway/models"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type notificationService struct {
	baseURL string
	client  *resty.Client
}

func NewNotificationService(baseURL string, client *resty.Client) NotificationService {
	return &notificationService{baseURL: baseURL, client: client}
}

func (s *notificationService) SendNotification(req models.SendNotificationRequest) error {
	resp, err := s.client.R().
		SetBody(req).
		Post(s.baseURL + "/notifications")

	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("notification service error: %s", resp.String())
	}
	return nil
}
