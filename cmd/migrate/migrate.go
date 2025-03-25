package main

import (
    "database/sql"
	_ "github.com/mattn/go-sqlite3"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/sqlite3"
    _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/charmbracelet/log"
)

func migrateDB() error {
    db, err := sql.Open("sqlite3", "file:./db/hiss.sqlite3")
	if err != nil {
		return err
	}
    driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

    m, err := migrate.NewWithDatabaseInstance(
        "file://./sql/migrations",
        "hissenburg", driver)
    //err = m.Down() // or m.Step(2) if you want to explicitly set the number of migrations to run
    err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := migrateDB(); err != nil {
		log.Fatal(err)
	}
}
