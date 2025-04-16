package app

import (
	"errors"
	"fmt"
	"net"

	pb "github.com/vadim8q258475/user-microservice/pb"
	service "github.com/vadim8q258475/user-microservice/service"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	service    *service.UserService
	grpcServer *grpc.Server
	port       string
	logger     *zap.Logger
}

func NewApp(
	service *service.UserService,
	grpcServer *grpc.Server,
	logger *zap.Logger,
	port string,
) *App {
	return &App{
		service:    service,
		grpcServer: grpcServer,
		logger:     logger,
		port:       port,
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		return err
	}
	pb.RegisterUserServiceServer(a.grpcServer, a.service)
	a.logger.Info(fmt.Sprintf("grpc server start on port %s", a.port))
	if err := a.grpcServer.Serve(l); err != nil {
		return errors.New("failed to serve")
	}
	return nil
}
