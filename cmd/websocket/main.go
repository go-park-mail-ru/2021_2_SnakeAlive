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

	http.HandleFunc("/connect", corsMiddleware(delivery.Connect))
	http.HandleFunc("/", corsMiddleware(delivery.HandleRequest))

	err = http.ListenAndServe(cfg.Port, nil)
	if err != nil {
		log.Fatal("failed to start server")
		return
	}
}

func corsMiddleware(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://snakehastrip.ru") // set domain
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}
