package constants

import "errors"

var ErrNotFound = errors.New("Quotes not found")
var ErrIDNotFound = errors.New("Quotes didn't finded with the same ID")
var ErrMustBePositive = errors.New("Number can't be zero")