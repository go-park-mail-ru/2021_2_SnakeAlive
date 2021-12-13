package usecase

import (
	"context"
	"encoding/json"
	"snakealive/m/internal/models"
	trip "snakealive/m/internal/services/trip/models"
	"snakealive/m/internal/websocket/repository"

	"github.com/gorilla/websocket"
)

type WebsocketUseCase interface {
	Update(ctx context.Context, tripId int) (*trip.Trip, error)
	Connect(userId int, conn *websocket.Conn)
	SendResponce(users []int, responce models.TripResponce) error
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

func (u websocketUseCase) SendResponce(users []int, responce models.TripResponce) error {
	result, err := json.Marshal(responce)
	if err != nil {
		return err
	}

	conns := u.websocketRepository.GetConnections(users)

	for _, conn := range conns {
		conn.WriteJSON(result)
	}
	return nil
}
