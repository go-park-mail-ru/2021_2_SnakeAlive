package cookieRepository

import (
	"context"
	"snakealive/m/internal/domain"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
)

type cookieStorage struct {
	dataHolder *pgxpool.Pool
}

func NewCookieStorage(DB *pgxpool.Pool) domain.CookieStorage {
	return &cookieStorage{dataHolder: DB}
}

const AddCookieQuery = `INSERT INTO Cookies ("hash", "user_id") VALUES ($1, $2)`

const GetCookieQuery = `SELECT U.id, U.name, U.surname, U.password, U.email, U.description, U.avatar
						FROM Users AS U JOIN Cookies AS C ON U.id = C.user_id
						WHERE C.hash = $1`

const DeleteCookieQuery = `DELETE FROM Cookies WHERE hash = $1`

func (c *cookieStorage) Add(key string, userId int) error {
	conn, err := c.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		AddCookieQuery,
		key,
		userId,
	)
	return err
}

func (c *cookieStorage) Get(value string) (user domain.User, err error) {

	conn, err := c.dataHolder.Acquire(context.Background())
	if err != nil {
		return user, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		GetCookieQuery,
		value,
	).Scan(&user.Id, &user.Name, &user.Surname, &user.Password, &user.Email, &user.Description, &user.Avatar)

	return user, err

}
func (c *cookieStorage) Delete(value string) error {
	conn, err := c.dataHolder.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(),
		DeleteCookieQuery,
		value,
	)
	return err
}
