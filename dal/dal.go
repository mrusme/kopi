package dal

import (
	"database/sql"
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	sqlite "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

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
	embeddedMigrations, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return err
	}
	driver, err := sqlite.WithInstance(dal.db, &sqlite.Config{})
	m, err := migrate.NewWithInstance(
		"iofs",
		embeddedMigrations,
		"sqlite",
		driver,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func New(dbString string) *DAL {
	var err error

	dal := new(DAL)

	if dal.db, err = sql.Open("sqlite", dbString); err != nil {
		log.Fatal(fmt.Errorf("failed to open database connection: %w", err))
	}
	if err := dal.db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("failed to ping database: %w", err))
	}

	return dal
}

func Open(dbFile string, devMode bool) (*DAL, error) {
	var mode string = "rwc"
	if devMode {
		mode = "memory"
	}

	dbString := fmt.Sprintf(
		"file:%s?cache=shared&mode=%s&_pragma=foreign_keys(1)",
		dbFile,
		mode,
	)
	db := New(dbString)
	err := db.Init()
	if err != nil {
		return nil, err
	}

	return db, nil
}
