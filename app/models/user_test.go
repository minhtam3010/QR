package models

import (
	"fmt"
	"testing"
)

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

	result := GetUserById(1, "Le Minh Tam", "minhtam")
	if result != user {
		t.Error("Error :(((")
	} else {
		t.Log("PASSED")
	}
	fmt.Println("Minh Tam")
}
