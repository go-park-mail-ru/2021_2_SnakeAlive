package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"snakealive/m/internal/websocket/delivery"
	"snakealive/m/internal/websocket/repository"
	"snakealive/m/internal/websocket/usecase"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	url := "postgres://tripadvisor:12345@localhost:5432/tripadvisor"
	dbpool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	delivery := delivery.NewWebSocketDelivery(usecase.NewWebsocketUseCase(repository.NewWebsocketRepository(dbpool)))

	http.HandleFunc("/connect", delivery.Connect)
	http.HandleFunc("/", delivery.HandleRequest)

	http.ListenAndServe(":5050", nil)
}
