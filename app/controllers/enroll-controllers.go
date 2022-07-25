package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/minhtam3010/qr/app/models"
	"github.com/minhtam3010/qr/app/utils"
)

func CreateEnroll(w http.ResponseWriter, r *http.Request) {
	createEnroll := &models.Enroll{}
	utils.ParseBody(r, createEnroll)
	// errT := transactions.TransactionsEnroll(createEnroll.User.ID, createEnroll.User.EntityCode, createEnroll.User.Username, createEnroll.Guardian.ID)
	// if errT != nil {
	// 	w.Write([]byte("Your input was duplicated in database"))
	// 	panic(errT)
	// }

	enroll, err := createEnroll.CreateEnroll()
	if err != nil {
		panic(err)
	}
	if res, err := json.Marshal(enroll); err == nil {
		WriteResponse(w, res)
	} else {
		panic(err)
	}
}
