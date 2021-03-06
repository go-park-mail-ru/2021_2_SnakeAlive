package reviewRepository

import (
	"context"
	"errors"
	logs "snakealive/m/internal/logger"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/domain"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type reviewStorage struct {
	dataHolder *pgxpool.Pool
}

func NewReviewStorage(DB *pgxpool.Pool) domain.ReviewStorage {
	return &reviewStorage{dataHolder: DB}
}

func (u *reviewStorage) Add(value domain.Review, userId int) error {
	logger := logs.GetLogger()

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.AddReviewQuery,
		value.Title,
		value.Text,
		value.Rating,
		userId,
		value.PlaceId,
	)
	return err
}

func (u *reviewStorage) GetListByPlace(id int) (domain.Reviews, error) {
	logger := logs.GetLogger()

	reviews := make(domain.Reviews, 0)
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return reviews, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		cnst.GetReviewsByPlaceQuery,
		id)
	if err != nil {
		logger.Error("error while getting places from database")
		return reviews, err
	}

	var review domain.Review
	for rows.Next() {
		err = rows.Scan(&review.Id, &review.Title, &review.Text, &review.Rating, &review.UserId, &review.PlaceId)
		reviews = append(reviews, review)
	}
	if rows.Err() != nil {
		logger.Error("error while scanning places from database")
		return reviews, err
	}
	if len(reviews) == 0 {
		logger.Error("no reviews found")
		return reviews, errors.New("no reviews")
	}
	return reviews, err
}

func (u *reviewStorage) Get(id int) (domain.Review, error) {
	logger := logs.GetLogger()

	var review domain.Review
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return review, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		cnst.GetReviewByIdQuery,
		id,
	).Scan(&review.Id, &review.Title, &review.Text, &review.Rating, &review.UserId, &review.PlaceId)
	if err != nil {
		logger.Error("error while scanning places: ", zap.Error(err))
		return review, err
	}
	return review, err
}

func (u *reviewStorage) Delete(id int) error {
	logger := logs.GetLogger()

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.DeleteReviewQuery,
		id,
	)
	return err
}

func (u *reviewStorage) GetReviewAuthor(id int) int {
	logger := logs.GetLogger()

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return 0
	}
	defer conn.Release()

	var author int
	err = conn.QueryRow(context.Background(),
		cnst.GetReviewAuthorQuery,
		id,
	).Scan(&author)

	if err != nil {
		logger.Error("error while finding author of review: ", zap.Error(err))
		return 0
	}
	return author
}
