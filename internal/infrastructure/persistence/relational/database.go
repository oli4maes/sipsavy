package relational

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
)

func MigrateDb() error {
	connString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		panic("connection string env variable not set")
	}
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Ingredient{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(Cocktail{})
	if err != nil {
		return err
	}

	return nil
}
