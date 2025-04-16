package service

import (
	"strconv"

	pb "github.com/vadim8q258475/user-microservice/pb"
	repo "github.com/vadim8q258475/user-microservice/repo"
)

func UserModelToGetResponse(model repo.UserModel) *pb.User {
	return &pb.User{
		Id:          strconv.Itoa(int(model.ID)),
		CreatedData: model.CreatedAt.String(),
		Email:       model.Email,
		Password:    model.Password,
	}
}

func UserCreateRequestToModel(request *pb.CreateRequest) repo.UserModel {
	return repo.UserModel{
		Email:    request.Email,
		Password: request.Password,
	}
}

func GetCreateResponse(query string) *pb.CreateResponse {
	return &pb.CreateResponse{Query: query}
}
