package userRepository

import (
	"context"
	"fmt"
	"snakealive/m/domain"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type userStorage struct {
	dataHolder *pgxpool.Pool
}

func NewUserStorage(DB *pgxpool.Pool) domain.UserStorage {
	return &userStorage{dataHolder: DB}
}

func (u *userStorage) Add(value domain.User) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`INSERT INTO Users ("name", "surname", "password", "email") VALUES ($1, $2, $3, $4)`,
		value.Name,
		value.Surname,
		value.Password,
		value.Email,
	)
	return err
}

func (u *userStorage) Delete(id int) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`DELETE FROM Users WHERE id = $1`,
		id,
	)
	return err
}

func (u *userStorage) Get(key string) (value domain.User, err error) {
	var user domain.User

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Error while adding user")
		return user, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT id, name, surname, password, email
		FROM Users WHERE email = $1`,
		key,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email)

	return user, err
}

func (u *userStorage) Update(id int, value domain.User) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`UPDATE Users SET "name" = $1, "surname" = $2, "email" = $3, "password" = $4 WHERE id = $5`,
		value.Name,
		value.Surname,
		value.Email,
		value.Password,
		id,
	)
	return err
}
