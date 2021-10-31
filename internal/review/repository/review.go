package reviewRepository

import (
	"context"
	"fmt"
	pr "snakealive/m/internal/place/repository"
	"snakealive/m/pkg/domain"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type reviewStorage struct {
	dataHolder *pgxpool.Pool
}

func NewRewiewStorage(DB *pgxpool.Pool) domain.ReviewStorage {
	return &reviewStorage{dataHolder: DB}
}

func (u *reviewStorage) Add(value domain.Review) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`INSERT INTO reviews (id, title, text, rating, user_id, created_at) VALUES ($1, $2, $3, $4)`,
		value.Title,
		value.Text,
		value.Rating,
		value.User_id,
	)
	return err
}

func (u *reviewStorage) Get(key string) (domain.Reviews, error) {
	var reviews domain.Reviews
	var place domain.Place

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Error while getting user")
		return reviews, err
	}
	defer conn.Release()
	place, _ = pr.NewPlaceStorage().Get(key)

	rows, err := conn.Query(context.Background(),
		`SELECT id, name, surname, password, email
		FROM Users WHERE email = $1`,
		place.Name,
	)
	i := 0
	for rows.Next() {
		err = rows.Scan(&reviews[i].Id, &reviews[i].Title, &reviews[i].Text, &reviews[i].Rating, &reviews[i].User_id, &reviews[i].Place_id, &reviews[i].Created_at)
		i += 1
	}
	return reviews, err
}

// func (u *reviewStorage) Delete(id int) error {
// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		fmt.Printf("Connection error while deleting user ", err)
// 		return err
// 	}
// 	defer conn.Release()

// 	_, err = conn.Exec(context.Background(),
// 		`DELETE FROM Users WHERE id = $1`,
// 		id,
// 	)
// 	return err
// }

// func (u *reviewStorage) GetByEmail(key string) (value domain.Review, err error) {
// 	var user domain.User

// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		fmt.Printf("Error while getting user")
// 		return user, err
// 	}
// 	defer conn.Release()

// 	err = conn.QueryRow(context.Background(),
// 		`SELECT id, name, surname, password, email
// 		FROM Users WHERE email = $1`,
// 		key,
// 	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email)

// 	return user, err
// }

// func (u *reviewStorage) GetById(id int) (value domain.Review, err error) {
// 	var user domain.User

// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		fmt.Printf("Error while adding user")
// 		return user, err
// 	}
// 	defer conn.Release()

// 	err = conn.QueryRow(context.Background(),
// 		`SELECT id, name, surname, password, email
// 		FROM Users WHERE id = $1`,
// 		id,
// 	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email)

// 	return user, err
// }

// func (u *reviewStorage) Update(id int, value domain.Review) error {
// 	conn, err := u.dataHolder.Acquire(context.Background())
// 	if err != nil {
// 		fmt.Printf("Connection error while adding user ", err)
// 		return err
// 	}
// 	defer conn.Release()

// 	_, err = conn.Exec(context.Background(),
// 		`UPDATE Users SET "name" = $1, "surname" = $2, "email" = $3, "password" = $4 WHERE id = $5`,
// 		value.Name,
// 		value.Surname,
// 		value.Email,
// 		value.Password,
// 		id,
// 	)
// 	return err
// }
