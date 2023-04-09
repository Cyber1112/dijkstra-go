package createPlace

type InputCreatePlace struct {
	Name      string  `json:"name" validate:"required"`
	Latitude  float32 `json:"latitude" validate:"required"`
	Longitude float32 `json:"longitude" validate:"required"`
}
