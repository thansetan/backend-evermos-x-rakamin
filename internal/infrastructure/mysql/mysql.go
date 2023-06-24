package mysql

import (
	"final_project/internal/helper"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConf struct {
	Username           string `mapstructure:"MYSQL_USERNAME"`
	Password           string `mapstructure:"MYSQL_PASSWORD"`
	DbName             string `mapstructure:"MYSQL_DBNAME"`
	Host               string `mapstructure:"MYSQL_HOST"`
	Port               int    `mapstructure:"MYSQL_PORT"`
	Schema             string `mapstructure:"MYSQL_SCHEMA"`
	LogMode            bool   `mapstructure:"MYSQL_LOG_MODE"`
	MaxLifetime        int    `mapstructure:"MYQSL_MAX_LIFETIME"`
	MinIdleConnections int    `mapstructure:"MYSQL_MIN_IDLE_CONNECTION"`
	MaxOpenConnections int    `mapstructure:"MYSQL_MAX_OPEN_CONNECTION"`
}

const currentfilepath = "internal/infrastructure/mysql/mysql.go"

func DatabaseInit(v *viper.Viper) *gorm.DB {
	var mysqlConfig MysqlConf
	err := v.Unmarshal(&mysqlConfig)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed init database mysql : %s", err.Error()))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("Cannot conenct to database : %s", err.Error()))
	}

	_, err = db.DB()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("Cannot conenct to database : %s", err.Error()))
	}

	// TODO POOLING CONNECTION

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "⇨ MySQL status is connected")
	RunMigration(db)

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("Failed to close connection to database : %s", err.Error()))
	}

	dbSQL.Close()

}
