package cli

import "flag"

// ParseInitFlags parses the flags for initializing Mirage.
func ParseInitFlags() *InitVCSData {
	dbEngine := flag.String("db-engine", "", "The database type. MySQL or PostgreSQL are supported.")
	connectionString := flag.String("connection-string", "", "The connection string for the database to be tracked.")

	flag.Parse()

	return &InitVCSData{
		DBEngine:         dbEngine,
		ConnectionString: connectionString,
	}
}
