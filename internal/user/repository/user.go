package userRepository

import (
	"context"
	"fmt"
	"snakealive/m/pkg/domain"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type UserStorage struct {
	dataHolder *pgxpool.Pool
}

func NewUserStorage(DB *pgxpool.Pool) domain.UserStorage {
	return &UserStorage{dataHolder: DB}
}

func (u *UserStorage) Add(value domain.User) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`INSERT INTO Users ("name", "surname", "password", "email", "avatar") VALUES ($1, $2, $3, $4, $5)`,

		value.Name,
		value.Surname,
		value.Password,
		value.Email,
		value.Avatar,
	)
	return err
}

func (u *UserStorage) Delete(id int) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while deleting user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`DELETE FROM Users WHERE id = $1`,
		id,
	)
	return err
}

func (u *UserStorage) DeleteByEmail(user domain.User) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while deleting user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`DELETE FROM Users WHERE email = $1`,
		user.Email,
	)
	return err
}

func (u *UserStorage) GetByEmail(key string) (value domain.User, err error) {
	var user domain.User

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Error while getting user")
		return user, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT id, name, surname, password, email, avatar
		FROM Users WHERE email = $1`,
		key,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Avatar)

	return user, err
}

func (u *UserStorage) GetById(id int) (value domain.User, err error) {
	var user domain.User

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Error while adding user")
		return user, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		`SELECT id, name, surname, password, email, avatar
		FROM Users WHERE id = $1`,
		id,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Avatar)

	return user, err
}

func (u *UserStorage) Update(id int, value domain.User) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		fmt.Printf("Connection error while adding user ", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		`UPDATE Users SET "name" = $1, "surname" = $2, "email" = $3, "password" = $4, "avatar" = $5 WHERE id = $6`,
		value.Name,
		value.Surname,
		value.Email,
		value.Password,
		value.Avatar,
		id,
	)
	return err
}
