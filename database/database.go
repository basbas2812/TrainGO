package database

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)
