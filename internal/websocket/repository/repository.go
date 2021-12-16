package repository

import (
	"context"

	"snakealive/m/internal/services/trip/models"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v4/pgxpool"
)

type WebsocketRepository interface {
	GetTripById(ctx context.Context, id int) (value *models.Trip, err error)
	GetTripAuthors(ctx context.Context, id int) ([]int, error)
	AddConnection(userId int, conn *websocket.Conn)
	GetConnections(users []int) []*websocket.Conn
	ValidateSession(ctx context.Context, hash string) (int, error)
}

type websocketRepository struct {
	dataHolder *pgxpool.Pool
}

func NewWebsocketRepository(DB *pgxpool.Pool) WebsocketRepository {
	return &websocketRepository{dataHolder: DB}
}

func (r *websocketRepository) GetTripById(ctx context.Context, id int) (value *models.Trip, err error) {
	var trip models.Trip

	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return &trip, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		getTripQuery,
		id,
	).Scan(&trip.Id, &trip.Title, &trip.Description)
	if err != nil {
		return &trip, err
	}

	rows, err := conn.Query(context.Background(),
		getPlaceForTripQuery,
		id)
	if err != nil {
		return &trip, err
	}
	defer rows.Close()

	var place models.Place
	for rows.Next() {
		_ = rows.Scan(&place.Id, &place.Name, &place.Tags, &place.Description, &place.Rating, &place.Country, &place.Photos, &place.Day)
		trip.Sights = append(trip.Sights, place)
	}

	rows, err = conn.Query(context.Background(),
		getAlbumsByTripQuery,
		id)
	if err != nil {
		return &trip, err
	}

	var album models.Album
	for rows.Next() {
		_ = rows.Scan(&album.Id, &album.Title, &album.Description, &album.Photos)
		trip.Albums = append(trip.Albums, album)
	}

	users, err := r.GetTripAuthors(ctx, trip.Id)
	if err != nil {
		return &trip, err
	}

	trip.Users = users

	return &trip, err
}

func (r *websocketRepository) GetTripAuthors(ctx context.Context, id int) ([]int, error) {
	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return []int{}, err
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(),
		getTripUsersQuery,
		id)
	if err != nil {
		return []int{}, err
	}
	defer rows.Close()

	var ids []int
	var userId int
	for rows.Next() {
		_ = rows.Scan(&userId)
		ids = append(ids, userId)
	}
	return ids, err
}

func (r *websocketRepository) AddConnection(userId int, conn *websocket.Conn) {
	conns.mu.Lock()
	conns.storage[userId] = conn
	conns.mu.Unlock()
}

func (r *websocketRepository) GetConnections(users []int) []*websocket.Conn {
	sockets := make([]*websocket.Conn, 0)
	for _, user := range users {
		conns.mu.Lock()
		sockets = append(sockets, conns.storage[user])
		conns.mu.Unlock()
	}
	return sockets
}

func (r *websocketRepository) ValidateSession(ctx context.Context, hash string) (int, error) {
	var userId int

	conn, err := r.dataHolder.Acquire(context.Background())
	if err != nil {
		return userId, err
	}
	defer conn.Release()

	err = conn.QueryRow(context.Background(),
		validateUserSession,
		hash,
	).Scan(&userId)
	if err != nil {
		return userId, err
	}

	return userId, nil
}
