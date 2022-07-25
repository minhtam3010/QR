package models

import (
	"log"
	"time"

	"github.com/minhtam3010/qr/app/config"
)

type User struct {
	ID            int    `json:"id"`
	EntityCode    int    `json:"entityCode"`
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

func CheckTimeUpdateUser(u *User) {
	switch {
	case dateupdated.Unix() <= 0:
		u.DateUpdated = 0
	case dateupdated.Unix() > 0:
		u.DateUpdated = dateupdated.Unix()
	}
}

func GetUsers() ([]User, error) {
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
			CheckTimeUpdateUser(&user)
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
	return res, nil
}

func GetUserById(id int, params ...string) (User, error) {
	db := config.GetDB()
	defer db.Close()

	switch {
	case len(params) == 2:
		username = params[0]
		fullname = params[1]
	default:
		log.Println("Error!!!")
	}

	rows, err := db.Query("SELECT * FROM users WHERE id=? AND username=? AND fullname=?", id, username, fullname)
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	for rows.Next() {
		err = rows.Scan(&id, &entityID, &username, &fullname, &password, &email, &address, &bod, &phone, &qualification, &slogan, &role, &hobby, &datecreated, &dateupdated)
		if err != nil {
			CheckTimeUpdateUser(&user)
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
	return user, nil
}

func (u *User) CreateUser() (User, error) {
	db := config.GetDB()
	defer db.Close()

	// change unix to datetime
	datecreated = time.Unix(u.DateCreated, 0)
	dateupdated = time.Unix(u.DateUpdated, 0)

	switch {
	case dateupdated.Unix() <= 0:
		create, err := db.Prepare("INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		create.Exec(u.ID, u.EntityCode, u.Username, u.Fullname, u.Password, u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan, u.Role, u.Hobby, datecreated)
		log.Println("INSERT Successfully")
	default:
		create, err := db.Prepare("INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}

		create.Exec(u.ID, u.EntityCode, u.Username, u.Fullname, u.Password, u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan, u.Role, u.Hobby, datecreated, dateupdated)
		log.Println("INSERT Successfully")
	}

	return *u, nil
}

func (u *User) UpdateUser(id int, username string) (User, error) {
	db := config.GetDB()
	defer db.Close()

	getUser, err := db.Query("SELECT * FROM users WHERE ID=? AND username=?", id, username)
	if err != nil {
		log.Printf("Not found the user %v\n", id)
	}
	for getUser.Next() {
		err = getUser.Scan(&id, &entityID, &username, &fullname, &password, &email, &address, &bod, &phone, &qualification, &slogan, &role, &hobby, &datecreated, &dateupdated)
		if err != nil {
			CheckTimeUpdateUser(u)
		}
		u.ID = id
		u.Username = username
		u.DateCreated = datecreated.Unix()
		u.DateUpdated = time.Now().Unix()
	}
	updateForm, err := db.Prepare("UPDATE users SET id=?, entitycode=?, username=?, fullname=?, password=?, email=?, address=?, bod=?, phone=?, qualification=?, slogan=?, role=?, hobby=?, datecreated=?, dateupdated=? WHERE id=? AND username=?")
	if err != nil {
		panic(err.Error())
	}
	updateForm.Exec(u.ID, u.EntityCode, u.Username, u.Fullname, u.Password, u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan, u.Role, u.Hobby, time.Unix(u.DateCreated, 0), time.Unix(u.DateUpdated, 0), u.ID, u.Username)
	log.Println("UPDATED Successfully")
	return *u, nil
}

func DeleteUser(id int, params ...string) error {
	db := config.GetDB()
	defer db.Close()
	delForm, err := db.Prepare("DELETE FROM Users WHERE id= ? OR fullname=?")
	if err != nil {
		panic(err.Error())
	}
	if len(params) == 1 {
		delForm.Exec(id, params[0])
	} else {
		delForm.Exec(id, "")
	}

	log.Println("DELETED SUCCESSFULLY")
	return nil
}
