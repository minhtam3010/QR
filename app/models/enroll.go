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
	e.User.TransactionsUser()
	e.Guardian.TransactionsGuardian()
	return *e, nil
}

var (
	userID, EntityCode, GuardianID int
	usernameT                      string
)

func (u *User) TransactionsUser() {
	db := config.GetDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	rowsAllUser, _ := db.Query("SELECT ID, EntityCode, Username FROM users")
	datecreated := time.Unix(u.DateCreated, 0)
	dateupdated := time.Unix(u.DateUpdated, 0)

	switch {
	case dateupdated.Unix() <= 0:
		tx.Exec("INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", u.ID, u.EntityCode, u.Username, u.Fullname, u.Password, u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan, u.Role, u.Hobby, datecreated)
	default:
		tx.Exec("INSERT INTO Users(ID, EntityCode, Username, Fullname, Password, Email, Address, BOD, Phone, Qualification, Slogan, Role, Hobby, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", u.ID, u.EntityCode, u.Username, u.Fullname, u.Password, u.Email, u.Address, u.BOD, u.Phone, u.Qualification, u.Slogan, u.Role, u.Hobby, datecreated, dateupdated)
	}
	for rowsAllUser.Next() {
		errInside1 := rowsAllUser.Scan(&userID, &EntityCode, &usernameT)
		if errInside1 != nil {
			panic(errInside1)
		} else {
			if userID == u.ID || EntityCode < 1 || EntityCode > 3 || usernameT == u.Username {
				fmt.Println("Conflict")
				_ = tx.Rollback()
			}
		}
	}
	if errCommit := tx.Commit(); errCommit != nil {
		fmt.Println(errCommit)
	}
}

func (g *Guardian) TransactionsGuardian() {
	db := config.GetDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	RowAllGuardian, _ := db.Query("SELECT ID FROM guardians")
	datecreated := time.Unix(g.DateCreated, 0)
	dateupdated := time.Unix(g.DateUpdated, 0)

	switch {
	case dateupdated.Unix() <= 0:
		tx.Exec("INSERT INTO Guardians(ID, Fullname, Email, Address, BOD, Phone, Qualification, Role, DateCreated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)", g.ID, g.Fullname, g.Email, g.Address, g.BOD, g.Phone, g.Qualification, g.Role, datecreated)
	default:
		tx.Exec("INSERT INTO Guardians(ID, Fullname, Email, Address, BOD, Phone, Qualification, Role, DateCreated, DateUpdated) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", g.ID, g.Fullname, g.Email, g.Address, g.BOD, g.Phone, g.Qualification, g.Role, datecreated, dateupdated)
	}

	for RowAllGuardian.Next() {
		errInside1 := RowAllGuardian.Scan(&GuardianID)
		fmt.Println(GuardianID)
		if errInside1 != nil {
			panic(errInside1)
		} else {
			if GuardianID == g.ID {
				_ = tx.Rollback()
			}
		}
	}
	if errCommit := tx.Commit(); errCommit != nil {
		fmt.Println(errCommit)
	}
}
