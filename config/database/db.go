package database

import (
	"fmt"
	"log"
	"mygram-api/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {
	var (
		env      = os.Getenv("ENV")
		host     = os.Getenv("PGHOST")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDBNAME")
		port     = os.Getenv("PGPORT")
		timeZone = os.Getenv("TIMEZONE")
		dsn      = ""
		db       *gorm.DB
		err      error
	)

	if env == "production" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=%s", host, user, password, dbname, port, timeZone)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbname, port, timeZone)
	}

	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{FullSaveAssociations: true}); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err = db.AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{}); err != nil {
		log.Fatal("Error migrating database: ", err.Error())
	}

	return db
}
