package main

import (
	"log"
	"net/http"
	"snakealive/m/internal/websocket/config"
	"snakealive/m/internal/websocket/delivery"
	"snakealive/m/internal/websocket/repository"
	"snakealive/m/internal/websocket/usecase"
	"snakealive/m/pkg/helpers"
)

func main() {

	var cfg config.Config
	if err := cfg.Setup(); err != nil {
		log.Fatal("failed to setup cfg: ", err)
		return
	}

	dbpool, err := helpers.CreatePGXPool(cfg.Ctx, cfg.DBUrl, cfg.Logger)
	if err != nil {
		log.Fatal("failed to connect to db: ", err)
		return
	}
	defer dbpool.Close()

	delivery := delivery.NewWebSocketDelivery(usecase.NewWebsocketUseCase(repository.NewWebsocketRepository(dbpool)))

	http.HandleFunc("/connect", delivery.Connect)
	http.HandleFunc("/", delivery.HandleRequest)

	err = http.ListenAndServe(cfg.Port, nil)
	if err != nil {
		log.Fatal("failed to start server")
		return
	}
}
