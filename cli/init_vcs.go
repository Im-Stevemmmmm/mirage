package cli

import (
	"errors"
	"flag"
	"os"
)

func parseFlags() (*initVCSData, error) {
	dbEngine := *flag.String("db-engine", "", "The database type. MySQL or PostgreSQL are supported.")
	endpointURL := *flag.String("endpoint-url", "", "The connection string for the database to be tracked.")

	flag.Parse()

	if len(dbEngine) == 0 {
		return nil, errors.New("the flag db-engine was not specified")
	}

	return &initVCSData{
		DBEngine:    dbEngine,
		EndpointURL: endpointURL,
	}, nil
}

func cloneDB(data *initVCSData) {}

// InitVCS initializes the Mirage VCS system
func InitVCS() {
	data, err := parseFlags()
	if err != nil {
		panic(err)
	}

	os.Mkdir(".mirage", 0755)

	cloneDB(data)
}

type initVCSData struct {
	DBEngine    string
	EndpointURL string
}
