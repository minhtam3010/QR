package models

import "testing"

func TestCreateEnroll(t *testing.T) {
	enroll := Enroll{
		User{
			ID:            10,
			EntityCode:    2,
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
		}, Guardian{
			ID:            4,
			Fullname:      "Le Minh Tam Parent",
			Email:         "parentmt@gmail.com",
			Address:       "None",
			BOD:           "30/10/1984",
			Phone:         "0000",
			Qualification: "None",
			Role:          "Mother",
			DateCreated:   1658678118,
			DateUpdated:   0,
		},
	}

	res, err := enroll.CreateEnroll()
	if err != nil {
		panic(err)
	}

	if res == enroll {
		t.Log("PASSED")
	}
	if res.User == enroll.User {
		t.Log("PASSED")
	}
	if res.Guardian == enroll.Guardian {
		t.Log("PASSED")
	}
	if res != enroll || res.User != enroll.User || res.Guardian != enroll.Guardian{
		t.Error("ERROR")
	} 
}
