package routers

import (
	"fmt"
	"net/http"

	"github.com/Marchowes/ListMap/dbmysql"
	"github.com/Marchowes/ListMap/listmapper"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// GetUsers gets users
func GetUsers(c *gin.Context) {
	log := c.MustGet(listmapper.LoggingKey).(*logrus.Entry)
	//db := c.MustGet(listmapper.MySQLClientKey).(*dbmysql.MySQLClient)
	//params, err := users.NewQueryParams(c.Request.URL)
	//if err != nil {
	//	log.WithError(err).Info("Invalid Query Parameters")
	//	Abort(c, InvalidQueryParameters)
	//	return
	//}
	log.Info("hi")

}

// GetUser gets users
func GetUser(c *gin.Context) {
	log := c.MustGet(listmapper.LoggingKey).(*logrus.Entry)
	db := c.MustGet(listmapper.MySQLClientKey).(*dbmysql.MySQLClient)
	username := c.Param("username")
	user, prob := db.GetUser(username, log)
	if prob != nil {
		Abort(c, prob)
		return
	}
	c.JSON(http.StatusOK, &user)
}

// PutUsers puts users
func PutUsers(c *gin.Context) {
	fmt.Println("holla")
}
