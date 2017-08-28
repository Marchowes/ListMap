package users

import (
	"net/url"

	"github.com/dyninc/qstring"
)

// User Struct
type User struct {
	Username  string
	Email     string
	Password  string
	Level     string
	Created   string
	LastLogin string
}

// QueryParameters is a struct for the supported Users query parameters.
type QueryParameters struct {
	Username string `qstring:"username,omitempty"`
	Email    string `qstring:"email,omitempty"`
	Password string `qstring:"password,omitempty"`
	Level    string `qstring:"level,omitempty"`
}

// NewQueryParams returns a new QueryParameters object
func NewQueryParams(u *url.URL) (*QueryParameters, error) {
	params := &QueryParameters{}
	query := u.Query()
	err := qstring.Unmarshal(query, params)
	if err != nil {
		return nil, err
	}
	return params, nil
}
