package util

import "errors"

var (
	ErrNameToLong  = errors.New("Name to long")
	ErrNameToShort = errors.New("Name to short, min 3")
	ErrPassWrong   = errors.New("Password incorrect")
	ErrNoSlots     = errors.New("No slots available")
	ErrNoPlayer    = errors.New("No players left")
	ErrToManySlots = errors.New("To many slots")
	ErrToLessSlots = errors.New("To less slots, min 2")
)
