package userRepository

import (
	"context"
	"snakealive/m/internal/domain"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type UserStorage struct {
	dataHolder *pgxpool.Pool
}

func NewUserStorage(DB *pgxpool.Pool) domain.UserStorage {
	return &UserStorage{dataHolder: DB}
}

func (u *UserStorage) Add(value domain.User) error {
	logger := logs.GetLogger()
	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.AddUserQuery,
		value.Name,
		value.Surname,
		value.Password,
		value.Email,
		value.Avatar,
	)
	return err
}

func (u *UserStorage) Delete(id int) error {
	logger := logs.GetLogger()

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.DeleteUserByIdQuery,
		id,
	)
	return err
}

func (u *UserStorage) DeleteByEmail(user domain.User) error {
	logger := logs.GetLogger()

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.DeleteUserByEmailQuery,
		user.Email,
	)
	return err
}

func (u *UserStorage) GetByEmail(key string) (value domain.User, err error) {
	logger := logs.GetLogger()
	var user domain.User

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return domain.User{}, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		cnst.GetUserByEmailQuery,
		key,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Avatar)

	return user, err
}

func (u *UserStorage) GetById(id int) (value domain.User, err error) {
	logger := logs.GetLogger()
	var user domain.User

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return user, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		cnst.GetUserByIdQuery,
		id,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Avatar)

	return user, err
}

func (u *UserStorage) GetPublicById(id int) (value domain.PublicUser, err error) {
	logger := logs.GetLogger()
	var user domain.PublicUser

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return user, err
	}
	defer conn.Release()
	const GetPublicUserByIdQuery = `SELECT id, name, surname, avatar FROM Users WHERE id = $1`
	err = conn.QueryRow(context.Background(),
		GetPublicUserByIdQuery,
		id,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Avatar)

	return user, err
}

func (u *UserStorage) Update(id int, value domain.User) error {
	logger := logs.GetLogger()

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.UpdateUserQuery,
		value.Name,
		value.Surname,
		value.Email,
		value.Password,
		value.Avatar,
		id,
	)
	return err
}

func (u *UserStorage) AddAvatar(id int, avatar string) error {
	logger := logs.GetLogger()

	conn, err := u.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.AddAvatarQuery,
		avatar,
		id,
	)
	return err
}
