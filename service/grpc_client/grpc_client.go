package user_serviceClient

import (
	"user_service/config"
)

// user_serviceClientI ...
type user_serviceClientI interface {
}

// user_serviceClient ...
type user_serviceClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*user_serviceClient, error) {
	return &user_serviceClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
