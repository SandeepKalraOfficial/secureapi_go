package models

import "time"

// Person model
type Person struct {
	ID    string    `json:"id"`
	FName string    `json:"fname"`
	LName string    `json:"lname"`
	DOB   time.Time `json:"dob"`
}
