package repository

import (
	"context"

	"snakealive/m/internal/services/review/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ReviewRepository interface {
	Add(ctx context.Context, value *models.Review, userId int) (int, error)
	Get(ctx context.Context, id int) (*models.Review, error)
	GetListByPlace(ctx context.Context, id int, limit int, skip int) (*[]models.Review, error)
	Delete(ctx context.Context, id int) error
	GetReviewAuthor(ctx context.Context, id int) int
	GetReviewsAmount(ctx context.Context, id int) (int, error)
}

type reviewRepository struct {
	dataHolder *pgxpool.Pool
}

func NewReviewRepository(DB *pgxpool.Pool) ReviewRepository {
	return &reviewRepository{dataHolder: DB}
}

func (r *reviewRepository) Add(ctx context.Context, value *models.Review, userId int) (int, error) {
	insertedId := 0
	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return insertedId, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		AddReviewQuery,
		value.Title,
		value.Text,
		value.Rating,
		userId,
		value.PlaceId,
	).Scan(&insertedId)
	if err != nil {
		return insertedId, err
	}

	return insertedId, err
}

func (r *reviewRepository) Get(ctx context.Context, id int) (*models.Review, error) {
	review := new(models.Review)
	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return review, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		GetReviewByIdQuery,
		id,
	).Scan(&review.Id, &review.Title, &review.Text, &review.Rating, &review.UserId, &review.PlaceId)
	if err != nil {
		return review, err
	}

	return review, err
}
func (r *reviewRepository) GetListByPlace(ctx context.Context, id int, limit int, skip int) (*[]models.Review, error) {
	var reviews []models.Review

	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return &reviews, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		GetReviewsByPlaceQuery,
		id, limit, skip)
	if err != nil {
		return &reviews, err
	}
	defer rows.Close()

	var review models.Review
	for rows.Next() {
		err = rows.Scan(&review.Id, &review.Title, &review.Text, &review.Rating, &review.UserId)
		reviews = append(reviews, review)
	}
	if rows.Err() != nil {
		return &reviews, err
	}

	return &reviews, err
}

func (r *reviewRepository) Delete(ctx context.Context, id int) error {
	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		DeleteReviewQuery,
		id,
	)
	return err
}

func (r *reviewRepository) GetReviewAuthor(ctx context.Context, id int) int {
	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return 0
	}
	defer conn.Release()

	var author int
	err = conn.QueryRow(context.Background(),
		GetReviewAuthorQuery,
		id,
	).Scan(&author)

	if err != nil {
		return 0
	}
	return author
}

func (r *reviewRepository) GetReviewsAmount(ctx context.Context, id int) (int, error) {
	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	var amount int
	err = conn.QueryRow(context.Background(),
		GetReviewAmount,
		id,
	).Scan(&amount)

	if err != nil {
		return 0, err
	}
	return amount, nil
}
