package database

var (
	// Handlers is all of the supported databases
	Handlers = map[string]dbHandler{
		"postgres": postgres{},
	}
)

type dbHandler interface {
	Clone(cs string)
}
