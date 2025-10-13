package service

import (
	"Main/internal/domain"
	"Main/internal/ports"
	pb "Main/proto"
	"context"
	"log"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo ports.UserRepository // <-- Наша зависимость от порта
}

// NewUserServer - конструктор для нашего сервиса
func NewUserServer(repo ports.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Получен запрос на создание юзера: email=%s\n", req.GetEmail())

	// Создаем доменную модель из запроса
	user := &domain.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(), // Пока без хеширования
	}

	// Используем порт для сохранения
	err := s.repo.Save(ctx, user)
	if err != nil {
		log.Printf("Ошибка сохранения пользователя: %v", err)
		return nil, err // В будущем здесь будут красивые gRPC ошибки
	}

	log.Printf("Пользователь успешно сохранен с ID: %s", user.ID)

	// Возвращаем ID из обновленной модели
	return &pb.CreateUserResponse{UserId: user.ID}, nil
}

// ... метод Login пока оставляем без изменений ...
func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// ...
	return &pb.LoginResponse{Message: "Логин пока не реализован"}, nil
}
