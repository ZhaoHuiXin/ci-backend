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
	WandInstance.DB = db
	WandInstance.RDS = rds
	return WandInstance
}
