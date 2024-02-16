package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type Client struct {
	client *redis.Client
}

func NewClient(address string) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,  
	})

	if err := client.Ping().Err(); err != nil {
		return nil, fmt.Errorf("failed to ping Redis server at %s: %v", address, err)
	}

	return &Client{client: client}, nil
}

func (c *Client) Get(key string) (string, error) {
	value, err := c.client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil // rreturn nil for "key not found" error
		}
		log.Printf("Error retrieving value for key '%s' from Redis: %v\n", key, err)
		return "", fmt.Errorf("failed to get value for key '%s' from Redis: %v", key, err)
	}
	return value, nil
}

func (c *Client) Set(key, value string, expiration time.Duration) error {
	err := c.client.Set(key, value, expiration).Err()
	if err != nil {
		log.Printf("Error setting value for key '%s' in Redis: %v\n", key, err)
		return fmt.Errorf("failed to set value for key '%s' in Redis: %v", key, err)
	}
	return nil
}
