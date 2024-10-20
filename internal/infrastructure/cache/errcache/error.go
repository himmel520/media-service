package errcache

import "errors"

// CACHE
var (
	ErrKeyNotFound = errors.New("there is no current data")
)
