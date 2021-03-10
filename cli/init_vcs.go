package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/Im-Stevemmmmm/mirage/database"
)

const (
	gitignoreWriteErrMsg = "Something went wrong while writing to the .gitignore file!"
)

func parseFlags() initVCSData {
	dbEngine := flag.String("db-engine", "", "The database type. MySQL or PostgreSQL are supported.")
	connectionString := flag.String("connection-string", "", "The connection string for the database to be tracked.")

	flag.Parse()

	return initVCSData{
		DBEngine:         *dbEngine,
		ConnectionString: *connectionString,
	}
}

// InitVCS initializes the Mirage VCS system
func InitVCS() {
	data := parseFlags()
	if len(data.ConnectionString) == 0 {
		fmt.Println("No value for connection-string specified")
		return
	}
	if len(data.DBEngine) == 0 {
		fmt.Println("No value for db-engine specified")
		return
	}

	h := database.Handlers[data.DBEngine]
	if h == nil {
		fmt.Printf("The database engine %s is not supported\n", h)
		return
	}

	// Begin initialization
	os.Mkdir(".mirage", 0755)
	appendGitignore()

	os.Chdir(".mirage")
	os.Mkdir("local", 0755)
	os.Create("settings.json")

	json, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile("settings.json", json, 0644)

	h.Clone(data.ConnectionString)
}

func appendGitignore() {
	flags := os.O_APPEND | os.O_WRONLY

	gitignore, err := os.OpenFile(".gitignore", flags, 0644)

	var exists bool
	if errors.Is(err, os.ErrNotExist) {
		exists = false

		gitignore, err = os.OpenFile(".gitignore", flags|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println(gitignoreWriteErrMsg)
			return
		}
	} else {
		exists = true
	}

	var lb string
	if exists {
		lb = "\n"
	} else {
		lb = ""
	}

	s := fmt.Sprintf("%s%s", lb, "# Mirage directories\n.mirage/local\n")
	if _, err := gitignore.WriteString(s); err != nil {
		fmt.Println(gitignoreWriteErrMsg)
	}
}

type initVCSData struct {
	DBEngine         string
	ConnectionString string
}
