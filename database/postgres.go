package database

import (
	"database/sql"
	"os"
	"os/exec"

	// Postgres driver
	_ "github.com/lib/pq"
)

func (p postgres) Clone(cs string) error {
	db, err := sql.Open("postgres", cs)
	if err != nil {
		return err
	}
	defer db.Close()

	var currentUser string
	dataUser := db.QueryRow("SELECT current_user")
	dataUser.Scan(&currentUser)

	var currentDB string
	dataDB := db.QueryRow("SELECT current_database()")
	dataDB.Scan(&currentDB)

	dumpFile, err := os.OpenFile(".mirage/master.pgsql", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer dumpFile.Close()

	cmd := exec.Command("pg_dump", "-U", currentUser, currentDB)
	dump, err := cmd.Output()
	if err != nil {
		return err
	}

	if _, err := dumpFile.WriteString(string(dump)); err != nil {
		return err
	}

	return nil
}

func (p postgres) ApplyMigration(m Migration) error {
	return nil
}

type postgres struct{}
