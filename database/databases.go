package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	cfg "myIndia/configuration"
	"strings"
)

var (
	// MysqlDB PgDB is the postgres connection handle
	MysqlDB *gorm.DB
)

func DbInit() *gorm.DB {

	//DB Connection syntax using env keys
	//dbconnection = "DB_USER:DB_PASSWORD@tcp(DB_HOST:DB_PORT)/DB_NAME?parseTime=true"

	dsn := cfg.GetConfig().Mysql.MysqlConnectionInfo()

	log.Infof("dsn values: %v", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Info(strings.Repeat("!", 40))
		log.Info("😏 Could Not Establish Mysql DB Connection")
		log.Info(strings.Repeat("!", 40))
		log.Fatal(err)
	}

	log.Info(strings.Repeat("-", 40))
	log.Info("😀 Connected To Mysql DB")
	log.Info(strings.Repeat("-", 40))

	MysqlDB = db

	return MysqlDB
}
