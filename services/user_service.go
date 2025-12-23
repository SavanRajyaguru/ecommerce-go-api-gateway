package services

import (
	"ecommerce-go-api-gateway/models"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type userService struct {
	baseURL string
	client  *resty.Client
}

func NewUserService(baseURL string, client *resty.Client) UserService {
	return &userService{baseURL: baseURL, client: client}
}

func (s *userService) Login(req models.LoginRequest) (*models.LoginResponse, error) {
	resp, err := s.client.R().
		SetBody(req).
		Post(s.baseURL + "/login")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("user service error: %s", resp.String())
	}

	var loginResp models.LoginResponse
	if err := json.Unmarshal(resp.Body(), &loginResp); err != nil {
		return nil, err
	}
	return &loginResp, nil
}

func (s *userService) Register(req models.CreateUserRequest) (*models.User, error) {
	resp, err := s.client.R().
		SetBody(req).
		Post(s.baseURL + "/register")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("user service error: %s", resp.String())
	}

	var user models.User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	resp, err := s.client.R().
		Get(fmt.Sprintf("%s/users/%d", s.baseURL, id))

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("user service error: %s", resp.String())
	}

	var user models.User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, err
	}
	return &user, nil
}
