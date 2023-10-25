package interfaces

import (
	"anote/internal/errors"
)

type ESClient interface {
	Search(index string, query string) ([]any, *errors.AppError)
}
