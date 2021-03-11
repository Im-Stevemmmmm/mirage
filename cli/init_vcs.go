package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Im-Stevemmmmm/mirage/database"
	"github.com/Im-Stevemmmmm/mirage/vcs"
)

const (
	gitignoreWriteErrMsg = "Something went wrong while writing to the .gitignore file!"
)

// InitVCS initializes the Mirage VCS system
func InitVCS(data *InitVCSData) {
	cs := *data.ConnectionString
	if len(cs) == 0 {
		fmt.Println("No value for connection-string specified")
		return
	}

	dbEngine := *data.DBEngine
	if len(dbEngine) == 0 {
		fmt.Println("No value for db-engine specified")
		return
	}

	h := database.Handlers[dbEngine]
	if h == nil {
		fmt.Printf("The database engine %s is not supported\n", h)
		return
	}

	// Begin initialization
	os.Mkdir(vcs.RootDir, 0755)
	appendGitignore()

	os.Mkdir(vcs.LocalDir, 0755)
	os.Create(vcs.ConfigPath)
	os.Create(vcs.StatePath)

	json, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(vcs.ConfigPath, json, 0644)

	if err := h.Clone(*data.ConnectionString); err != nil {
		panic(err)
	}

	vcs.CreateBranch("master")
	vcs.CheckoutBranch("master")
}

func appendGitignore() {
	flags := os.O_APPEND | os.O_WRONLY

	gitignore, err := os.OpenFile(".gitignore", flags, 0644)
	defer gitignore.Close()

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

// InitVCSData is the data for initializing Mirage
type InitVCSData struct {
	DBEngine         *string
	ConnectionString *string
}
