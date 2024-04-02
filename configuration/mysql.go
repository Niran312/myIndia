package configuration

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"strconv"
)

type MysqlData struct {
	Host     string `json:"DB_HOST"`
	Port     int    `json:"DB_PORT"`
	Name     string `json:"DB_NAME"`
	User     string `json:"DB_USER"`
	Password string `json:"DB_PASSWORD"`
}

func (s MysqlData) MysqlConnectionInfo() string {

	info := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true", s.User, s.Password, s.Host, s.Port, s.Name)

	log.Infof("DB connect string: ", info)
	return info
}

func GetMsqlData() MysqlData {

	log.Infof("ENV values: %v", os.Getenv("DB_PORT"))

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	log.Infof("Port number: ", port)

	return MysqlData{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}
