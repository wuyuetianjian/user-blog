package data

import (
	"qizhan/backend/internal/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(c *conf.Bootstrap) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(c.MySQL.DSN), &gorm.Config{})
}
