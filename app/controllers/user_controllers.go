package controllers

import (
	"log"
	"net/http"

	"github.com/minhtam3010/qr/app/config"
	"github.com/minhtam3010/qr/app/models"
)

func GetUsers(id int, fullname string, username string) []models.User {
	db := config.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM User ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	user := models.User{}
	res := []models.User{}
	for rows.Next() {
		var id, entityID int
		var datecreated, dateupdated int64
		var username, fullname, password, email, address, bod, phone, qualification, slogan, role, hobby string
		err = rows.Scan(&id, &entityID, &username, &fullname, &password, &email, &address, &bod, &phone, &qualification, &slogan, &role, &hobby, &datecreated, &dateupdated)
		if err != nil {
			panic(err.Error())
		}
		user.ID = id
		user.EntityCode = entityID
		user.Username = username
		user.Fullname = fullname
		user.Password = password
		user.Email = email
		user.Address = address
		user.BOD = bod
		user.Phone = phone
		user.Qualification = qualification
		user.Slogan = slogan
		user.Role = role
		user.Hobby = hobby
		user.DateCreated = datecreated
		user.DateUpdated = dateupdated

		res = append(res, user)
	}

	return res
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	if r.Method == "POST" {
		ID := r.FormValue("id")
		EntityCode := r.FormValue("entityID")
		Username := r.FormValue("username")
		Fullname := r.FormValue("fullname")
		Password := r.FormValue("password")
		Email := r.FormValue("email")
		Address := r.FormValue("address")
		BOD := r.FormValue("bod")
		Phone := r.FormValue("phone")
		Qualification := r.FormValue("qualification")
		Slogan := r.FormValue("slogan")
		Role := r.FormValue("role")
		Hobby := r.FormValue("hobby")
		DateCreated := r.FormValue("datecreated")
		DateUpdated := r.FormValue("dateupdated")

		create, err := db.Prepare("INSERT INTO User(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		create.Exec(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated, DateUpdated)
		log.Println("INSERT Successfully")
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
