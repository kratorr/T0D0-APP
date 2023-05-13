package main

import (
	"context"
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
	// swagger embed files
)

type Server struct {
	httpServer *http.Server
}

// TODO вынести модели на уровень модуля
// TODO создать структуру ответа\ошибки

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8005
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	logger, _ := zap.NewDevelopment()

	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	// config, err :=
	config.InitConfig()

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
	// efer db.Close() потом разобраться с закрытием подключения
	if err != nil {
		zap.L().Sugar().Fatal(err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, viper.GetString("secret_key"))
	handlers := handler.NewHandler(services, viper.GetString("secret_key"))

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
