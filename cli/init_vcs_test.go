package cli_test

import (
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

	if _, err := os.Stat(".gitignore"); os.IsNotExist(err) {
		t.Fatal("GCC")
	}
}
