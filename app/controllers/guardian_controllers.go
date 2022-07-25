package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/minhtam3010/qr/app/models"
	"github.com/minhtam3010/qr/app/utils"
)

func GetGuardians(w http.ResponseWriter, r *http.Request) {
	guardians, _ := models.GetGuardians()
	if res, err := json.Marshal(guardians); err == nil {
		WriteResponse(w, res)
	} else {
		panic(err)
	}
}

func GetGuardianById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guardianID := vars["id"]
	fullname := vars["fullname"]
	fmt.Println(fullname)
	ID, err := strconv.ParseInt(guardianID, 0, 0)
	if err != nil {
		panic(err)
	}
	guardianDetails, err := models.GetGuardiansByID(int(ID), fullname)
	if err != nil {
		panic(err)
	}
	if res, err := json.Marshal(guardianDetails); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func CreateGuardian(w http.ResponseWriter, r *http.Request) {
	createGuardian := &models.Guardian{}
	utils.ParseBody(r, createGuardian)
	guardian, err := createGuardian.CreateGuardian()
	if err != nil {
		panic(err)
	}
	if res, err := json.Marshal(guardian); err == nil {
		WriteResponse(w, res)
	}else{
		log.Println("Error :(((")
	}
}

func UpdateGuardian(w http.ResponseWriter, r *http.Request) {
	updateGuardian := &models.Guardian{}
	utils.ParseBody(r, updateGuardian)
	vars := mux.Vars(r)
	guardianID := vars["id"]
	ID, err := strconv.ParseInt(guardianID, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}
	guardianDetails, err := updateGuardian.UpdateGuardian(int(ID))
	if err != nil{
		panic(err)
	}
	if res, err := json.Marshal(guardianDetails); err == nil {
		WriteResponse(w, res)
	}else {
		panic(err)
	}

}

func DeleteGuardian(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guardianID := vars["id"]
	ID, err := strconv.ParseInt(guardianID, 0, 0)

	if err != nil {
		panic(err)
	}
	if err := models.DeleteGuardian(int(ID)); err != nil {
		panic(err)
	}
	w.Write([]byte("DELETED guarian Successfully"))
}

