package cookieRepository

import (
	"context"
	"snakealive/m/internal/domain"
	cnst "snakealive/m/pkg/constants"
	logs "snakealive/m/pkg/logger"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type cookieStorage struct {
	dataHolder *pgxpool.Pool
}

func NewCookieStorage(DB *pgxpool.Pool) domain.CookieStorage {
	return &cookieStorage{dataHolder: DB}
}

func (c *cookieStorage) Add(key string, userId int) error {
	logger := logs.GetLogger()

	conn, err := c.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.AddCookieQuery,
		key,
		userId,
	)
	return err
}

func (c *cookieStorage) Get(value string) (user domain.User, err error) {
	logger := logs.GetLogger()

	conn, err := c.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return user, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		cnst.GetCookieQuery,
		value,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Description, &user.Avatar)

	return user, err

}
func (c *cookieStorage) Delete(value string) error {
	logger := logs.GetLogger()

	conn, err := c.dataHolder.Acquire(context.Background())
	if err != nil {
		logger.Error("error while aquiring connection")
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		cnst.DeleteCookieQuery,
		value,
	)
	return err
}
