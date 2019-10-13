package utils

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/prometheus/common/log"
)

var once sync.Once

func GetDB(dbType, connectAddr string) (db *gorm.DB, err error) {
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	once.Do(func() {
		db, err = gorm.Open(dbType, connectAddr)
		if err != nil {
			log.Fatalf("Failed to open DB %s, error: %v", dbType, err)
			return
		}
	})

	return db, nil
}

func GetRDS(redisURL string) (redisClient *redis.Client, err error) {
	// opt, err := redis.ParseURL("redis://:qwerty@localhost:6379/1")
	once.Do(func() {
		optConfig, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatalf("Failed to open RDS, error: %v", err)
			return
		}
		redisClient := redis.NewClient(optConfig)
		pong, err := redisClient.Ping().Result()
		if err != nil || pong != "PONG" {
			log.Fatalf("Redis init failed, %s", err)
			return
		}
	})
	return
}
