package updatePlace

import "time"

type InputUpdatePlace struct {
	ID        uint    `validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Latitude  float32 `json:"latitude" validate:"required"`
	Longitude float32 `json:"longitude" validate:"required"`
	UpdatedAt time.Time
}
