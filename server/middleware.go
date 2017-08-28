package server

import (
	"github.com/Marchowes/ListMap/listmapper"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// LoggingMiddleware injects a logrus logger.
func LoggingMiddleware(c *gin.Context) {

	logger := logrus.WithFields(logrus.Fields{
		"Path": c.Request.URL.EscapedPath(),
	})
	c.Set(listmapper.LoggingKey, logger)

}
