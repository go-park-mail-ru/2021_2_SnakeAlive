package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var auth = map[string]string{
	"alex@mail.ru": "pass",
}

type User struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Token struct {
	Token string `json:"token,omitempty"`
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	user := new(User)
	if err := decoder.Decode(user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password, found := auth[user.Email]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if password != user.Password {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	t := Token{"token"}
	bytes, err := json.Marshal(t)

	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		w.Write([]byte("{}"))
		return
	}

	w.Write(bytes)
}

func handlerDecorator(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Request-ID")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/login", handlerDecorator(loginPage))

	fmt.Println("starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("failed to start server:", err)
		return
	}
}
