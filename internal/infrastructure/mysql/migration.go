package mysql

import (
	"final_project/internal/dao"
	"final_project/internal/helper"
	"fmt"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&dao.User{},
		&dao.Address{},
		&dao.Category{},
		&dao.Product{},
		&dao.ProductLog{},
		&dao.ProductPhoto{},
		&dao.Store{},
		&dao.Transaction{},
		&dao.TransactionDetail{},
	)

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Failed Database Migrated : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}
