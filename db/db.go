package db

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbHandler *gorm.DB

func Init() *gorm.DB {
	var err error

	dsn := "host=" + viper.GetString("database.host") + " user=" + viper.GetString("database.user") + " password=" + viper.GetString("database.password") + " dbname=" + viper.GetString("database.name") + " port=" + viper.GetString("database.port") + " sslmode=disable"
	dbHandler, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sql, err := dbHandler.DB()
	if err != nil {
		log.Fatal(err)
	}

	sql.SetMaxIdleConns(49)
	sql.SetMaxOpenConns(50)
	sql.SetConnMaxIdleTime(time.Second * 5)
	sql.SetConnMaxLifetime(time.Minute)

	return dbHandler
}

func GetDB() *gorm.DB {
	return dbHandler
}
