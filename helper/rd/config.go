package rd

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASS"),
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ping)
}
