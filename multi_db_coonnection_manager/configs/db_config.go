package configs

import (
	"fmt"
	"log"
	"multidbmanager/dto"

	"gorm.io/gorm"
)

type DBConnection *gorm.DB

// DatabaseConnection is an interface that represents a database connection.
type DatabaseConnection interface {
	Connect() (DBConnection, error)
}

// DBConfig represents the configuration for a database.
type DBConfig struct {
	IdentificationName string // IdentificationName using for get the specific database connection.
	DB                 string
	User               string
	Password           string `json:"_"`
	Host               string
	Port               string
	Type               string // Type of the database ("mysql", "postgres", "mssql" or etc...)
	SSLMode            string
	TimeZone           string
	dialector          gorm.Dialector
}

// Connect establishes a database connection based on the provided configuration.
func (config *DBConfig) Connect() (DBConnection, error) {
	db, err := gorm.Open(config.dialector, &gorm.Config{})
	return db, err
}

// NewConnection creates a new DatabaseConnection based on the given config.
func (config *DBConfig) NewConnection() (DBConnection, error) {
	var dbConnection DatabaseConnection
	switch config.Type {
	case "mysql":
		dbConnection = &MySQLConnection{Config: config}
	case "postgres":
		dbConnection = &PostgresConnection{Config: config}
	default:
		return nil, fmt.Errorf("Unsupported database type: %s", config.Type)
	}

	// create new connection
	con, err := dbConnection.Connect()
	if err != nil {
		return nil, err
	}

	//AutoMigrate dtos
	err = con.Statement.AutoMigrate(&dto.User{})
	if err != nil {
		return nil, err
	}

	return con, nil
}

func closeDBConnection(con DBConnection) error {
	sql, err := con.Statement.DB.DB()
	if err != nil {
		return err
	}

	sql.Close()
	return nil
}

func CloseDBConnections() {
	for _, con := range databaseConnections {
		err := closeDBConnection(con)
		if err != nil {
			log.Print(err)
		}
	}
}

// database connections
var databaseConnections map[string]DBConnection

func InitDBConnections() {
	// Initialize database connections
	databaseConnections = make(map[string]DBConnection)
	// Define your database configurations
	dbConfigs := []DBConfig{
		{
			IdentificationName: "TEST_MYSQL_CON",
			DB:                 "testapp",
			User:               "root",
			Password:           "root",
			Host:               "localhost",
			Port:               "3306",
			TimeZone:           "Local",
			Type:               "mysql",
		},
		{
			IdentificationName: "TEST_POSTGRES_CON",
			Host:               "localhost",
			User:               "postgres",
			Password:           "postgres",
			Port:               "5432",
			DB:                 "testapp",
			SSLMode:            "disable",
			TimeZone:           "Asia/Shanghai",
			Type:               "postgres",
		},
	}

	// Connect to each database and store the connection in the map
	for _, config := range dbConfigs {
		db, err := config.NewConnection()
		if err != nil {
			log.Fatalf("Failed to connect to %s database: %v", config.DB, err)
		}
		databaseConnections[config.IdentificationName] = db
	}

}

func GetDBConnection(identificationName string) DBConnection {
	con, ok := databaseConnections[identificationName]
	if !ok {
		log.Fatalf("%s database connection not found", identificationName)
		return nil
	}
	return con
}
