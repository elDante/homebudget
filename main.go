package main

import (
	"flag"

	"github.com/elDante/homebudget/config"
	"github.com/elDante/homebudget/database"
	"github.com/elDante/homebudget/models"
)

func main() {
	configPath := flag.String("config", "config.toml", "Path to TOML config")
	flag.Parse()

	conf := config.Parse(configPath)
	db := database.Connector(&conf.Database)
	defer db.Close()
	database.MigrateDB(db)
	models.CurrencyFixtures(db)
}
