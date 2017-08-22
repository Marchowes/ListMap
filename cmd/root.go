package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Marchowes/ListMap/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultMySQLUser     = "listmapper"
	defaultMySQLPassword = "password"
	defaultMySQLDB       = "listmap"
	defaultMySQLHost     = "localhost:3306"
	defaultHost          = "0.0.0.0"
	defaultPort          = 9000
)

var (
	host                    string
	port                    int
	mySQLHost               string
	mySQLUser               string
	database                string
	password                string
	readTimeout             time.Duration
	defaultMySQLReadTimeout = 5 * time.Second
)

// RootCmd blah
var RootCmd = &cobra.Command{
	Use:   "ListMapper [OPTIONS]",
	Short: "ListMapper API",
	Long:  `Yep, ListMapper API`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.NewServer()
		srv.MySQLClient = server.MySQLConnect()
		srv.Run()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(
		&host,
		"host",
		"",
		defaultHost,
		"Specify host to run the server on")
	viper.BindPFlag("host", RootCmd.PersistentFlags().Lookup("host"))
	RootCmd.PersistentFlags().IntVarP(
		&port,
		"port",
		"p",
		defaultPort,
		"Set the port to bind the HTTP server to")
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))

	RootCmd.PersistentFlags().StringVarP(
		&mySQLHost,
		"mysql-host",
		"",
		defaultMySQLHost,
		"Mysql host to connect to")
	viper.BindPFlag("mysql.host", RootCmd.PersistentFlags().Lookup("mysql-host"))
	RootCmd.PersistentFlags().StringVarP(
		&mySQLUser,
		"mysql-user",
		"",
		defaultMySQLUser,
		"User to connect to the mysql database as")
	viper.BindPFlag("mysql.user", RootCmd.PersistentFlags().Lookup("mysql-user"))
	RootCmd.PersistentFlags().StringVarP(
		&database,
		"mysql-database",
		"",
		defaultMySQLDB,
		"Name of the mysql database to connect to")
	viper.BindPFlag("mysql.database", RootCmd.PersistentFlags().Lookup("mysql-database"))
	RootCmd.PersistentFlags().StringVarP(
		&password,
		"mysql-password",
		"",
		defaultMySQLPassword,
		"Password for authenticating to the mysql database")
	viper.BindPFlag("mysql.password", RootCmd.PersistentFlags().Lookup("mysql-password"))
	RootCmd.PersistentFlags().DurationVarP(
		&readTimeout,
		"mysql-readtimeout",
		"",
		defaultMySQLReadTimeout,
		"Read connection timeout for the mysql database connection")
	viper.BindPFlag("mysql.readtimeout", RootCmd.PersistentFlags().Lookup("mysql-readtimeout"))

}

// Execute blah.
func Execute() {
	RootCmd.Execute()
}

func initConfig() {
	viper.SetEnvPrefix("listmapper")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetConfigName("listmapper")
	viper.AddConfigPath("/etc/listmapper")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
