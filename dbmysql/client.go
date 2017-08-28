package dbmysql

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	maxIdleConnections = 20
	maxOpenConnections = 20
)

// MySQLClient struct.
type MySQLClient struct {
	db *sqlx.DB
}

// NewMySQLClient returns a mysql client
func NewMySQLClient(user, pass, host, name string, readTimeout time.Duration) (*MySQLClient, error) {
	connectionConfig := mysql.Config{User: user,
		Passwd:          pass,
		Net:             "tcp",
		Addr:            host,
		DBName:          name,
		MultiStatements: true,
		ReadTimeout:     readTimeout,
	}

	db, err := sqlx.Connect("mysql", connectionConfig.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)

	return &MySQLClient{db: db}, nil
}
