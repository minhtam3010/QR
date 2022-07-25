package models

import (
	"log"

	"github.com/minhtam3010/qr/app/config"
)

type Guardian_Student struct {
	GuardianID int `json:"guardianID"`
	StudentID  int `json:"studentID"`
}

func (gs *Guardian_Student) CreateGS() {
	db := config.GetDB()
	defer db.Close()

	create, err := db.Prepare("INSERT INTO guardian_student(GuardianID, StudentID) VALUES(?, ?)")
	if err != nil {
		panic(err)
	}
	create.Exec(gs.GuardianID, gs.StudentID)
	log.Println("INSERT Successfully")
}
