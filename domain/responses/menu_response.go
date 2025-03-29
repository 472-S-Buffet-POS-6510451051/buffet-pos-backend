package responses

import (
	"github.com/google/uuid"
)

type BaseMenu struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	CategoryID  *uuid.UUID `json:"categoryId"`
	ImageURL    *string    `json:"imageUrl"`
	IsAvailable bool       `json:"isAvailable"`
}

type MenuDetail struct {
	BaseMenu
}

type NumberMenu struct {
	BaseMenu
	Number 		int 	`json:"number"`
}
