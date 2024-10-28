package repoerr

import (
	"net/http"
)

// DB
const (
	// foreign key violation: 23503
	FKViolation = "23503"
	// unique violation: 23505
	UniqueConstraint = "23505"
)

type RepoError struct {
	message string
	status  int
}

func NewRepoErr(message string, status int) *RepoError {
	return &RepoError{
		message: message,
		status:  status,
	}
}

func (e *RepoError) Error() string {
	return e.message
}

func (e *RepoError) Status() int {
	return e.status
}

// Logo
var (
	ErrLogoNotFound   = NewRepoErr("logo not found", http.StatusNotFound)
	ErrLogoExist      = NewRepoErr("logo url must be unique", http.StatusConflict)
	ErrLogoDependency = NewRepoErr("cannot delete object, it is linked to adv", http.StatusConflict)
)

// Ошибки Color
var (
	ErrColorNotFound        = NewRepoErr("color not found", http.StatusNotFound)
	ErrColorHexExist        = NewRepoErr("colors hex must be unique", http.StatusConflict)
	ErrColorDependencyExist = NewRepoErr("cannot delete adv because there is record reference to adv", http.StatusConflict)
)

// Ошибки TG
var (
	ErrTGNotFound        = NewRepoErr("tg not found", http.StatusNotFound)
	ErrTGExist           = NewRepoErr("tg url must be unique", http.StatusConflict)
	ErrTGDependencyExist = NewRepoErr("cannot delete tg because there is record reference to adv", http.StatusConflict)
)

// Ошибки Adv
var (
	ErrAdvDependencyNotExist = NewRepoErr("cannot add or update adv because there is no record reference to color, logo or tg", http.StatusConflict)
	ErrAdvNotFound           = NewRepoErr("adv not found", http.StatusNotFound)
)
