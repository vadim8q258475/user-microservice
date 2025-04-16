package repo

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email    string
	Password string
}

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&UserModel{},
	)

	return db, err
}
