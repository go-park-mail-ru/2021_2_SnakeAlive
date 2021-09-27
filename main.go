package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
<<<<<<< HEAD
	"strconv"

	"snakealive/m/http"
=======
	"log"
	ent "snakealive/m/entities"
	DB "snakealive/m/storage"
	"snakealive/m/validate"
	"strconv"
>>>>>>> feature/tests

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

<<<<<<< HEAD
func Router() *router.Router {
	r := router.New()

	r.POST("/login", http.Login)
	r.POST("/register", http.Registration)
	r.GET("/country/{name}", http.PlacesList)
	return r
=======
const CookieName = "SnakeAlive"

type placesListJSON struct {
	Places []ent.Place
>>>>>>> feature/tests
}

func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 10)
<<<<<<< HEAD
=======
}

func SetCookie(ctx *fasthttp.RequestCtx, cookie string, user ent.User) {
	var c fasthttp.Cookie
	c.SetKey(CookieName)
	c.SetValue(cookie)
	ctx.Response.Header.SetCookie(&c)

	DB.CookieDB[cookie] = user
}

func CheckCookie(ctx *fasthttp.RequestCtx) bool {
	if _, found := DB.CookieDB[string(ctx.Request.Header.Cookie(CookieName))]; !found {
		return false
	}
	return true
}

func Router() *router.Router {
	r := router.New()

	r.POST("/login", login)
	r.POST("/register", registration)
	r.GET("/country/{name}", placesList)
	return r
>>>>>>> feature/tests
}

func corsMiddleware(handler func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*") // set domain
		ctx.Response.Header.Set("Content-Type", "application/json; charset=utf8")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Request-ID")
		ctx.Response.Header.Set("Access-Control-Expose-Headers", "Authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.Set("Access-Control-Max-Age", "3600")

		if bytes.Equal(ctx.Method(), []byte(fasthttp.MethodOptions)) {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		handler(ctx)
	}
}

<<<<<<< HEAD
=======
func placesList(ctx *fasthttp.RequestCtx) {
	if !CheckCookie(ctx) {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return
	}

	param, _ := ctx.UserValue("name").(string)
	if _, found := DB.PlacesDB[param]; !found {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)

	response := placesListJSON{DB.PlacesDB[param]}
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		ctx.Write([]byte("{}"))
		return
	}

	ctx.Write(bytes)
}

func login(ctx *fasthttp.RequestCtx) {
	user := new(ent.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if _, found := DB.AuthDB[user.Email]; !found {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	password := DB.AuthDB[user.Email].Password

	if password != user.Password {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	SetCookie(ctx, Hash(user.Email), DB.AuthDB[user.Email])
}

func registration(ctx *fasthttp.RequestCtx) {
	user := new(ent.User)

	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if !validate.Validate(*user) {
		log.Printf("error while validate user:")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if _, found := DB.AuthDB[user.Email]; found {
		log.Printf("User with this email already exists")
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	DB.AuthDB[user.Email] = *user

	ctx.SetStatusCode(fasthttp.StatusOK)
	SetCookie(ctx, Hash(user.Email), DB.AuthDB[user.Email])
}

>>>>>>> feature/tests
func main() {
	fmt.Println("starting server at :8080")

	r := Router()

<<<<<<< HEAD
	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
=======
	if err := fasthttp.ListenAndServe(":8081", corsMiddleware(r.Handler)); err != nil {
>>>>>>> feature/tests
		fmt.Println("failed to start server:", err)
		return
	}
}
