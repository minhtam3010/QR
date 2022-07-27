package usecase

import (
	"fmt"
	"log"

	"github.com/minhtam3010/qr/app/config"
	"github.com/minhtam3010/qr/app/models"
)

type Enroll struct {
	User     models.User     `json:"user"`
	Guardian models.Guardian `json:"guardian"`
}

func (e *Enroll) CreateEnroll() (Enroll, error) {
	gs := models.Guardian_Student{}
	user, err := e.User.CreateUser()
	tx := config.GetTx()
	fmt.Println(err)
	if err != nil {
		// rollback
		_ = tx.Rollback()
		log.Printf("Error1: %v", err)
		return Enroll{}, err
	}

	guardian, err := e.Guardian.CreateGuardian()
	fmt.Println(err)
	if err != nil {
		// rollback
		_ = tx.Rollback()
		log.Printf("Error1: %v", err)
		return Enroll{}, err
	}

	if errComit := models.TX.Commit(); errComit != nil {
		fmt.Println("Error")
	} else {
		gs.CreateGS(user, guardian) //tx this shit
	}

	return *e, nil
}
