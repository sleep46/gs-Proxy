package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/sleep46/gs-Proxy/cache"
    "github.com/sleep46/gs-Proxy/redis"
    "github.com/sleep46/gs-Proxy/proxy"
    "github.com/sleep46/gs-Proxy/server"
    "github.com/sleep46/gs-Proxy/config"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    cache := cache.NewLRUCache(cfg.CacheCapacity, cfg.GlobalExpiry)

    redisClient, err := redis.NewClient(cfg.RedisAddress)
    if err != nil {
        log.Fatalf("Error initializing Redis client: %v", err)
    }

    proxy := proxy.NewProxy(cache, redisClient)

    httpServer := server.NewServer(cfg.Port, proxy)

    log.Printf("Server listening on port %d\n", cfg.Port)
    if err := httpServer.ListenAndServe(); err != nil {
        log.Fatalf("Error starting HTTP server: %v", err)
    }
}
