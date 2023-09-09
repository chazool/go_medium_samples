package configs

import (
	"fmt"
	"modulemanager/dto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbCon *gorm.DB

type DB struct {
	_        struct{}
	Host     string
	User     string
	Password string `json:"_"`
	Port     string
	DB       string
	SSLMode  string
	TimeZone string
}

type DBOperator struct {
	Con *gorm.DB
}

func (db *DB) NewConnect() (*gorm.DB, error) {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	dialector := postgres.Open(fmt.Sprintf(dsn, db.Host, db.User, db.Password, db.DB, db.Port, db.SSLMode, db.TimeZone))
	con, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	con.AutoMigrate(&dto.User{})

	return con, nil
}

func (db *DBOperator) Save(value interface{}) error {
	return db.Con.Create(value).Error
}

func (db *DBOperator) Find(dest interface{}) error {
	return db.Con.Find(dest).Error
}

func InitDBConnection() {
	dbConfig := DB{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Port:     "5432",
		DB:       "testapp",
		SSLMode:  "disable",
		TimeZone: "Asia/Shanghai",
	}
	dbcon, err := dbConfig.NewConnect()
	if err != nil {
		return
	}
	dbCon = dbcon
}

func GetDBConnection() *gorm.DB {
	return dbCon
}
