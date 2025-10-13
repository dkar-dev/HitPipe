package main

import (
	pb "Main/proto"
	"github.com/joho/godotenv"
	"os"

	"Main/internal/adapters/postgres"
	"Main/internal/config"
	"Main/internal/service"
	"fmt"
	"log"
	"net"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func main() {
	// Loading .env file
	// Checking that the program is running in production mode
	// Depends on the "ENV" variable
	// Take values: "production", "development", "preprod"
	env := os.Getenv("ENV")
	if env != "production" && env != "preprod" {
		err := godotenv.Load() // Loading .env from root directory
		if err != nil {
			log.Fatalf("error loading .env file: %v", err)
		}
	}

	// Loading configuration from yaml file
	// The loader and its Load() method is implemented in loader.go
	cfg, err := config.Load[config.Config]("config.yaml")
	if err != nil {
		log.Fatalf("error loading config.yaml file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SSLMode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalf("cannot connect to the database: %v", err)
	}
	defer db.Close()
	log.Println("")

	// 3. Создание зависимостей (порты и адаптеры)
	userRepository := postgres.NewUserRepository(db)

	// 4. Создание сервиса с зависимостями
	userServer := service.NewUserServer(userRepository)

	// 5. Запуск gRPC сервера
	lis, err := net.Listen("tcp", cfg.GRPCServer.Port)
	if err != nil {
		log.Fatalf("Не удалось слушать порт: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userServer) // <-- Регистрируем наш обновленный сервис
	log.Printf("gRPC сервер слушает порт %s\n", cfg.GRPCServer.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
