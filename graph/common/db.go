package common

import (
	"gorm_gql/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	dsn := "host=localhost user=root password=password dbname=simple_bank port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(model.Todo{})
	return db, nil
}
