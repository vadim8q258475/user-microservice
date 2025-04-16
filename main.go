package main

import (
	"fmt"
	"os"

	"github.com/vadim8q258475/user-microservice/app"
	"github.com/vadim8q258475/user-microservice/repo"
	"github.com/vadim8q258475/user-microservice/service"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	if dir, err := os.Getwd(); err == nil {
		fmt.Println("Current directory:", dir)
	}

	godotenv.Load(".env")

	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	fmt.Println(dsn)

	db, err := repo.InitDB(dsn)
	if err != nil {
		panic(err)
	}

	repo := repo.NewUserRepository(db)

	service := service.NewUserService(repo, logger)

	server := grpc.NewServer()

	app := app.NewApp(service, server, logger, port)

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
