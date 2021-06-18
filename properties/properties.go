package properties

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/uniplaces/carbon"
	"time"
)

var (
	redisHost = ""
	redisPort = ""
)

type genericResponse struct {
	Message      string      `json:"message,omitempty"` //nolint:govet
	Time         string      `json:"time,omitempty"`
	ResponseCode string      `json:"response_code,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func contextRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})
	return client
}

func WriteRedis(key string, body string, expiry time.Duration) (string, error) {
	ctx := context.Background()
	client := contextRedis()
	defer client.Close()
	return client.Set(ctx, key, body, expiry).Result()
}

func ReadRedis(key string) (string, error) {
	ctx := context.Background()
	resultMessage, err := contextRedis().Get(ctx, key).Result()
	defer contextRedis().Close()
	return resultMessage, err
}

func PurgeRedis(key string) error {
	ctx := context.Background()
	err := contextRedis().Del(ctx, key)
	defer contextRedis().Close()
	return err.Err()
}

func NResponse(message string, responseCode string, data interface{}) genericResponse {
	return genericResponse{
		Message:      message,
		Data:         data,
		Time:         carbon.Now().DateTimeString(),
		ResponseCode: responseCode,
	}
}
