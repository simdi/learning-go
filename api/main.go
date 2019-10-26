package main

import (
	"net/http"

	resp "github.com/nicklaw5/go-respond"
	"github.com/urfave/negroni"
)

type Response struct {
	Success bool `json:"success"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		resp.NewResponse(w).Ok(&Response{true})
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
