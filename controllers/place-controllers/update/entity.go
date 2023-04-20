package updatePlace

import "time"

type InputUpdatePlace struct {
	ID        uint    `validate:"required"`
	Weight    float32 `json:"weight" validate:"required"`
	UpdatedAt time.Time
}
