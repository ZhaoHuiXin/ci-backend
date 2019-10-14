package utils

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type Wand struct {
	DB  *gorm.DB
	RDS *redis.Client
}

var WandInstance *Wand

func NewWand(db *gorm.DB, rds *redis.Client) *Wand {
	w := &Wand{
		DB:db,
		RDS:rds}
	WandInstance = w
	return WandInstance
}
