package ports

type Entity interface {
	GetFieldAdresses() []any
}

type DBConnection interface {
	QueryOne(dest Entity, query string, args ...any) error
	QueryMultiple(dest []Entity, query string, args ...any) error
	Exec(query string, args ...any) error
}
