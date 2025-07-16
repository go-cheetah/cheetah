package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
	conn *gorm.DB
}

func (db *SQLiteDB) Connect(connectionString string) error {
	var err error
	db.conn, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	return err
}

func (db *SQLiteDB) Close() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *SQLiteDB) GetDB() *gorm.DB {
	return db.conn
}
