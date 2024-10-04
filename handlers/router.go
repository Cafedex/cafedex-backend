package handlers

import (
	"net/http"

	"github.com/rs/cors"
)

type Response struct {
	Msg  string
	Code int
}

func CreateRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"Ya Motha\": \"Fuck Ya Life\"}"))
	})
	handler := cors.Default().Handler(mux)
	return handler
}
