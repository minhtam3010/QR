package controllers

import (
	"log"
	"net/http"

	"github.com/minhtam3010/qr/app/config"
	// "github.com/minhtam3010/qr/app/models"
)

// func GetGuardians(w http.ResponseWriter, r *http.Request) {
// 	db := config.GetDB()
// 	selfDB, err := db.Query("SELECT * FROM User ORDER BY id ASC")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	guardian := models.Guardian{}
// 	res := []models.Guardian{}
//     for selfDB.Next() {
//         var id int
//         var datecreated, dateupdated int64
//         var fullname, email, address, bod, phone, qualification, role string
//         err = selfDB.Scan(&id, &fullname, &email, &address, &bod, &phone, &qualification, &role, &datecreated, &dateupdated)
//         if err != nil {
//             panic(err.Error())
//         }
//         guardian.ID = id
//         guardian.Fullname = fullname
//         guardian.Email = email
//         guardian.Address = address
//         guardian.BOD = bod
//         guardian.Phone = phone
//         guardian.Qualification = qualification
//         guardian.Role = role
//         guardian.DateCreated = datecreated
//         guardian.DateUpdated = dateupdated
//         res = append(res, guardian)
//     }
//     defer db.Close()
// }

func CreateGuardian(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	if r.Method == "POST" {
		ID := r.FormValue("id")
		Fullname := r.FormValue("fullname")
		Email := r.FormValue("email")
		Address := r.FormValue("address")
		BOD := r.FormValue("bod")
		Phone := r.FormValue("phone")
		Qualification := r.FormValue("qualification")
		Role := r.FormValue("role")
		DateCreated := r.FormValue("datecreated")
		DateUpdated := r.FormValue("dateupdated")

		create, err := db.Prepare("INSERT INTO User(ID, Fullname, Email, Address, BOD, Phone, Qualification, Role, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		create.Exec(ID, Fullname, Email, Address, BOD, Phone, Qualification, Role, DateCreated, DateUpdated)
		log.Println("INSERT Successfully")
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
