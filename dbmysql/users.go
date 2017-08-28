package dbmysql

import (
	"fmt"

	"github.com/Marchowes/ListMap/listmapper"
	"github.com/Sirupsen/logrus"
	"github.com/moogar0880/problems"
)

var (
	foo = "bar"
)

// UserQueryParams object, holds user data.
type UserQueryParams struct {
	Username string `db:"user_name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Level    string `db:"user_level"`
}

// User Struct
type User struct {
	Username string `db:"user_name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Level    string `db:"user_level"`
	//Created   string `db:"cust_id"`
	//LastLogin string `db:"cust_id"`
}

// NewUserQueryParams returns a populated UserQueryParams object
func NewUserQueryParams(username, email, password, level string) *UserQueryParams {

	return &UserQueryParams{Username: username, Email: email, Password: password, Level: level}
}

// GetUsers gets users from mysql DB.
func (db *MySQLClient) GetUsers(params *UserQueryParams) {
	// stuff here

}

// GetUser gets single from mysql DB.
func (db *MySQLClient) GetUser(username string, log *logrus.Entry) (*User, problems.StatusProblem) {
	user := User{}
	query := fmt.Sprintf("SELECT user_name, email, password, user_level FROM users where USER_NAME = '%s';", username)
	err := db.db.Get(&user, query)
	if err != nil {
		log.WithError(err).Error("MySQL: Problem Getting User")
		return nil, listmapper.InternalServerError
	}
	return &user, nil

}
