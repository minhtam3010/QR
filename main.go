package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhtam3010/qr/app/config"
	"github.com/minhtam3010/qr/app/external/api"
)

func main() {
	r := mux.NewRouter()
	log.Println("Server started on: http://localhost:8080")
	api.RegisterRoutes(r)
	http.Handle("/", r)

	config.ConnectDB()
	db := config.GetDB()
	defer db.Close()
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
