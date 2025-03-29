package responses

import "github.com/google/uuid"

type BaseCustomer struct {
	ID    uuid.UUID `json:"id"`
	Phone string    `json:"phone"`
	Point int       `json:"point"`
}