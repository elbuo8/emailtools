package main

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"net/http"
	"net/mail"
)

func main() {
	m := martini.Classic()
	m.Get("/:email", func(params martini.Params, res http.ResponseWriter) {
		email, e := mail.ParseAddress(params["email"])
		if e == nil {
			email, _ := json.Marshal(email)
			res.Write([]byte(email))
			res.WriteHeader(200)
		} else {
			res.Write([]byte("{}"))
			res.WriteHeader(400)
		}
	})
	m.Run()
}
