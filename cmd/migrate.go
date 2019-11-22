package cmd

import (
	"github.com/elDante/homebudget/config"
	"github.com/elDante/homebudget/database"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	Long:  `This subcommand run database migrations and upload fixtures`,
	Run: func(cmd *cobra.Command, args []string) {
		cnf := config.Parse(&ConfigPath)
		db := database.Connector(&cnf.Database)
		database.MigrateDB(db)
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}
