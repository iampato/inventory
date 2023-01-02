package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() *gorm.DB {
	// print current path
	if _, err := os.Stat("/go/src/inventory/user/config.ini"); err == nil {
		fmt.Printf("******************* File exists *******************\n")
	} else {
		fmt.Printf("******************* File does not exist *******************\n")
	}
	cfg, err := ini.Load("/go/src/inventory/user/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	dbUrl := cfg.Section("postgresql").Key("host").String()
	dbPort, _ := cfg.Section("postgresql").Key("port").Int64()
	dbUser := cfg.Section("postgresql").Key("user").String()
	dbPassword := cfg.Section("postgresql").Key("password").String()
	dbName := cfg.Section("postgresql").Key("dbname").String()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", dbUrl, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Fail to connect to DB: %v", err)
		os.Exit(1)
	}

	return db
}
