package reviewRepository

import (
	"context"
	"errors"
	"fmt"
	"snakealive/m/pkg/domain"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type reviewStorage struct {
	dataHolder *pgxpool.Pool
}

func NewReviewStorage(DB *pgxpool.Pool) domain.ReviewStorage {
	return &reviewStorage{dataHolder: DB}
}

func (u *reviewStorage) Add(value domain.Review, userId int) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Print("Connection error while adding review ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`INSERT INTO public.reviews (title, text, rating, user_id, place_id) VALUES ($1, $2, $3, $4, $5)`,
		value.Title,
		value.Text,
		value.Rating,
		userId,
		value.PlaceId,
	)
	return err
}

func (u *reviewStorage) GetListByPlace(id int) (domain.Reviews, error) {
	reviews := make(domain.Reviews, 0)
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Print(err)
		return reviews, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		`SELECT id, title, text, rating, user_id, place_id FROM Reviews
		Where place_id = $1`,
		id)
	if err != nil {
		fmt.Printf("Error while getting places")
		return reviews, err
	}

	var review domain.Review
	for rows.Next() {
		err = rows.Scan(&review.Id, &review.Title, &review.Text, &review.Rating, &review.UserId, &review.PlaceId)
		reviews = append(reviews, review)
	}
	if rows.Err() != nil {
		fmt.Printf("Error while scanning places")
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
		fmt.Print("Error while getting user", err)
		return review, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT id, title, text, rating, user_id, place_id FROM Reviews
		Where id = $1`,
		id,
	).Scan(&review.Id, &review.Title, &review.Text, &review.Rating, &review.UserId, &review.PlaceId)
	if err != nil {
		fmt.Print("Error while scanning places", err)
		return review, err
	}
	return review, err
}

func (u *reviewStorage) Delete(id int) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Print("Connection error while deleting user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`DELETE FROM Reviews WHERE id = $1`,
		id,
	)
	return err
}

func (u *reviewStorage) GetReviewAuthor(id int) int {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Print("Connection error while deleting user ", err)
		return 0
	}
	defer conn.Release()

	var author int
	err = conn.QueryRow(context.Background(),
		`SELECT user_id
	FROM Reviews WHERE id = $1`,
		id,
	).Scan(&author)

	if err != nil {
		return 0
	}
	return author
}
