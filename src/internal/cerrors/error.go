package cerrors

import (
	"errors"
)

var (
	ErrAliasAlreadySaved       = errors.New("alias for provided url is already saved")
	ErrAliasForURLDoesNotExist = errors.New("alias for provided url does not exist")
	ErrAliasDoesNotExist       = errors.New("alias does not exist")
	ErrInvalidUrl              = errors.New("invalid url")
)
