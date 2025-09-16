package permissions

import (
	"pttep-vr-api/pkg/config"
	pttep_vr_db "pttep-vr-api/pkg/repository/pttep-vr-db"
)

type Service struct {
	config     *config.Config
	repository pttep_vr_db.Interface
}

func New(config *config.Config, repository pttep_vr_db.Interface) *Service {
	return &Service{
		config:     config,
		repository: repository,
	}
}
