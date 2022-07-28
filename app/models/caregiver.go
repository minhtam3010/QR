package models

import (
	"errors"
	"log"
	"time"

	"github.com/minhtam3010/qr/app/config"
)

type Caregiver struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	BOD         string `json:"bod"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	DateCreated int64  `json:"datecreated"`
	DateUpdated int64  `json:"dateupdated"`
}

var (
	name string
)

func CheckTimeUpdateCaregiver(c *Caregiver) {
	switch {
	case dateupdated.Unix() <= 0:
		c.DateUpdated = 0
	case dateupdated.Unix() > 0:
		c.DateUpdated = dateupdated.Unix()
	}
}

func Check(ID int) error {
	db := config.GetDB()

	rows, err := db.Query("SELECT id FROM caregivers ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	for rows.Next(){
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}else if (id == ID) {
			return nil
		}
	}
	return errors.New("not found any id in table")
}

func GetCaregivers() ([]Caregiver, error) {
	db := config.GetDB()

	rows, err := db.Query("SELECT * FROM caregivers ORDER BY id ASC")
	if err != nil {
		log.Println(err.Error())
		return []Caregiver{}, errors.New("error while querying")
	}
	caregiver := Caregiver{}
	res := []Caregiver{}
	for rows.Next() {
		err = rows.Scan(&id, &name, &bod, &address, &phone, &datecreated, &dateupdated)
		if err != nil {
			CheckTimeUpdateCaregiver(&caregiver)
		}
		caregiver.ID = id
		caregiver.BOD = bod
		caregiver.Address = address
		caregiver.Phone = phone
		caregiver.DateCreated = datecreated.Unix()
		caregiver.DateUpdated = dateupdated.Unix()

		res = append(res, caregiver)
	}
	return res, nil
}

func GetCaregiverByID(ID int) (Caregiver, error) {
	db := config.GetDB()

	row, err := db.Query("SELECT * FROM caregivers WHERE id=?", ID)
	if err != nil {
		log.Println(err.Error())
		return Caregiver{}, errors.New("error while querying or no ID")
	}
	caregiver := Caregiver{}

	for row.Next() {
		err = row.Scan(&id, &name, &bod, &address, &phone, &datecreated, &dateupdated)
		if err != nil {
			CheckTimeUpdateCaregiver(&caregiver)
		}

		caregiver.ID = id
		caregiver.Name = name
		caregiver.BOD = bod
		caregiver.Address = address
		caregiver.Phone = phone
		caregiver.DateCreated = datecreated.Unix()
		caregiver.DateUpdated = dateupdated.Unix()
	}
	return caregiver, nil
}

func (cg *Caregiver) CreateCaregiver() (Caregiver, error) {
	loop := true
	db := config.GetDB()

	TX = config.GetTx()

	datecreated = time.Unix(cg.DateCreated, 0)
	dateupdated = time.Unix(cg.DateUpdated, 0)

	rowsAllCaregiver, err := db.Query("SELECT ID FROM caregivers")
	if err != nil {
		log.Println("OK?")
		loop = false
		log.Println(err.Error())
	}

	switch {
	case datecreated.Unix() <= 0:
		_, err := TX.Exec(`INSERT INTO caregivers(ID, Name, BOD, Address, Phone, DateCreated)
							VALUES(?, ?, ?, ?, ?, ?)`, cg.ID, cg.Name, cg.BOD, cg.Address,
			cg.Phone, datecreated)
		if err != nil {
			panic(err.Error())
		}
	default:
		_, err := TX.Exec(`INSERT INTO caregivers(ID, Name, BOD, Address, Phone, DateCreated, DateUpdated)
		VALUES(?, ?, ?, ?, ?, ?, ?)`, cg.ID, cg.Name, cg.BOD, cg.Address,
			cg.Phone, datecreated, dateupdated)
		if err != nil {
			panic(err.Error())
		}
	}
	if loop {
		for rowsAllCaregiver.Next(){
			err = rowsAllCaregiver.Scan(&id)
			if err != nil {
				panic(err)
			} else {
				if id == cg.ID {
					_ = TX.Rollback()
					return Caregiver{}, errors.New("error while creating Caregiver")
				}
			}
		}
	}
	if errCommit := TX.Commit(); errCommit != nil {
		log.Println("Error")
	}
	return *cg, nil
}

func (cg *Caregiver) UpdateCaregiver(id int) (Caregiver, error){
	db := config.GetDB()

	TX = config.GetTx()
	hold := cg.ID
	getCG, err := db.Query("SELECT id, DateCreated, DateUpdated FROM caregivers WHERE ID=?", id)
	if err != nil {
		panic(err)
	}
	for getCG.Next() {
		err = getCG.Scan(&id, &datecreated, &dateupdated)
		if err != nil {
			CheckTimeUpdateCaregiver(cg)
		}
		cg.ID = id
		cg.DateCreated = datecreated.Unix()
		cg.DateUpdated = time.Now().Unix()
	}
	if id != hold {
		_ = TX.Rollback()
		return Caregiver{}, errors.New("error while updating Caregivers")
	}
	_, err = TX.Exec(`UPDATE caregivers SET ID=?, Name=?, BOD=?, Address=?, Phone=?, DateCreated=?, DateUpdated=? WHERE ID=?`, id, cg.Name, cg.BOD, cg.Address, cg.Phone, time.Unix(cg.DateCreated, 0), time.Unix(cg.DateUpdated, 0), id)
	if err != nil {
		panic(err)
	}else if errCommit := TX.Commit(); errCommit != nil {
		log.Println("Error")
	}
	return *cg, nil
}

func DeleteCaregiver(id int) error {
	TX := config.GetTx()
	err := Check(id)
	if err != nil {
		return errors.New("not found id in that table")
	}
	_, err = TX.Exec("DELETE FROM caregivers WHERE ID=?", id)
	if err != nil {
		return errors.New("error while deleting caregiver")
	} else if errCommit := TX.Commit(); errCommit != nil {
		log.Println("Error")
	}
	return nil
}