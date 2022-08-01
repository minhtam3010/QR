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

func WriteResponse(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := models.GetUsers()
	if res, err := json.Marshal(users); err == nil {
		WriteResponse(w, res)
	} else {
		panic(err)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	username := vars["username"]
	fullname := vars["fullname"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		panic(err)
	}
	userDetails, err := models.GetUserById(int(ID), username, fullname)
	if err != nil {
		panic(err)
	}
	if res, err := json.Marshal(userDetails); err == nil {
		WriteResponse(w, res)
	} else {
		log.Printf("Error: %v\n", err)
	}
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	createUser := &models.User{}
	utils.ParseBody(r, createUser)
	user, err := createUser.CreateUser()
	if err != nil {
		w.Write([]byte("Cannot duplicate ID or Username"))
		return
	}
	if errCommit := models.TX.Commit(); err != nil {
		panic(errCommit)
	}

	if res, err := json.Marshal(user); err == nil {
		WriteResponse(w, res)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	updateUser := &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["id"]
	username := vars["username"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}
	userDetails, err := updateUser.UpdateUser(int(ID), username)
	if err != nil {
		w.Write([]byte("U cannot change ID or username"))
		return
	} else if res, err := json.Marshal(userDetails); err == nil {
		WriteResponse(w, res)
	} else {
		panic(err)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)

	if err != nil {
		panic(err)
	} else if err = models.DeleteUser(int(ID)); err != nil {
		w.Write([]byte("Not found ID in this table"))
	} else {
		w.Write([]byte("DELETED Successfully"))
	}
}
