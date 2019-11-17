package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-redis/redis"

	"github.com/elDante/homebudget/config"
	"github.com/elDante/homebudget/database"
)

func RedisConnector(config *config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		DB:   config.Database,
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Fatal("Can't connect to Redis server")
	}
	return client
}

func main() {
	configPath := flag.String("config", "config.toml", "Path to TOML config")
	flag.Parse()

	conf := config.Parse(configPath)
	db := database.Connector(&conf.Database)
	defer db.Close()
	database.MigrateDB(db)
}
