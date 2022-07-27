package models

import (
	"errors"
	"log"
	"time"

	"github.com/minhtam3010/qr/app/config"
)

type Guardian struct {
	ID            int    `json:"id"`
	Fullname      string `json:"fullname"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	BOD           string `json:"bod"`
	Phone         string `json:"phone"`
	Qualification string `json:"qualification"`
	Role          string `json:"role"`
	DateCreated   int64  `json:"datecreated"`
	DateUpdated   int64  `json:"dateupdated"`
}

var GuardianID int

func CheckTimeUpdateGuardian(g *Guardian) {
	switch {
	case dateupdated.Unix() <= 0:
		g.DateUpdated = 0
	case dateupdated.Unix() > 0:
		g.DateUpdated = dateupdated.Unix()
	}
}

func GetGuardians() ([]Guardian, error) {
	db := config.GetDB()
	// defer db.Close()

	rows, err := db.Query("SELECT * FROM guardians ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	guardian := Guardian{}
	res := []Guardian{}
	for rows.Next() {

		err = rows.Scan(&id, &fullname, &email, &address, &bod, &phone, &qualification, &role, &datecreated, &dateupdated)
		if err != nil {
			CheckTimeUpdateGuardian(&guardian)
		}
		guardian.ID = id
		guardian.Fullname = fullname
		guardian.Email = email
		guardian.Address = address
		guardian.BOD = bod
		guardian.Phone = phone
		guardian.Qualification = qualification
		guardian.Role = role
		guardian.DateCreated = datecreated.Unix()
		guardian.DateUpdated = dateupdated.Unix()

		res = append(res, guardian)
	}
	return res, nil
}

func GetGuardiansByID(id int, params ...string) (Guardian, error) {
	db := config.GetDB()
	// defer db.Close()

	switch {
	case len(params) == 1:
		fullname = params[0]
	case len(params) > 1:
		log.Println("Error!!!")
	}

	rows, err := db.Query("SELECT * FROM guardians WHERE id=? AND fullname=?", id, fullname)
	if err != nil {
		panic(err.Error())
	}

	guardian := Guardian{}
	for rows.Next() {
		err := rows.Scan(&id, &fullname, &email, &address, &bod, &phone, &qualification, &role, &datecreated, &dateupdated)

		if err != nil {
			CheckTimeUpdateGuardian(&guardian)
		}
		guardian.ID = id
		guardian.Fullname = fullname
		guardian.Email = email
		guardian.Address = address
		guardian.BOD = bod
		guardian.Phone = phone
		guardian.Qualification = qualification
		guardian.Role = role
		guardian.DateCreated = datecreated.Unix()
	}
	return guardian, nil
}

func (g *Guardian) CreateGuardian() (Guardian, error) {
	// return Guardian{}, errors.New("fail at create guardian")
	db := config.GetDB()
	// defer DB.Close()
	// TX, _ = DB.Begin()
	// tx, err := DB.Begin()

	// change unix to datetime
	datecreated = time.Unix(g.DateCreated, 0)
	dateupdated = time.Unix(g.DateUpdated, 0)

	RowAllGuardian, errQ := db.Query("SELECT ID FROM guardians")

	if errQ != nil {
		panic(errQ)
	}

	switch {
	case dateupdated.Unix() <= 0:
		_, err := TX.Exec(`INSERT INTO Guardians(ID, Fullname, Email, Address,
			BOD, Phone, Qualification, Role, DateCreated)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			g.ID, g.Fullname, g.Email,
			g.Address, g.BOD, g.Phone,
			g.Qualification, g.Role, datecreated)

		if err != nil {
			panic(err.Error())
		}
	default:
		_, err := TX.Exec(`INSERT INTO Guardians(ID, Fullname, Email, Address,
				BOD, Phone, Qualification, Role, DateCreated, DateUpdated)
				VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			g.ID, g.Fullname, g.Email,
			g.Address, g.BOD, g.Phone,
			g.Qualification, g.Role,
			datecreated, dateupdated)
		if err != nil {
			panic(err.Error())
		}
	}
	for RowAllGuardian.Next() {
		errInside1 := RowAllGuardian.Scan(&GuardianID)
		if errInside1 != nil {
			panic(errInside1)
		} else {
			if GuardianID == g.ID {
				// _ = TX.Rollback()
				return Guardian{}, errors.New("error while creating Guardian")
			}
		}
	}
	return *g, nil
}

func (g *Guardian) UpdateGuardian(id int) (Guardian, error) {
	db := config.GetDB()
	// defer db.Close()

	getGuardian, err := db.Query("SELECT * FROM Guardians WHERE ID=?", id)
	if err != nil {
		log.Printf("Not found the Guardians %v\n", id)
	}
	for getGuardian.Next() {
		err = getGuardian.Scan(&id, &fullname, &email, &address, &bod, &phone, &qualification, &role, &datecreated, &dateupdated)
		if err != nil {
			CheckTimeUpdateGuardian(g)
		}
		g.ID = id
		g.DateCreated = datecreated.Unix()
		g.DateUpdated = time.Now().Unix()
	}
	updateForm, err := db.Prepare("UPDATE guardians SET id=?, fullname=?, email=?, address=?, bod=?, phone=?, qualification=?, role=?, datecreated=?, dateupdated=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	updateForm.Exec(g.ID, g.Fullname, g.Email, g.Address, g.BOD, g.Phone, g.Qualification, g.Role, time.Unix(g.DateCreated, 0), time.Unix(g.DateUpdated, 0), g.ID)
	log.Println("UPDATED Guardian Successfully")
	return *g, nil
}

func DeleteGuardian(id int) error {
	db := config.GetDB()
	// defer db.Close()
	delForm, err := db.Prepare("DELETE FROM Guardians WHERE id= ?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)

	log.Println("DELETED Guardian SUCCESSFULLY")
	return nil
}
