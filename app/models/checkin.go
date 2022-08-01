package models

// import (
// 	"time"

// 	"github.com/minhtam3010/qr/app/config"
// )

// // CI stand for checkin
// type CheckIn struct {
// 	TimeCI      int64  `json:"timeci"`
// 	UnicodeID   int    `json:"unicodeid"`
// 	CaregiverID int    `json:"caregiverid"`
// 	Reason      string `json:"reason"`
// }

// var reason = "Go to school"

// func (ci *CheckIn) CreateCheckIn() (CheckIn, error) {
// 	TX = config.GetTx()

// 	_, err := TX.Exec(`INSERT INTO checkins(TimeCI, UnicodeID, CaregiverID, Reason)`,
// 						time.Now(), )
// }
