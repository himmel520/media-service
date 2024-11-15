package cache

import "errors"

var (
	ErrKeyNotFound = errors.New("there is no current data")
)
