package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/vadim8q258475/user-microservice/pb"
	repo "github.com/vadim8q258475/user-microservice/repo"

	"go.uber.org/zap"
)

const timeoutMS = 300 * time.Millisecond

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo repo.UserRepository
	log  *zap.Logger
}

func NewUserService(repo repo.UserRepository, logger *zap.Logger) *UserService {
	return &UserService{repo: repo, log: logger}
}

func (s *UserService) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	listCtx, cancel := context.WithTimeout(ctx, timeoutMS)
	defer cancel()
	models, err := s.repo.List(listCtx)
	s.log.Info("Geting users")
	if err != nil {
		s.log.Error(fmt.Sprintf("Failed to get users; error: %s", err.Error()))
		return nil, err
	}
	response := &pb.ListResponse{Users: make([]*pb.User, len(models))}

	for i, model := range models {
		response.Users[i] = UserModelToGetResponse(model)
	}
	s.log.Info("Users get successfuly")
	return response, nil
}

func (s *UserService) GetByEmail(ctx context.Context, request *pb.GetReuqest) (*pb.User, error) {
	getCtx, cancel := context.WithTimeout(ctx, timeoutMS)
	defer cancel()
	model, err := s.repo.GetByEmail(getCtx, request.Email)
	s.log.Info("Geting user by email")
	if err != nil {
		s.log.Error(fmt.Sprintf("Failed to get user by email; error: %s", err.Error()))
		return nil, err
	}
	response := UserModelToGetResponse(model)

	s.log.Info("User by email get successfuly")
	return response, nil
}

func (s *UserService) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	model := UserCreateRequestToModel(request)
	crtCtx, cancel := context.WithTimeout(ctx, timeoutMS)
	defer cancel()

	s.log.Info(fmt.Sprintf("Creating user; email: %s, password: %s", request.Email, request.Password))
	err := s.repo.Create(crtCtx, model)
	if err != nil {
		s.log.Error(fmt.Sprintf("Failed to create user; error: %s", err.Error()))
		return nil, err
	}

	s.log.Info("User created successfuly")
	response := GetCreateResponse("User created successfuly")
	return response, nil
}
