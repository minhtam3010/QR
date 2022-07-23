package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhtam3010/qr/app/external/api"
)

func main() {
	r := mux.NewRouter()
	log.Println("Server started on: http://localhost:8080")
	api.RegisterRoutes(r)
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
