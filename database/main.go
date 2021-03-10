package database

var (
	// DBHandlers is all of the supported databases
	DBHandlers = map[string]dbHandler{
		"postgres": postgres{},
	}
)

type dbHandler interface {
	clone()
}
