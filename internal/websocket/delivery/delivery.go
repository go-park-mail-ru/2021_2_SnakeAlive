package delivery

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"snakealive/m/internal/models"
	"snakealive/m/internal/websocket/usecase"
	cnst "snakealive/m/pkg/constants"

	"github.com/gorilla/websocket"
)

type WebsocketDelivery interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
	Connect(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request, request models.UsersTripRequest)
	Delete(w http.ResponseWriter, request models.UsersTripRequest)
}

type websocketDelivery struct {
	websocketUsecase usecase.WebsocketUseCase
}

func NewWebSocketDelivery(websocketUsecase usecase.WebsocketUseCase) WebsocketDelivery {
	return &websocketDelivery{
		websocketUsecase: websocketUsecase,
	}
}

func (d *websocketDelivery) Connect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got in Connect  ", r.URL.String())
	cookie, err := r.Cookie(cnst.CookieName)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("{}"))
		return
	}

	userId, err := d.websocketUsecase.ValidateSession(r.Context(), cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("{}"))
		return
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("error while connecting: %s", err)
		_, _ = w.Write([]byte("{}"))
		return
	}

	d.websocketUsecase.Connect(userId, conn)
}

func (d *websocketDelivery) HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got in Handle request  ", r.URL.String())
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)

	decoder := json.NewDecoder(r.Body)

	request := new(models.UsersTripRequest)
	err := decoder.Decode(request)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		_, _ = w.Write([]byte("{}"))
		return
	}

	if request.Message == "delete" {
		d.Delete(w, *request)
	}

	if request.Message == "update" {
		d.Update(w, r, *request)
	}
}

func (d *websocketDelivery) Update(w http.ResponseWriter, r *http.Request, request models.UsersTripRequest) {
	trip, err := d.websocketUsecase.Update(r.Context(), request.TripId)
	if err != nil {
		log.Printf("error while getting trip info: %s", err)
		_, _ = w.Write([]byte("{}"))
		return
	}

	err = d.websocketUsecase.SendUpdateResponce(trip.Users, models.TripResponce{
		Message: request.Message,
		Trip:    *trip,
	})
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		_, _ = w.Write([]byte("{}"))
		return
	}
}

func (d *websocketDelivery) Delete(w http.ResponseWriter, request models.UsersTripRequest) {
	err := d.websocketUsecase.SendDeleteResponce(request.Users, models.TripRequest{
		Message: request.Message,
		TripId:  request.TripId,
	})
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		_, _ = w.Write([]byte("{}"))
		return
	}
}
