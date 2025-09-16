package ping

import (
	"pttep-vr-api/pkg/config"
)

type Service struct {
	config *config.Config
}

func New(config *config.Config) *Service {
	return &Service{
		config: config,
	}
}

func (s *Service) Ping() (string, error) {
	return "Pong", nil
}

type Interface interface {
	Ping() (string, error)
}
