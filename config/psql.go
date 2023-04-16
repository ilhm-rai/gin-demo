package config

import (
	"context"
	"fmt"
	"time"

	"github.com/ilhm-rai/go-middleware/entity"
	"github.com/ilhm-rai/go-middleware/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func NewPostgresDatabase(configuration Config) *gorm.DB {
	ctx, cancel := NewPostgresContext()
	defer cancel()

	host := configuration.Get("PSQL_HOST")
	user := configuration.Get("PSQL_USER")
	password := configuration.Get("PSQL_PASS")
	dbname := configuration.Get("PSQL_DB")
	port := configuration.Get("PSQL_PORT")
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	exception.PanicIfNeeded(err)

	DB.Debug().Exec(`
	DO $$ BEGIN
		CREATE TYPE role_gm AS ENUM ('ADMIN', 'USER');
	EXCEPTION
		WHEN duplicate_object THEN null;
	END $$;`)
	DB.Debug().AutoMigrate(entity.User{}, entity.Product{})

	DB.WithContext(ctx)
	return DB
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func GetDB() *gorm.DB {
	return DB
}
