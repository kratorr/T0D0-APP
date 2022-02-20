package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	"todo/pkg/config"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Server struct {
	httpServer *http.Server
}

// TODO вынести модели на уровень модуля
// TODO создать структуру ответа\ошибки

func main() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	// config, err :=
	config.InitConfig()
	fmt.Println(viper.ConfigFileUsed())
	fmt.Println(viper.GetString("postgres_host"))

	// if err != nil {
	//		zap.L().Sugar().Error("config error", err.Error())/
	//	}

	db, err := repository.NewPostgresDB(repository.PostgresConfig{
		Host:     viper.GetString("postgres_host"),
		Port:     viper.GetInt("postgres_port"),
		Username: viper.GetString("postgres_username"),
		Password: viper.GetString("postgres_password"),
		DBname:   viper.GetString("postgres_dbname"),
		SSLMode:  viper.GetString("postgres_sslmode"),
	})
	if err != nil {
		zap.L().Sugar().Fatal(err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	g := initRoutes(handlers)
	server := new(Server)

	go func() {
		if err := server.Run(8085, g); err != nil {
			zap.L().Sugar().Fatal("Error start server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	zap.L().Sugar().Info("Shutdown server")
	if err := server.Shutdown(context.Background()); err != nil {
		zap.L().Sugar().Error("Error stop server: %s", err.Error())
	}
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func initRoutes(routes *handler.Handler) *gin.Engine {
	g := gin.New()
	g.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	routes.InitRoutes(g)
	return g
}

func (s *Server) Run(port int, handler http.Handler) error {
	portStr := strconv.Itoa(port)

	s.httpServer = &http.Server{
		Addr:           ":" + portStr,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
