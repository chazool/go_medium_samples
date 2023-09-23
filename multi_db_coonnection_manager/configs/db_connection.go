package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
)

// MySQLConnection implements DatabaseConnection for MySQL.
type MySQLConnection struct {
	Config *DBConfig
}

// Connect connects to a MySQL database and returns a GORM DB instance.
func (m *MySQLConnection) Connect() (DBConnection, error) {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s"
	m.Config.dialector = mysql.Open(fmt.Sprintf(dsn, m.Config.User, m.Config.Password, m.Config.Host, m.Config.Port, m.Config.DB, m.Config.TimeZone))
	db, err := m.Config.Connect()
	return db, err
}

// PostgresConnection implements DatabaseConnection for PostgreSQL.
type PostgresConnection struct {
	Config *DBConfig
}

// Connect connects to a PostgreSQL database and returns a GORM DB instance.
func (p *PostgresConnection) Connect() (DBConnection, error) {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	p.Config.dialector = postgres.Open(fmt.Sprintf(dsn, p.Config.Host, p.Config.User, p.Config.Password, p.Config.DB, p.Config.Port, p.Config.SSLMode, p.Config.TimeZone))
	db, err := p.Config.Connect()
	return db, err
}
