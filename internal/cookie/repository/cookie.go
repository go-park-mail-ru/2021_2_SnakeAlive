package cookieRepository

import (
	"context"
	logs "snakealive/m/internal/logger"
	"snakealive/m/pkg/domain"

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
		`INSERT INTO Cookies ("hash", "user_id") VALUES ($1, $2)`,
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
		`SELECT U.id, U.name, U.surname, U.password, U.email
		FROM Users AS U
		JOIN Cookies AS C ON U.id = C.user_id
		WHERE C.hash = $1`,
		value,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email)

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
		`DELETE FROM Cookies WHERE hash = $1`,
		value,
	)
	return err
}
