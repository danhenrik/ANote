package interfaces

import (
	"anote/internal/errors"
	"reflect"
)

type DBConnection interface {
	QueryOne(t reflect.Type, query string, args ...any) (any, *errors.AppError)
	QueryMultiple(t reflect.Type, query string, args ...any) (any, *errors.AppError)
	Exec(query string, args ...any) *errors.AppError
}
