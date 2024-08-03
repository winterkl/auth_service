package app

import (
	"auth/internal/config"
	v1 "auth/internal/controller/gRPC/v1"
	auth_usecase "auth/internal/domain/auth/usecase"
	auth_repo "auth/internal/infrastructure/repo/auth"
	cache_repo "auth/internal/infrastructure/repo/cache"
	"auth/pkg/grpc_server"
	"auth/pkg/jwt_auth"
	"auth/pkg/postgres"
	"auth/pkg/redis_db"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	gRPCServer *grpc_server.GRPCServer
}

func New(cfg *config.Config) *App {

	// Подключение к Redis
	redisDB := redis_db.New(cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.DB, cfg.Redis.Port)

	// Подключение к PSql
	psql, err := postgres.New(cfg.DataBase.User, cfg.DataBase.Password, cfg.DataBase.Host, cfg.DataBase.DB, cfg.DataBase.SSLMode, cfg.DataBase.Port)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %s", err.Error())
	}

	// Инициализация пакета JWT
	jwtAuth := jwt_auth.NewJwtAuth(cfg.JwtAuth.Key)

	// Инициализация Repositories
	cacheRepo := cache_repo.New(redisDB)
	authRepo := auth_repo.New(psql)

	// Инициализация UseCases
	useCases := v1.UseCase{
		Auth: auth_usecase.New(authRepo, cacheRepo, jwtAuth),
	}

	// Инициализация gRPC сервера
	grpcServer := grpc_server.New(cfg.GRPCServer.PORT, cfg.GRPCServer.Host)

	v1.Register(grpcServer, useCases)

	return &App{
		gRPCServer: grpcServer,
	}
}

func (app *App) Run() {

	go app.gRPCServer.Run()

	//Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.gRPCServer.GracefulStop()
}
