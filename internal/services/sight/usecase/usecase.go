package usecase

import (
	"context"

	"snakealive/m/internal/services/sight/models"
	"snakealive/m/internal/services/sight/repository"
)

type SightUsecase interface {
	GetSightsByCountry(ctx context.Context, country string) ([]models.Sight, error)
	GetSightsByIDs(ctx context.Context, ids []int64) ([]models.Sight, error)
	GetSightByID(ctx context.Context, id int) (models.Sight, error)
}

type sightUsecase struct {
	repo repository.SightRepository
}

func (s *sightUsecase) GetSightsByIDs(ctx context.Context, ids []int64) ([]models.Sight, error) {
	return s.repo.GetSightByIDs(ctx, ids)
}

func (s *sightUsecase) GetSightsByCountry(ctx context.Context, country string) ([]models.Sight, error) {
	return s.repo.GetSightsByCountry(ctx, country)
}

func (s *sightUsecase) GetSightByID(ctx context.Context, id int) (models.Sight, error) {
	return s.repo.GetSightByID(ctx, id)
}

func NewSightUsecase(repo repository.SightRepository) SightUsecase {
	return &sightUsecase{repo: repo}
}
