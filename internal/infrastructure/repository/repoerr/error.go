package repoerr

import "errors"

// DB
var (
	// foreign key violation: 23503
	FKViolation = "23503"
	// unique violation: 23505
	UniqueConstraint = "23505"
)

// Logo
var (
	ErrLogoNotFound   = errors.New("logo not found")
	ErrLogoExist      = errors.New("logo url must be unique")
	ErrLogoDependency = errors.New("cannot delete object, it is linked to adv")
)

// Color
var (
	ErrColorNotFound        = errors.New("color not found")
	ErrColorHexExist        = errors.New("colors hex must be unique")
	ErrColorDependencyExist = errors.New("cannot delete adv because there is record reference to adv")
)

// TG
var (
	ErrTGNotFound        = errors.New("tg not found")
	ErrTGExist           = errors.New("tg url must be unique")
	ErrTGDependencyExist = errors.New("cannot delete tg because there is record reference to adv")
)

// Adv
var (
	ErrAdvDependencyNotExist = errors.New("cannot add or update adv because there is no record reference to color, logo or tg")
	ErrAdvNotFound           = errors.New("adv not found")
)