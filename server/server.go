package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// Server struct
type Server struct {
	MySQLClient *sqlx.DB
	router      *gin.Engine
}

// NewServer creates and returns a new Server instance
func NewServer() *Server {
	srv := &Server{router: gin.New(), MySQLClient: MySQLConnect()}

	//srv.ConfigureRoutes()
	return srv
}

// PUT ROUTES HERE

//func (srv *Server) ConfigureRoutes() {
//}

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
