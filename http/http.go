package http

import (
	"encoding/json"
	"hash/fnv"
	"log"
	ent "snakealive/m/entities"
	DB "snakealive/m/storage"
	"snakealive/m/validate"
	"strconv"

	"github.com/valyala/fasthttp"
)

const CookieName = "SnakeAlive"

type placesListJSON struct {
	Places []ent.Place
}

func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 10)
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

func Login(ctx *fasthttp.RequestCtx) {
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

func Registration(ctx *fasthttp.RequestCtx) {
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

func PlacesList(ctx *fasthttp.RequestCtx) {
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
