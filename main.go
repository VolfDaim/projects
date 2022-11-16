package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"net/http"
	"projects/handlers"
	"projects/repository"
	service2 "projects/service"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (server *Server) Run(port string, handler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return server.httpServer.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}

func main() {
	if err := InitConfig(); err != nil {
		panic(err)
	}
	repos := repository.NewRepository()
	service := service2.NewService(repos)
	handler := handlers.NewHandler(service)

	server := new(Server)
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		panic(err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
