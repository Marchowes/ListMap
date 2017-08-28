package server

import (
	"time"

	"github.com/Marchowes/ListMap/dbmysql"
	"github.com/spf13/viper"
)

// MySQLConnect connects to the MySQL Database and returns a client.
func MySQLConnect() *dbmysql.MySQLClient {
	readTimeout, err := time.ParseDuration(viper.GetString("mysql.readtimeout"))
	if err != nil {
		panic(err.Error())
	}
	client, err := dbmysql.NewMySQLClient(viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.database"),
		readTimeout)
	if err != nil {
		panic(err.Error())
	}
	return client
}
