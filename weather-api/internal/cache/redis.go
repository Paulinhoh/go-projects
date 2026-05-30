package cache

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS"), // vem do .env
		Password: "",                 // se tiver senha
		DB:       0,                  // banco padrão
	})
)

func Set(city string, weather []byte) error {
	if err := client.Set(context.Background(), "weather:"+city, weather, 2*time.Minute).Err(); err != nil {
		return errors.New("Failed to set value in the redis instance: " + err.Error())
	}
	return nil
}

func Get(city string) (string, error) {
	val, err := client.Get(context.Background(), "weather:"+city).Result()
	if err == redis.Nil {
		return "", err
	} else if err != nil {
		return "", errors.New("Failed to connect in the redis instance: " + err.Error())
	}

	return val, nil
}
