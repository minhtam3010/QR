package models

import (
	"testing"
)

func TestGuardian(t *testing.T) {
	guardian := Guardian{
		ID:            1,
		Fullname:      "Le Minh Tam Parent",
		Email:         "parentmt@gmail.com",
		Address:       "None",
		BOD:           "30/10/1984",
		Phone:         "0000",
		Qualification: "None",
		Role:          "Mother",
		DateCreated:   1658678118,
		DateUpdated:   0,
	}
	
	result := GetGuardiansByID(1)

	if result != guardian {
		t.Error("Error :((((")
	} else {
		t.Log("PASSED")
	}
}
