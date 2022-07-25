package models

import (
	"log"

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

func GetGuardians() []Guardian {
	db := config.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM guardians ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	guardian := Guardian{}
	res := []Guardian{}
	for rows.Next() {

		err = rows.Scan(&id, &fullname, &email, &address, &bod, &phone, &qualification, &role, &datecreated, &dateupdated)
		if err != nil {
			panic(err.Error())
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
	return res
}

func GetGuardiansByID(id int, params ...string) Guardian {
	db := config.GetDB()
	defer db.Close()

	switch {
	case len(params) == 1:
		fullname = params[0]
	case len(params) == 2:
		fullname = params[0]
		email = params[1]
	case len(params) == 3:
		log.Println("Error!!!")
	}

	rows, err := db.Query("SELECT * FROM guardians WHERE id=? OR fullname=? OR email=?", id, fullname, email)
	if err != nil {
		panic(err.Error())
	}

	guardian := Guardian{}
	for rows.Next() {
		err := rows.Scan(&id, &fullname, &email, &address, &bod, &phone, &qualification, &role, &datecreated, &dateupdated)

		if err != nil {
			switch {
			case dateupdated.Unix() <= 0:
				guardian.DateUpdated = 0
			case dateupdated.Unix() > 0:
				guardian.DateUpdated = dateupdated.Unix()
			default:
				panic(err.Error())
			}
			// if dateupdated.Unix() < 0{

			// }else{
			// 	panic(err.Error())
			// }
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
	return guardian
}
