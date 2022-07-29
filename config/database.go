package config

import (
	"fmt"
	"go-checkin/config/credential"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func NewDbMssql() *gorm.DB {

	var (
		dbHost   = credential.DbHost
		port     = credential.DbPort
		user     = credential.DbUsername
		password = credential.DbPassword
		database = credential.DbName
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, dbHost, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	idle, _ := strconv.Atoi(os.Getenv("SET_MAX_IDLE_CONN"))
	open, _ := strconv.Atoi(os.Getenv("SET_MAX_OPEN_CONN"))

	pool, err := db.DB()
	pool.SetMaxIdleConns(idle)
	pool.SetMaxOpenConns(open)

	return db.Debug()
}
