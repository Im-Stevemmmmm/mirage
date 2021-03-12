package database

var (
	// Handlers maps a database name to its implementation.
	Handlers = map[string]handler{
		"postgres": postgres{},
	}
)

type handler interface {
	Clone(cs string) error
	ApplyMigration(m Migration) error
}

// Migration is a database migration. Apply is the sql statments to be executed
// within the database to apply the migration. Revert is the sql statments
// required to undo the migration.
type Migration struct {
	Apply  []string
	Revert []string
}
