package postgres

import (
	"Online-Test-GO/internal/pkg/config"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"io/ioutil"
	"log"
	"strings"
)

func NewPostgres() *bun.DB {
	dsn := "postgres://" + config.GetConf().DBUsername + ":" + config.GetConf().DBPassword + "@" +
		config.GetConf().DBHost + ":" + config.GetConf().DBPort + "/" + config.GetConf().DBName +
		"?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	if err := runMigrations(db); err != nil {
		log.Printf("Error running migrations: %v", err)
	}

	return db
}

func runMigrations(db *bun.DB) error {
	migrationFile := "internal/pkg/script/migrations.sql"

	log.Printf("Running migration file: %s", migrationFile)
	content, err := ioutil.ReadFile(migrationFile)
	if err != nil {
		return err
	}

	queries := strings.Split(string(content), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, err = db.Exec(query)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			return err
		}
	}

	log.Printf("Migrations completed successfully")
	return nil
}
