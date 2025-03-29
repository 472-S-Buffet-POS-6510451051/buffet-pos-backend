package exceptions

import "errors"

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrDuplicatedPhone 	= errors.New("phone number is already exist")
	ErrNotEnoughPoints 	= errors.New("not enough points, must have 10 point")
	ErrInvalidPoint 	= errors.New("point must be greater than 0")
	ErrIncorrectPIN     = errors.New("incorrect pin")
	ErrPointLimit       = errors.New("point is full")
)
