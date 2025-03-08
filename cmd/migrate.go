package cmd

import (
	"fmt"
	"log"

	"michiru/internal/db"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn, err := db.Connect()
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		driver, err := postgres.WithInstance(dbConn.DB, &postgres.Config{})
		if err != nil {
			log.Fatalf("Failed to create migration driver: %v", err)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://db/migrations",
			"postgres", driver)
		if err != nil {
			log.Fatalf("Failed to create migrate instance: %v", err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to run migrations: %v", err)
		}

		fmt.Println("Migrations ran successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
