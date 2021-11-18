package reviewRepository

import (
	"context"
	"errors"
	"snakealive/m/internal/domain"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type reviewStorage struct {
	dataHolder *pgxpool.Pool
}

func NewReviewStorage(DB *pgxpool.Pool) domain.ReviewStorage {
	return &reviewStorage{dataHolder: DB}
}

const AddReviewQuery = `INSERT INTO public.reviews (title, text, rating, user_id, place_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

const GetReviewByIdQuery = `SELECT id, title, text, rating, user_id, place_id FROM Reviews WHERE id = $1`

const DeleteReviewQuery = `DELETE FROM Reviews WHERE id = $1`

const GetReviewAuthorQuery = `SELECT user_id FROM Reviews WHERE id = $1`

func (u *reviewStorage) Add(value domain.Review, userId int) (int, error) {
	insertedId := 0
	conn, err := u.dataHolder.Acquire(context.Background())
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

func (u *reviewStorage) GetListByPlace(id int, limit int, skip int) (domain.ReviewsNoPlace, error) {
	reviews := make(domain.ReviewsNoPlace, 0)
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return reviews, err
	}
	defer conn.Release()
	const GetReviewsByPlaceQuery = `SELECT id, title, text, rating, user_id FROM Reviews WHERE place_id = $1 LIMIT $2 OFFSET $3`
	rows, err := conn.Query(context.Background(),
		GetReviewsByPlaceQuery,
		id, limit, skip)
	if err != nil {
		return reviews, err
	}

	var review domain.ReviewNoPlace
	for rows.Next() {
		err = rows.Scan(&review.Id, &review.Title, &review.Text, &review.Rating, &review.UserId)
		reviews = append(reviews, review)
	}
	if rows.Err() != nil {
		return reviews, err
	}
	if len(reviews) == 0 {
		return reviews, errors.New("no reviews")
	}
	return reviews, err
}

func (u *reviewStorage) Get(id int) (domain.Review, error) {
	var review domain.Review
	conn, err := u.dataHolder.Acquire(context.Background())
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

func (u *reviewStorage) Delete(id int) error {
	conn, err := u.dataHolder.Acquire(context.Background())
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

func (u *reviewStorage) GetReviewAuthor(id int) int {
	conn, err := u.dataHolder.Acquire(context.Background())
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
