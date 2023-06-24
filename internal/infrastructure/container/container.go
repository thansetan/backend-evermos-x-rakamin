package container

import (
	"final_project/internal/helper"
	"final_project/internal/infrastructure/mysql"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var v *viper.Viper

const currentfilepath = "internal/infrastructure/container/container.go"

type (
	Container struct {
		Mysqldb *gorm.DB
		Apps    *Apps
	}

	Apps struct {
		Name      string `mapstructure:"APP_NAME"`
		Host      string `mapstructure:"HOST"`
		Version   string `mapstructure:"VERSION"`
		Address   string `mapstructure:"ADDRESS"`
		HttpPort  int    `mapstructure:"HTTP_PORT"`
		SecretJwt string `mapstructure:"JWT_SECRET"`
	}
)

func loadEnv() {
	projectDirName := "backend-evermos-x-rakamin"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	v.SetConfigFile(string(rootPath) + `/.env`)
}

func init() {
	v = viper.New()

	v.AutomaticEnv()
	loadEnv()

	path, err := os.Executable()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("os.Executable panic : %s", err.Error()))
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed read config : %s", err.Error()))
	}

	err = v.ReadInConfig()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed init config : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed read configuration file")
}

func AppsInit(v *viper.Viper) (apps Apps) {
	err := v.Unmarshal(&apps)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprint("Error when unmarshal configuration file : ", err.Error()))
	}
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed when unmarshal configuration file")
	return
}

func InitContainer() (cont *Container) {
	apps := AppsInit(v)
	mysqldb := mysql.DatabaseInit(v)

	return &Container{
		Apps:    &apps,
		Mysqldb: mysqldb,
	}

}
