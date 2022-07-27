package models

import (
	"github.com/minhtam3010/qr/app/config"
)

type Guardian_Student struct {
	GuardianID int `json:"guardianID"`
	StudentID  int `json:"studentID"`
}

func (gs *Guardian_Student) CreateGS(user User, guardian Guardian) error {
	TX := config.GetDB()

	_, err := TX.Exec("INSERT INTO guardian_student(GuardianID, StudentID) VALUES(?, ?)", user.ID, guardian.ID)
	if err != nil {
		panic(err)
	}
	// log.Println("INSERT Successfully")
	return nil
}
