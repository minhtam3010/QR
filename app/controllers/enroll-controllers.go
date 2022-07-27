package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/minhtam3010/qr/app/usecase"
	"github.com/minhtam3010/qr/app/utils"
)

func CreateEnroll(w http.ResponseWriter, r *http.Request) {
	createEnroll := &usecase.Enroll{}
	utils.ParseBody(r, createEnroll)
	e, err := createEnroll.CreateEnroll()
	if err != nil {
		w.Write([]byte("Error :(((("))
		// panic(err)
	}else{
		if res, err := json.Marshal(e); err == nil {
			WriteResponse(w, res)
		}
	}
}
