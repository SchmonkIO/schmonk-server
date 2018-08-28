package util

import "errors"

var (
	ErrNameToLong = errors.New("Name to long")
	ErrPassWrong  = errors.New("Password incorrect")
	ErrNoSlots    = errors.New("No slots available")
	ErrNoPlayer   = errors.New("No players left")
)
