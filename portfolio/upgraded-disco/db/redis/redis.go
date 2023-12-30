package main

import (
	"context"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

// TODO: use hexagonal arch
// TODO: use interface based approach for backend
func main() {
	fmt.Print(1)
	opt, err := redis.ParseURL("redis://test@localhost:6379/")
	//opt, err := redis.ParseURL(config.FetchConfig().RedisURI)
	if err != nil {
		log.Info().Err(fmt.Errorf("error while connecting to Redis: %v", err))
	}
	fmt.Print(2, opt)
	client := redis.NewClient(opt)

	ctx := context.Background()
	fmt.Println(11)
	err = client.Set(ctx, "foo", "bar", 0).Err()
	fmt.Print(3)
	if err != nil {
		log.Info().Err(fmt.Errorf("error while setting up redis client: %v", err))
	}

	val, err := client.Get(ctx, "foo").Result()
	fmt.Print(4)
	if err != nil {
		log.Info().Err(fmt.Errorf("error while getting data for key %v: %v", "foo", err))
	}
	fmt.Println("foo", val)

	session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
	for k, v := range session {
		fmt.Print(5)
		err := client.HSet(ctx, "user-session:123", k, v).Err()
		fmt.Print(6)
		if err != nil {
			log.Info().Err(fmt.Errorf("error while setting data: %v", err))
		}
	}

	userSession := client.HGetAll(ctx, "user-session:123").Val()
	fmt.Println(userSession)
}
