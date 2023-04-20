package createPlace

type InputCreatePlace struct {
	Start  string  `json:"start" validate:"required"`
	End    string  `json:"end" validate:"required"`
	Weight float32 `json:"weight" validate:"required"`
}
