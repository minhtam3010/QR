package models

import (
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

	result := GetUsers()
	if user != result[1] {
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

	result := GetUserById(1, "Le Minh Tam")
	if result != user {
		t.Error("Error :(((")
	} else {
		t.Log("PASSED")
	}
}
