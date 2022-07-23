package models

import (
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

func GetUsers() []User {
	db := config.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM User ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	res := []User{}
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

func GetUserById(id int, fullname string, username string) User {
	db := config.GetDB()
	defer db.Close()
	// nId := r.URL.Query().Get("id")

	rows, err := db.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	user := User{}
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

	}
	return user
}
