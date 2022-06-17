package services

import "github.com/linyuan-t/quick-search/backend/config"

type Services struct {
}

func NewServices(config *config.Config) *Services {
	return &Services{}
}
