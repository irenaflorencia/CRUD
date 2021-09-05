package repository

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection(DSN string) (*gorm.DB, error) {
	db, err :gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
		return db, nil
	}
