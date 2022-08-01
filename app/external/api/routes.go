package api

import (
	"github.com/gorilla/mux"
	"github.com/minhtam3010/qr/app/controllers"
)

var RegisterRoutes = func(routes *mux.Router) {
	// API for User
	routes.HandleFunc("/user/", controllers.GetUsers).Methods("GET")
	routes.HandleFunc("/user/{id}/{username}/{fullname}/", controllers.GetUserById).Methods("GET")
	routes.HandleFunc("/user/", controllers.CreateStudent).Methods("POST")
	routes.HandleFunc("/user/{id}/{username}/", controllers.UpdateUser).Methods("PUT")
	routes.HandleFunc("/user/{id}/", controllers.DeleteUser).Methods("DELETE")

	// API for Guardian
	routes.HandleFunc("/guardian/", controllers.GetGuardians).Methods("GET")
	routes.HandleFunc("/guardian/{id}/{fullname}/", controllers.GetGuardianById).Methods("GET")
	routes.HandleFunc("/guardian/", controllers.CreateGuardian).Methods("POST")
	routes.HandleFunc("/guardian/{id}/", controllers.UpdateGuardian).Methods("PUT")
	routes.HandleFunc("/guardian/{id}/", controllers.DeleteGuardian).Methods("DELETE")

	// API for enroll
	routes.HandleFunc("/enroll/", controllers.CreateEnroll).Methods("POST")

	// API for caregiver
	routes.HandleFunc("/caregiver/", controllers.GetCaregivers).Methods("GET")
	routes.HandleFunc("/caregiver/{CaregiverID}/", controllers.GetCaregiverByID).Methods("GET")
	routes.HandleFunc("/caregiver/", controllers.CreateCaregiver).Methods("POST")
	routes.HandleFunc("/caregiver/{CaregiverID}/", controllers.UpdateCaregiver).Methods("PUT")
	routes.HandleFunc("/caregiver/{CaregiverID}/", controllers.DeleteCaregiver).Methods("DELETE")
}
