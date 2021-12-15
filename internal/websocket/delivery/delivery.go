package delivery

import (
	"encoding/json"
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
	cookie, err := r.Cookie(cnst.CookieName)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("{}"))
		return
	}

	userId, err := d.websocketUsecase.ValidateSession(r.Context(), cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("{}"))
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
		w.Write([]byte("{}"))
		return
	}

	d.websocketUsecase.Connect(userId, conn)
}

func (d *websocketDelivery) HandleRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	request := new(models.TripRequest)
	err := decoder.Decode(request)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}

	trip, err := d.websocketUsecase.Update(r.Context(), request.TripId)
	if err != nil {
		log.Printf("error while getting trip info: %s", err)
		w.Write([]byte("{}"))
		return
	}

	err = d.websocketUsecase.SendResponce(trip.Users, models.TripResponce{
		Message: request.Message,
		Trip:    *trip,
	})
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}
}
