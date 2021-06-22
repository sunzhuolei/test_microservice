package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var(
	Db  *gorm.DB
	Redis  *redis.Client
)
