package cmd

import (
	"github.com/elDante/homebudget/config"
	"github.com/elDante/homebudget/contrib"
	"github.com/elDante/homebudget/database"
	"github.com/elDante/homebudget/router"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run http server",
	Long:  `This subcommand run http server`,
	Run: func(cmd *cobra.Command, args []string) {
		cnf := config.Parse(&ConfigPath)
		redis := contrib.RedisConnector(&cnf.Redis)
		db := database.Connector(&cnf.Database)
		router := router.Router(db, redis, &cnf.Site)
		router.Run()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
