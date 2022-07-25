package models

import (
	"fmt"
	"log"
	"time"

	"github.com/minhtam3010/qr/app/config"
)

type Enroll struct {
	User     User     `json:"user"`
	Guardian Guardian `json:"guardian"`
}

func (e *Enroll) CreateEnroll() (Enroll, error) {
	err := e.Create()
	if err == 1 {
		guardian_student := &Guardian_Student{}
		guardian_student.StudentID = e.User.ID
		guardian_student.GuardianID = e.Guardian.ID
		guardian_student.CreateGS()
	} else {
		log.Println("Error :(((")
	}
	return *e, nil
}

var (
	userID, EntityCode, GuardianID int
	usernameT                      string
)

func (e *Enroll) Create() int {
	db := config.GetDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	rowsAllUser, _ := db.Query("SELECT ID, EntityCode, Username FROM users")
	datecreatedUser := time.Unix(e.User.DateCreated, 0)
	dateupdatedUSer := time.Unix(e.User.DateUpdated, 0)
	RowAllGuardian, _ := db.Query("SELECT ID FROM guardians")
	datecreatedGuardian := time.Unix(e.User.DateCreated, 0)
	dateupdatedGuardian := time.Unix(e.User.DateUpdated, 0)
	switch {
	case dateupdated.Unix() <= 0:
		tx.Exec("INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", e.User.ID, e.User.EntityCode, e.User.Username, e.User.Fullname, e.User.Password, e.User.Email, e.User.Address, e.User.BOD, e.User.Phone, e.User.Qualification, e.User.Slogan, e.User.Role, e.User.Hobby, datecreatedUser)
		tx.Exec("INSERT INTO Guardians(ID, Fullname, Email, Address, BOD, Phone, Qualification, Role, DateCreated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", e.Guardian.ID, e.Guardian.Fullname, e.Guardian.Email, e.Guardian.Address, e.Guardian.BOD, e.Guardian.Phone, e.Guardian.Qualification, e.Guardian.Role, datecreatedGuardian)

	default:
		tx.Exec("INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", e.User.ID, e.User.EntityCode, e.User.Username, e.User.Fullname, e.User.Password, e.User.Email, e.User.Address, e.User.BOD, e.User.Phone, e.User.Qualification, e.User.Slogan, e.User.Role, e.User.Hobby, datecreatedUser, dateupdatedUSer)
		tx.Exec("INSERT INTO Guardians(ID, Fullname, Email, Address, BOD, Phone, Qualification, Role, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", e.Guardian.ID, e.Guardian.Fullname, e.Guardian.Email, e.Guardian.Address, e.Guardian.BOD, e.Guardian.Phone, e.Guardian.Qualification, e.Guardian.Role, datecreatedGuardian, dateupdatedGuardian)
	}
	for rowsAllUser.Next() {
		errInside1 := rowsAllUser.Scan(&userID, &EntityCode, &usernameT)
		if errInside1 != nil {
			panic(errInside1)
		} else {
			if userID == e.User.ID || EntityCode < 1 || EntityCode > 3 || usernameT == e.User.Username {
				_ = tx.Rollback()
				return 2
			}
		}
	}
	for RowAllGuardian.Next() {
		errInside1 := RowAllGuardian.Scan(&GuardianID)
		if errInside1 != nil {
			panic(errInside1)
		} else {
			if GuardianID == e.Guardian.ID {
				_ = tx.Rollback()
				return 2
			}
		}
	}
	if errCommit := tx.Commit(); errCommit != nil {
		fmt.Println(errCommit)
	}
	return 1
}

// func (u *User) TransactionsUser() int{
// 	db := config.GetDB()
// 	defer db.Close()

// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	rowsAllUser, _ := db.Query("SELECT ID, EntityCode, Username FROM users")
// 	datecreated := time.Unix(u.DateCreated, 0)
// 	dateupdated := time.Unix(u.DateUpdated, 0)

// 	switch {
// 	case dateupdated.Unix() <= 0:
// 		tx.Exec("INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", u.ID, u.EntityCode, u.Username, u.Fullname, u.Password, u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan, u.Role, u.Hobby, datecreated)
// 	default:
// 		tx.Exec("INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", u.ID, u.EntityCode, u.Username, u.Fullname, u.Password, u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan, u.Role, u.Hobby, datecreated, dateupdated)
// 	}
// 	for rowsAllUser.Next() {
// 		errInside1 := rowsAllUser.Scan(&userID, &EntityCode, &usernameT)
// 		if errInside1 != nil {
// 			panic(errInside1)
// 		} else {
// 			if userID == u.ID || EntityCode < 1 || EntityCode > 3 || usernameT == u.Username {
// 				_ = tx.Rollback()
// 				return 2
// 			}
// 		}
// 	}
// 	if errCommit := tx.Commit(); errCommit != nil {
// 		fmt.Println(errCommit)
// 	}
// 	return 1
// }

// func (g *Guardian) TransactionsGuardian() int{
// 	db := config.GetDB()
// 	defer db.Close()

// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	RowAllGuardian, _ := db.Query("SELECT ID FROM guardians")
// 	datecreated := time.Unix(g.DateCreated, 0)
// 	dateupdated := time.Unix(g.DateUpdated, 0)

// 	switch {
// 	case dateupdated.Unix() <= 0:
// 		tx.Exec("INSERT INTO Guardians(ID, Fullname, Email, Address, BOD, Phone, Qualification, Role, DateCreated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", g.ID, g.Fullname, g.Email, g.Address, g.BOD, g.Phone, g.Qualification, g.Role, datecreated)
// 	default:
// 		tx.Exec("INSERT INTO Guardians(ID, Fullname, Email, Address, BOD, Phone, Qualification, Role, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", g.ID, g.Fullname, g.Email, g.Address, g.BOD, g.Phone, g.Qualification, g.Role, datecreated, dateupdated)
// 	}

// 	for RowAllGuardian.Next() {
// 		errInside1 := RowAllGuardian.Scan(&GuardianID)
// 		if errInside1 != nil {
// 			panic(errInside1)
// 		} else {
// 			if GuardianID == g.ID {
// 				_ = tx.Rollback()
// 				return 2
// 			}
// 		}
// 	}
// 	if errCommit := tx.Commit(); errCommit != nil {
// 		fmt.Println(errCommit)
// 	}
// 	return 1
// }
