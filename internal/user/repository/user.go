package userRepository

import (
	"context"
	"snakealive/m/internal/domain"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type UserStorage struct {
	dataHolder *pgxpool.Pool
}

func NewUserStorage(DB *pgxpool.Pool) domain.UserStorage {
	return &UserStorage{dataHolder: DB}
}

const AddUserQuery = `INSERT INTO Users ("name", "surname", "password", "email", "description") VALUES ($1, $2, $3, $4, $5)`

const DeleteUserByIdQuery = `DELETE FROM Users WHERE id = $1`

const DeleteUserByEmailQuery = `DELETE FROM Users WHERE email = $1`

const GetUserByEmailQuery = `SELECT id, name, surname, password, email, avatar, description FROM Users WHERE email = $1`

const GetUserByIdQuery = `SELECT id, name, surname, password, email, avatar, description FROM Users WHERE id = $1`

const UpdateUserQuery = `UPDATE Users SET "name" = $1, "surname" = $2, "email" = $3, "password" = $4,
	"description" = $5 WHERE id = $6`

const AddAvatarQuery = `UPDATE Users SET "avatar" = $1 WHERE id = $2`

func (u *UserStorage) Add(value domain.User) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		AddUserQuery,
		value.Name,
		value.Surname,
		value.Password,
		value.Email,
		value.Description,
	)
	return err
}

func (u *UserStorage) Delete(id int) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		DeleteUserByIdQuery,
		id,
	)
	return err
}

func (u *UserStorage) DeleteByEmail(user domain.User) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		DeleteUserByEmailQuery,
		user.Email,
	)
	return err
}

func (u *UserStorage) GetByEmail(key string) (value domain.User, err error) {
	var user domain.User

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return domain.User{}, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		GetUserByEmailQuery,
		key,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Avatar, &user.Description)

	return user, err
}

func (u *UserStorage) GetById(id int) (value domain.User, err error) {
	var user domain.User

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return user, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		GetUserByIdQuery,
		id,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Avatar, &user.Description)

	return user, err
}

func (u *UserStorage) GetPublicById(id int) (value domain.PublicUser, err error) {
	var user domain.PublicUser

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return user, err
	}
	defer conn.Release()
	const GetPublicUserByIdQuery = `SELECT id, name, surname, avatar, description FROM Users WHERE id = $1`
	err = conn.QueryRow(context.Background(),
		GetPublicUserByIdQuery,
		id,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Avatar, &user.Description)

	return user, err
}

func (u *UserStorage) Update(id int, value domain.User) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		UpdateUserQuery,
		value.Name,
		value.Surname,
		value.Email,
		value.Password,
		value.Description,
		id,
	)
	return err
}

func (u *UserStorage) AddAvatar(id int, avatar string) error {
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		AddAvatarQuery,
		avatar,
		id,
	)
	return err
}
