package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	CacheCapacity int
	GlobalExpiry  int
	Port          int
	RedisAddress  string
}

func LoadConfig() (*Config, error) {
	cacheCapacityStr := os.Getenv("CACHE_CAPACITY")
	cacheCapacity, err := strconv.Atoi(cacheCapacityStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse CACHE_CAPACITY: %v", err)
	}

	globalExpiryStr := os.Getenv("GLOBAL_EXPIRY")
	globalExpiry, err := strconv.Atoi(globalExpiryStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse GLOBAL_EXPIRY: %v", err)
	}

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PORT: %v", err)
	}

	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		return nil, fmt.Errorf("REDIS_ADDRESS is not set")
	}

	return &Config{
		CacheCapacity: cacheCapacity,
		GlobalExpiry:  globalExpiry,
		Port:          port,
		RedisAddress:  redisAddress,
	}, nil
}
