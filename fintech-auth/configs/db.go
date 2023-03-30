package configs

import (
	"fintechGo/internal/types"
	"log"

	"github.com/knadh/koanf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(conf *koanf.Koanf) *gorm.DB {

	dsn := conf.String("database.host")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error %v", err)
	}

	db.AutoMigrate(
		&types.AuthUser{},
	)
	return db

}
