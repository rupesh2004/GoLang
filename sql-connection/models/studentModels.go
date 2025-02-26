package models

type Student struct {
	Name      string `json:"name" binding:"required"`
	PRN       int    `json:"prn" binding:"required"`
	MobileNo  string `json:"mobileNo" binding:"required"`
	BirthDate string `json:"birthDate" binding:"required"`
}
