package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDB struct {
	conn *gorm.DB
}

// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (db *MySQLDB) Connect(connectionString string) error {
	var err error
	db.conn, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	return err
}

func (db *MySQLDB) Close() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *MySQLDB) GetDB() *gorm.DB {
	return db.conn
}

type Migrate bool

func (m Migrate) apply(o *Options) {
	o.Migrate = bool(m)
}

func WithMigrate(c bool) Option {
	return Migrate(c)
}

func migrate() {
}
