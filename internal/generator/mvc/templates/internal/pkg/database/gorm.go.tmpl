package database

import (
	"gorm.io/gorm"
)

type DB interface {
	Connect(connectionString string) error
	Close() error
	GetDB() *gorm.DB
}

type Options struct {
	Migrate bool
}

type Option interface {
	apply(*Options)
}

func New(dbType, connectionString string, opts ...Option) DB {
	var database DB

	switch dbType {
	case "sqlite":
		database = &SQLiteDB{}
	case "mysql":
		database = &MySQLDB{}
	default:
		return nil // errors.New("unsupported database type")
	}

	err := database.Connect(connectionString)
	if err != nil {
		return nil //, err
	}

	options := Options{
		Migrate: true,
	}
	for _, o := range opts {
		o.apply(&options)
	}

	return database //, nil
}
