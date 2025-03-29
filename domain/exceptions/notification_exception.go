package exceptions

import "errors"

var (
	ErrorInvalidStaffNotificationStatus = errors.New("invalid status provided")
)