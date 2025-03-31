package restaurantstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewSQLRestore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
