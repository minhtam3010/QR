package models

import (
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	user := User{
		ID:            2,
		EntityCode:    2,
		Username:      "quynhnhu",
		Fullname:      "Nguyen Giang Quynh Nhu",
		Password:      "888",
		Email:         "nhu@gmail.com",
		Address:       "None",
		BOD:           "20/10/2002",
		Phone:         "9999",
		Qualification: "None",
		Slogan:        "None",
		Role:          "CB",
		Hobby:         "None",
		DateCreated:   1666198800,
		DateUpdated:   1667062800,
	}

	result, err := GetUsers()
	if user != result[1] || err != nil {
		t.Error("Error :(((")
	} else {
		t.Log("PASSED")
	}
}

func TestUser(t *testing.T) {
	user := User{
		ID:            1,
		EntityCode:    1,
		Username:      "minhtam",
		Fullname:      "Le Minh Tam",
		Password:      "999",
		Email:         "tam@gmail.com",
		Address:       "None",
		BOD:           "30/10/2002",
		Phone:         "9999",
		Qualification: "None",
		Slogan:        "None",
		Role:          "Dev",
		Hobby:         "None",
		DateCreated:   1658665531,
		DateUpdated:   1658665531,
	}

	result, err := GetUserById(1, "Le Minh Tam")
	if result != user || err != nil {
		t.Error("Error :(((")
	} else {
		t.Log("PASSED")
	}
}

// func TestCreateUser(t *testing.T) {
// 	user := User{
// 		ID:            5,
// 		EntityCode:    5,
// 		Username:      "minhtam",
// 		Fullname:      "Le Minh Tam3",
// 		Password:      "999",
// 		Email:         "tam@gmail.com",
// 		Address:       "None",
// 		BOD:           "30/10/2002",
// 		Phone:         "9999",
// 		Qualification: "None",
// 		Slogan:        "None",
// 		Role:          "Dev",
// 		Hobby:         "None",
// 		DateCreated:   1658665531,
// 		DateUpdated:   0,
// 	}

// 	u, err := user.CreateUser()
// 	if err != nil {
// 		panic(err)
// 	}

// 	result, _ := GetUserById(5)
// 	log.Println(u.ID, result.ID)
// 	if result.ID != u.ID {
// 		t.Error("Error :((((")
// 	} else {
// 		t.Log("PASSED")
// 	}
// }

func TestDeleteUser(t *testing.T) {
	user := User{
		ID:            5,
		EntityCode:    5,
		Username:      "minhtam",
		Fullname:      "Le Minh Tam3",
		Password:      "999",
		Email:         "tam@gmail.com",
		Address:       "None",
		BOD:           "30/10/2002",
		Phone:         "9999",
		Qualification: "None",
		Slogan:        "None",
		Role:          "Dev",
		Hobby:         "None",
		DateCreated:   1658665531,
		DateUpdated:   0,
	}

	err := DeleteUser(5)

	log.Println("Delete OK")
	if err != nil {
		panic(err)
	}

	res, _ := GetUserById(5)
	log.Println(res.ID, "Delete")
	if res != user {
		t.Log("PASSED")
	} else {
		t.Error("Error :((((")
	}
}

func TestUpdateUser(t *testing.T) {
	user := User{
		EntityCode:    3,
		Fullname:      "Le Minh Tam",
		Password:      "999",
		Email:         "tam@gmail.com",
		Address:       "None",
		BOD:           "30/10/2002",
		Phone:         "9999",
		Qualification: "None",
		Slogan:        "None",
		Role:          "Dev Backend1",
		Hobby:         "None",
	}
	u, err := user.UpdateUser(7, "minhtam")
	log.Println(u.DateCreated)
	if err != nil {
		panic(err.Error())
	}

	res, _ := GetUserById(7, "minhtam", "Le Minh Tam")
	log.Println(res.DateCreated)
	if res.DateCreated != u.DateCreated {
		t.Error("Error :((((")
	} else {
		t.Log("PASSED")
	}
}
