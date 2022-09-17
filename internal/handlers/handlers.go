package handlers

import "github.com/gorilla/mux"

func ApplyHandlers(router *mux.Router) error {
	router.HandleFunc("/", Handler)
	router.HandleFunc("/echo", EchoHandler)
	return nil
}
