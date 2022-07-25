package api

import (
	"github.com/gorilla/mux"
	"github.com/minhtam3010/qr/app/controllers"
)

var RegisterRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/user/", controllers.GetUsers).Methods("GET")
	routes.HandleFunc("/user/{id}/{username}/{fullname}", controllers.GetUserById).Methods("GET")
	routes.HandleFunc("/user/", controllers.CreateStudent).Methods("POST")
	routes.HandleFunc("/user/{id}/{username}/", controllers.UpdateUser).Methods("PUT")
	routes.HandleFunc("/user/{id}/{username}/", controllers.DeleteUser).Methods("DELETE")
	// routes.HandleFunc("/user/", controllers.CreateStudent).Methods("POST")
	// routes.HandleFunc("/guardian/", controllers.GetGuardians).Methods("GET")
	// routes.HandleFunc("/guardian/", controllers.CreateGuardian).Methods("POST")
}
