package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/minhtam3010/qr/app/models"
	"github.com/minhtam3010/qr/app/utils"
)

func GetCaregivers(w http.ResponseWriter, r *http.Request) {
	allCG, err := models.GetCaregivers()
	if err != nil {
		w.Write([]byte("Error while querying :((("))
	} else if res, err := json.Marshal(allCG); err == nil {
		WriteResponse(w, res)
	}
}

func GetCaregiverByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cgID := vars["CaregiverID"]
	ID, err := strconv.ParseInt(cgID, 0, 0)
	if err != nil {
		log.Println("Error while parsing ID")
	}
	cgDetails, err := models.GetCaregiverByID(int(ID))
	if err != nil {
		w.Write([]byte("Error while getting caregiver"))
	}else if res, err := json.Marshal(cgDetails); err == nil {
		WriteResponse(w, res)
	}
}

func CreateCaregiver(w http.ResponseWriter, r *http.Request) {
	caregiver := &models.Caregiver{}
	utils.ParseBody(r, caregiver)
	cg, err := caregiver.CreateCaregiver()
	if err != nil {
		w.Write([]byte("Error while creating :((("))
	} else if res, err := json.Marshal(cg); err == nil {
		WriteResponse(w, res)
	}
}

func UpdateCaregiver(w http.ResponseWriter, r *http.Request) {
	caregiver := &models.Caregiver{}
	utils.ParseBody(r, caregiver)
	vars := mux.Vars(r)
	cgID := vars["CaregiverID"]
	ID, err := strconv.ParseInt(cgID, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}
	updateCG, err := caregiver.UpdateCaregiver(int(ID))
	if err != nil {
		w.Write([]byte("Error while updating Caregiver :(((("))
	}else if res, err := json.Marshal(updateCG); err == nil {
		WriteResponse(w, res)
	}
}

func DeleteCaregiver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["CaregiverID"], 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}
	err = models.DeleteCaregiver(int(ID))
	if err != nil {
		w.Write([]byte("Not found id in that table or Error while deleting Caregiver"))
	} else {
		w.Write([]byte("DELETED Successfully"))
	}
}