package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tensuqiuwulu/bupda-reset-otp/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection(DB *config.Database) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB.Username,
		DB.Password,
		DB.Address,
		strconv.Itoa(int(DB.Port)),
		DB.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Cannot connect to database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Cannot connect to database: " + err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("Cannot ping the database: " + err.Error())
	}

	fmt.Println("Success connect to database")

	sqlDB.SetMaxIdleConns(int(DB.MaxIdle))
	sqlDB.SetMaxOpenConns(int(DB.MaxOpen))
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(DB.MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Minute * time.Duration(DB.MaxIdleTime))
	return db
}

func Close(DB *gorm.DB) {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("database disconected")
}
