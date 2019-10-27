package main

import (
	"net/http"

	resp "github.com/nicklaw5/go-respond"
	"github.com/urfave/negroni"
)

type Response struct {
	Success bool `json:"success"`
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		resp.NewResponse(w).Ok(&Response{true})
	})
	mux.HandleFunc("/user", func(w http.ResponseWriter, req *http.Request) {
		if true {
			resp.NewResponse(w).DefaultMessage().Unauthorized(nil)
		}
	})

	n := negroni.New()
	recovery := negroni.NewRecovery()
	recovery.ErrorHandlerFunc = func(error interface{}) {
		// do something with the unexpected error
	}

	n.Use(recovery)
	n.UseHandler(mux)

	http.ListenAndServe(":8001", n)
}
