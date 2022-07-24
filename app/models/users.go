package models

import (
	"log"
	"time"

	"github.com/minhtam3010/qr/app/config"
)

type User struct {
	ID            int    `json:"id"`
	EntityCode    int    `json:"entityID"`
	Username      string `json:"username"`
	Fullname      string `json:"fullname"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	BOD           string `json:"bod"`
	Phone         string `json:"phone"`
	Qualification string `json:"qualification"`
	Slogan        string `json:"slogan"`
	Role          string `json:"role"`
	Hobby         string `json:"hobby"`
	DateCreated   int64  `json:"datecreated"`
	DateUpdated   int64  `json:"dateupdated"`
}

var (
	id, entityID                                                                                 int
	datecreated, dateupdated                                                                     time.Time
	username, fullname, password, email, address, bod, phone, qualification, slogan, role, hobby string
)

func GetUsers() []User {
	db := config.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	res := []User{}
	for rows.Next() {
		var id, entityID int
		var datecreated, dateupdated time.Time
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
		user.DateCreated = datecreated.Unix()
		user.DateUpdated = dateupdated.Unix()

		res = append(res, user)
	}

	return res
}

func GetUserById(id int, params ...string) User {
	db := config.GetDB()
	defer db.Close()

	switch {
	case len(params) == 1:
		fullname = params[0]
	case len(params) == 2:
		fullname = params[0]
		username = params[1]
	case len(params) == 3:
		log.Println("Error!!!")
	}

	rows, err := db.Query("SELECT * FROM users WHERE id=? OR fullname=? OR username=?", id, fullname, username)
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	for rows.Next() {
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
		user.DateCreated = datecreated.Unix()
		user.DateUpdated = dateupdated.Unix()

	}
	return user
}
