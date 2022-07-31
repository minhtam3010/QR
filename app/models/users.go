package models

import (
	"database/sql"
	"errors"
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

var (
	userID    int
	usernameT string
	TX        *sql.Tx
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
	// defer db.Close()

	rows, err := db.Query("SELECT * FROM Users ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	user := User{}
	res := []User{}
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

		res = append(res, user)
	}
	return res, nil
}

func GetUserById(id int, params ...string) (User, error) {
	db := config.GetDB()
	// defer db.Close()

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
	// return User{}, errors.New("err create user")

	db := config.GetDB()
	TX = config.GetTx()

	// change unix to datetime
	datecreated = time.Unix(u.DateCreated, 0)
	dateupdated = time.Unix(u.DateUpdated, 0)

	rowsAllUser, errQ := db.Query("SELECT ID, Username FROM users")
	if errQ != nil {
		panic(errQ)
	}
	switch {
	case dateupdated.Unix() <= 0:
		_, err := TX.Exec(`INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email,
							Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated)
							VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			u.ID, u.EntityCode, u.Username, u.Fullname, u.Password,
			u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan,
			u.Role, u.Hobby, datecreated)
		if err != nil {
			panic(err.Error())
		}
	default:
		_, err := TX.Exec(`INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email,
						Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated, DateUpdated)
						VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			u.ID, u.EntityCode, u.Username, u.Fullname, u.Password,
			u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan,
			u.Role, u.Hobby, datecreated)
		if err != nil {
			panic(err.Error())
		}
	}
	for rowsAllUser.Next() {
		errInside1 := rowsAllUser.Scan(&userID, &usernameT)
		if errInside1 != nil {
			panic(errInside1)
		} else {

			if userID == u.ID || u.EntityCode < 1 || u.EntityCode > 3 || usernameT == u.Username {
				return User{}, errors.New("error while creating User")
			}
		}
	}
	return *u, nil
}

func (u *User) UpdateUser(id int, username string) (User, error) {
	db := config.GetDB()
	TX := config.GetTx()

	if err := Check(id, "Users"); err != nil {
		return User{}, errors.New("error while updating user")
	}

	hold1, hold2 := u.ID, u.Username

	getUser, err := db.Query("SELECT ID, Username, DateCreated, DateUpdated FROM users WHERE ID=? AND username=?", id, username)
	if err != nil {
		log.Printf("Not found the user %v\n", id)
	}
	for getUser.Next() {
		err = getUser.Scan(&id, &username, &datecreated, &dateupdated)
		if err != nil {
			CheckTimeUpdateUser(u)
		}
		u.ID = id
		u.Username = username
		u.DateCreated = datecreated.Unix()
		u.DateUpdated = time.Now().Unix()
	}
	if hold1 != id && hold2 != username {
		return User{}, errors.New("cannot change ID and Username")
	}

	_, err = TX.Exec(`UPDATE users SET id=?, entitycode=?, username=?, fullname=?, password=?, email=?, address=?, bod=?, phone=?, qualification=?, 
								slogan=?, role=?, hobby=?, datecreated=?, dateupdated=? WHERE id=? AND username=?`,
		u.ID, u.EntityCode, u.Username, u.Fullname, u.Password, u.Email, u.Address, u.BOD, u.Phone, u.Qualification,
		u.Slogan, u.Role, u.Hobby, time.Unix(u.DateCreated, 0), time.Unix(u.DateUpdated, 0), u.ID, u.Username)

	if err != nil {
		panic(err.Error())
	}else if errCommit := TX.Commit(); errCommit != nil {
		log.Println("Error :(((")
	} else {
		log.Println("UPDATED User SUCCESSFULLY")
	}
	return *u, nil
}

func DeleteUser(id int, params ...string) error {
	TX = config.GetTx()

	err := Check(id, "Users")
	if err != nil {
		return errors.New("not found ID")
	}
	_, err = TX.Exec("DELETE FROM Users WHERE id= ?", id)
	if err != nil {
		panic(err.Error())
	} else if errCommit := TX.Commit(); errCommit != nil {
		return errors.New("error while commiting transaction")
	} else {
		log.Println("DELETED User SUCCESSFULLY")
	}
	return nil
}
