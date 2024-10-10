package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cafedex-backend/services"

	"github.com/rs/cors"
)

type Response struct {
	Msg  string
	Code int
}

func CreateRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Cafedex backend Api"))
		log.Default().Printf(r.Method + " at " + r.RequestURI)

	})

	// mux.HandleFunc("GET /api/guides/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")
	// 	fmt.Fprintf(w, "Fetching guide id: %s", id)
	// 	log.Default().Printf(r.Method + " at " + r.RequestURI)

	// 	services.GetGuideById(id)
	// })

	mux.HandleFunc("POST /api/guides/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Got it, creating a new Guide...")
		log.Default().Printf(r.Method + " at " + r.RequestURI)

		services.CreateGuide(w, r)
	})

	// mux.HandleFunc("PUT /api/guides/update/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")
	// 	fmt.Fprintf(w, "Roger that, updating guide id:%s", id)
	// 	log.Default().Printf(r.Method + " at " + r.RequestURI)

	// 	// services.UpdateGuide()
	// })

	// mux.HandleFunc("DELETE /api/guides/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")
	// 	fmt.Fprintf(w, "Understood, removing guide id:%s", id)
	// 	log.Default().Printf(r.Method + " at " + r.RequestURI)

	// 	// services.DeleteGuide(id)
	// })

	handler := cors.Default().Handler(mux)
	return handler
}
