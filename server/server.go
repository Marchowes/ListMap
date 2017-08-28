package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Marchowes/ListMap/dbmysql"
	"github.com/Marchowes/ListMap/listmapper"
	"github.com/Marchowes/ListMap/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Server struct
type Server struct {
	MySQLClient *dbmysql.MySQLClient
	router      *gin.Engine
	UserRoutes  *gin.RouterGroup
}

// NewServer creates and returns a new Server instance
func NewServer() *Server {
	srv := &Server{router: gin.New(), MySQLClient: MySQLConnect()}

	srv.ConfigureRoutes()
	return srv
}

// ConfigureRoutes Configures all the routes.
func (srv *Server) ConfigureRoutes() {
	srv.router.Use(injectMiddleware(srv))
	srv.buildUserRoutes()
}

func injectMiddleware(srv *Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Inject MySQL Client.
		c.Set(listmapper.MySQLClientKey, srv.MySQLClient)
		LoggingMiddleware(c)
	}
}

// Run runs the server
func (srv *Server) Run() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	on := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port"))
	go srv.router.Run(on)
	<-done
}

func (srv *Server) buildUserRoutes() {
	srv.UserRoutes = srv.router.Group("/users")
	{
		srv.UserRoutes.GET("",
			routers.GetUsers)
		srv.UserRoutes.GET("/:username",
			routers.GetUser)
		srv.UserRoutes.PUT("",
			routers.PutUsers)
	}
}
