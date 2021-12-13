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
	GetSightByTag(ctx context.Context, tag int64) ([]models.Sight, error)
	SearchSights(ctx context.Context, search string, skip, limit int64) ([]models.Sight, error)

	GetTags(ctx context.Context) ([]models.Tag, error)
}

type sightUsecase struct {
	repo repository.SightRepository
}

func (s *sightUsecase) GetTags(ctx context.Context) ([]models.Tag, error) {
	return s.repo.GetTags(ctx)
}

func (s *sightUsecase) SearchSights(ctx context.Context, search string, skip, limit int64) (sights []models.Sight, err error) {
	return s.repo.SearchSights(ctx, search, skip, limit)
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

func (s *sightUsecase) GetSightByTag(ctx context.Context, tag int64) ([]models.Sight, error) {
	return s.repo.GetSightByTag(ctx, tag)
}

func NewSightUsecase(repo repository.SightRepository) SightUsecase {
	return &sightUsecase{repo: repo}
}
