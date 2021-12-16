package usecase

import (
	"context"

	"snakealive/m/internal/models"
	trip "snakealive/m/internal/services/trip/models"
	"snakealive/m/internal/websocket/repository"

	"github.com/gorilla/websocket"
)

type WebsocketUseCase interface {
	Update(ctx context.Context, tripId int) (*trip.Trip, error)
	Connect(userId int, conn *websocket.Conn)
	SendUpdateResponce(users []int, responce models.TripResponce) error
	SendDeleteResponce(users []int, responce models.TripRequest) error
	ValidateSession(ctx context.Context, hash string) (int, error)
}

type websocketUseCase struct {
	websocketRepository repository.WebsocketRepository
}

func NewWebsocketUseCase(websocketRepository repository.WebsocketRepository) WebsocketUseCase {
	return &websocketUseCase{websocketRepository: websocketRepository}
}

func (u websocketUseCase) Update(ctx context.Context, tripId int) (*trip.Trip, error) {
	return u.websocketRepository.GetTripById(ctx, tripId)
}

func (u websocketUseCase) Connect(userId int, conn *websocket.Conn) {
	u.websocketRepository.AddConnection(userId, conn)
}

func (u websocketUseCase) SendUpdateResponce(users []int, responce models.TripResponce) error {
	conns := u.websocketRepository.GetConnections(users)

	for _, conn := range conns {
		_ = conn.WriteJSON(responce)
	}
	return nil
}

func (u websocketUseCase) SendDeleteResponce(users []int, responce models.TripRequest) error {
	conns := u.websocketRepository.GetConnections(users)

	for _, conn := range conns {
		_ = conn.WriteJSON(responce)
	}
	return nil
}

func (u websocketUseCase) ValidateSession(ctx context.Context, hash string) (int, error) {
	return u.websocketRepository.ValidateSession(ctx, hash)
}
