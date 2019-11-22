package contrib

import (
	"fmt"
	"log"

	"github.com/elDante/homebudget/config"
	"github.com/go-redis/redis"
)

// RedisConnector connect and return redis.Client
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
