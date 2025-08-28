// core package holds project configurationa and shared functionailities
package core

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB(config *Config) {
	dsn := "host=" + config.DBHost +
		" user=" + config.DBUser +
		" password=" + config.DBPassword +
		" dbname=" + config.DBName +
		" port=" + config.DBPort +
		" sslmode=" + config.DBSSLMode

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect to database: ", err)
	}
	log.Println("Database connected successfully!")
}
