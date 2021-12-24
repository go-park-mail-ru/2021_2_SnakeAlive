package usecase

import (
	"context"
	"math/rand"

	"snakealive/m/internal/services/sight/models"
	"snakealive/m/internal/services/sight/repository"
	"snakealive/m/pkg/constants"
)

type SightUsecase interface {
	GetSightsByCountry(ctx context.Context, country string) ([]models.Sight, error)
	GetSightsByIDs(ctx context.Context, ids []int64) ([]models.Sight, error)
	GetSightByID(ctx context.Context, id int) (models.Sight, error)
	GetSightByTag(ctx context.Context, tag int64) ([]models.Sight, error)
	SearchSights(ctx context.Context, req *models.SightsSearch) ([]models.Sight, error)

	GetTags(ctx context.Context) ([]models.Tag, error)
	GetRandomTags(ctx context.Context) ([]models.Tag, error)
}

type sightUsecase struct {
	repo repository.SightRepository
}

func (s *sightUsecase) GetTags(ctx context.Context) ([]models.Tag, error) {
	return s.repo.GetTags(ctx)
}

func (s *sightUsecase) GetRandomTags(ctx context.Context) ([]models.Tag, error) {
	tags, err := s.GetTags(ctx)
	if err != nil {
		return []models.Tag{}, err
	}
	if len(tags) <= constants.DefaultTagCount {
		return tags, err
	}

	var result = make([]models.Tag, constants.DefaultTagCount)
	for i, val := range rand.Perm(len(tags))[:constants.DefaultTagCount] {
		result[i] = tags[val]
	}

	return result, nil
}

func (s *sightUsecase) SearchSights(ctx context.Context, req *models.SightsSearch) (sights []models.Sight, err error) {
	return s.repo.SearchSights(ctx, req)
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
