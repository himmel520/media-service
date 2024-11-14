package repoerr

import (
	"errors"
)

// DB
const (
	// foreign key violation: 23503
	FKViolation = "23503"
	// unique violation: 23505
	UniqueConstraint = "23505"
)

var (
	ErrImageNotFound   = errors.New("image not found")
	ErrImageExist      = errors.New("image url must be unique")
	ErrImageDependencyExist = errors.New("cannot delete object, it is linked to adv")
)

// Ошибки Color
var (
	ErrColorNotFound        = errors.New("color not found")
	ErrColorHexExist        = errors.New("colors hex must be unique")
	ErrColorDependencyExist = errors.New("cannot delete adv because there is record reference to adv")
)

// Ошибки TG
var (
	ErrTGNotFound        = errors.New("tg not found")
	ErrTGExist           = errors.New("tg url must be unique")
	ErrTGDependencyExist = errors.New("cannot delete tg because there is record reference to adv")
)

// Ошибки Adv
var (
	ErrAdvDependencyNotExist = errors.New("cannot add or update adv because there is no record reference to color, Image or tg")
	ErrAdvNotFound           = errors.New("adv not found")
)
