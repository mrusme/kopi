package dal

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattn/go-sqlite3"
)

var ErrNotFound = errors.New("record not found")

type DAL struct {
	db *sql.DB
}

func (dal *DAL) DB() *sql.DB {
	return dal.db
}

func (dal *DAL) Close() {
	_ = dal.db.Close()
}

func (dal *DAL) Init() error {
	driver, err := sqlite3.WithInstance(dal.db, &sqlite3.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3",
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func New() *DAL {
	var err error

	dal := new(DAL)

	dbString := fmt.Sprintf("file:test.db?cache=shared&mode=memory")

	if dal.db, err = sql.Open("sqlite3", dbString); err != nil {
		log.Fatal(fmt.Errorf("failed to open database connection: %w", err))
	}
	if err := dal.db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("failed to ping database: %w", err))
	}

	return dal
}
