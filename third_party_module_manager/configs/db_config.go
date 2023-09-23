package configs

import (
	"fmt"
	"modulemanager/dto"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbCon DBConnection

type DBConnection *gorm.DB

// DBConfig defines the database configuration parameters.
type DBConfig struct {
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
	Con DBConnection
}

// NewConnection initializes the database connection based on the provided config.
func (db *DBConfig) NewConnection() (DBConnection, error) {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	dialector := postgres.Open(fmt.Sprintf(dsn, db.Host, db.User, db.Password, db.DB, db.Port, db.SSLMode, db.TimeZone))
	con, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	con.AutoMigrate(&dto.User{})

	return con, nil
}

// Save saves the given value to the database.
func (db *DBOperator) Save(value interface{}) error {
	return db.Con.Statement.Create(value).Error
}

// Find finds and retrieves data from the database and stores it in dest.
func (db *DBOperator) Find(dest interface{}) error {
	return db.Con.Statement.Find(dest).Error
}

// InitDBConnection initializes the database connection.
func InitDBConnection() {
	dbConfig := DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Port:     "5432",
		DB:       "testapp",
		SSLMode:  "disable",
		TimeZone: "Asia/Shanghai",
	}
	con, err := dbConfig.NewConnection()
	if err != nil {
		return
	}
	dbCon = con
}

// GetDBConnection returns the database connection.
func GetDBConnection() DBConnection {
	return dbCon
}
