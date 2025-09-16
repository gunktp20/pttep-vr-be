package permissions

import (
	"context"
	"pttep-vr-api/pkg/models"
	"pttep-vr-api/pkg/utils/pagination"
)

func (s *Service) Create(ctx context.Context, v models.Permission) (models.Permission, error) {
	v.IsActive = true
	v.IsRemove = false
	data, err := s.repository.InsertOnePermissions(ctx, v)
	if err != nil {
		return models.Permission{}, err
	}

	return data, nil
}

func (s *Service) Delete(ctx context.Context, id uint) error {
	_, err := s.repository.UpdateOnePermissionsIsRemove(ctx, models.Permission{
		ID:       id,
		IsRemove: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetByID(ctx context.Context, id uint) (models.Permission, error) {
	return s.repository.FindOnePermissions(ctx, id)
}

func (s *Service) Get(ctx context.Context, paginate *pagination.Pagination) ([]models.Permission, int64, error) {
	return s.repository.FindPermissions(ctx, paginate)
}

func (s *Service) UpdateIsActive(ctx context.Context, v models.Permission) error {
	_, err := s.repository.UpdateOnePermissionsIsActive(ctx, v)
	if err != nil {
		return err
	}
	return nil
}
