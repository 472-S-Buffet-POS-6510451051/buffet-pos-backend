package requests

type EditPricePerPerson struct {
	Price float64 `json:"price" validate:"required"`
}

type EditPointLimit struct {
	LimitPoint int `json:"limit_point" validate:"required"`
}

type EditUsePointPerPerson struct {
	UsePointPerPerson int `json:"use_point_per_person" validate:"required"`
}

type EditPriceFeeFoodOverWeight struct {
	PriceFeeFoodOverWeight float64 `json:"price_fee_food_overweight" validate:"required"`
}
