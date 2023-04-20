package models

type Place struct {
	ID     uint   `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Start  string `gorm:"size:255;"`
	End    string `gorm:"size:255;"`
	Weight float32
}
