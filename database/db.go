package database

import (
	"fmt"
	"log"

	"github.com/ilhm-rai/go-middleware/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	port     = "5432"
	dbname   = "simple-api"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	fmt.Println("successfully connected to database")
	db.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE role_gm AS ENUM ('ADMIN', 'USER');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;`)
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
