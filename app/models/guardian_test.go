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
	
	result1, _ := GetGuardiansByID(1, "Le Minh Tam Parent")
	result2, _ := GetGuardians()
	if result1 != guardian && result2[0] != guardian{
		t.Error("Error :((((")
	} else {
		t.Log("PASSED")
	}
}

// func TestCreateGuardian(t *testing.T) {
// 	guardian := Guardian{
// 		ID:            2,
// 		Fullname:      "Le Minh Tam Parent",
// 		Email:         "parentmt@gmail.com",
// 		Address:       "None",
// 		BOD:           "30/10/1984",
// 		Phone:         "0000",
// 		Qualification: "None",
// 		Role:          "Mother",
// 		DateCreated:   1658678118,
// 		DateUpdated:   0,
// 	}

// 	res, _ := guardian.CreateGuardian()
// 	if res != guardian {
// 		t.Error("Error :(((")
// 	} else {
// 		t.Log("PASSED")
// 	}
// }

func TestUpdate(t *testing.T) {
	guardian := Guardian{
		ID:            2,
		Fullname:      "Le Minh Tam222 Parent",
		Email:         "parentmt@gmail.com",
		Address:       "None",
		BOD:           "30/10/1984",
		Phone:         "0000",
		Qualification: "None",
		Role:          "Mother",
	}

	res, _ := guardian.UpdateGuardian(2)
	if res != guardian {
		t.Error("Error :(((")
	} else {
		t.Log("PASSED")
	}
}

func TestDelete(t *testing.T) {
	guardian := Guardian{
		ID:            2,
		Fullname:      "Le Minh Tam222 Parent",
		Email:         "parentmt@gmail.com",
		Address:       "None",
		BOD:           "30/10/1984",
		Phone:         "0000",
		Qualification: "None",
		Role:          "Mother",
	}

	err := DeleteGuardian(2)
	if err != nil {
		panic(err)
	}

	res, _ := GetGuardiansByID(2)
	if res != guardian {
		t.Log("PASSED")
	}else {
		t.Error("ERROR :(((")
	}
}