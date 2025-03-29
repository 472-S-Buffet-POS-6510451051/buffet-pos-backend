package exceptions

import "errors"

var (
	ErrInvoiceNotFound = errors.New("invoice not found")
	ErrTotalFoodWeightInvalid = errors.New("total food weight must be positive")
)
