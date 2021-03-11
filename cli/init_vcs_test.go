package cli_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Im-Stevemmmmm/mirage/cli"
)

func TestInitVCS(t *testing.T) {
	dbEngine := "postgres"
	cs := "1234"

	cli.InitVCS(&cli.InitVCSData{
		DBEngine:         &dbEngine,
		ConnectionString: &cs,
	})

	if err := pathExists(".gitignore"); err != nil {
		t.Fatal(err)
	}

	if err := pathExists(".mirage/config.json"); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile(".mirage/config.json")
	if err != nil {
		t.Fatal(err)
	}

	var config cli.InitVCSData
	json.Unmarshal(data, &config)

	if *config.ConnectionString != cs {
		t.Fatalf("ConnectionString in the config does not match; got %s expected %s", *config.ConnectionString, cs)
	}

	if *config.DBEngine != dbEngine {
		t.Fatalf("DBEngine in the config does not match; got %s expected %s", *config.DBEngine, dbEngine)
	}

	if err := pathExists(".mirage/master.pgsql"); err != nil {
		t.Fatal(err)
	}

	// Cleanup
	os.RemoveAll(".mirage")
	os.Remove(".gitignore")
}

func pathExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New(path + " was not found")
	}
	return nil
}
