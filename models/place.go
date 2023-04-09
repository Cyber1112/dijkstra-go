package models

type Place struct {
	ID        uint `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string
	Longitude float32
	Latitude  float32
}
