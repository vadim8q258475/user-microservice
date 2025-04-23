package service

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	pb "github.com/vadim8q258475/user-microservice/pb"
	"github.com/vadim8q258475/user-microservice/repo"

	"go.uber.org/zap"
)

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repo.NewMockUserRepository(ctrl)
	logger := zap.NewNop()
	service := NewUserService(mockRepo, logger)

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	request := pb.CreateRequest{Email: "email", Password: "Pass"}
	result, err := service.Create(context.Background(), &request)

	assert.NoError(t, err)
	assert.Equal(t, result.Query, "User created successfuly")

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(assert.AnError).Times(1)

	_, err = service.Create(context.Background(), &request)

	assert.Error(t, err)
}

func TestUserService_GetByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repo.NewMockUserRepository(ctrl)
	logger := zap.NewNop()
	service := NewUserService(mockRepo, logger)

	testUser := repo.UserModel{Email: "user1@test.com"}

	mockRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(testUser, nil).Times(1)

	request := pb.GetReuqest{Email: "user1@test.com"}
	result, err := service.GetByEmail(context.Background(), &request)

	assert.NoError(t, err)
	assert.Equal(t, result.Email, testUser.Email)

	mockRepo.EXPECT().GetByEmail(gomock.Any(), gomock.Any()).Return(repo.UserModel{}, assert.AnError).Times(1)

	_, err = service.GetByEmail(context.Background(), &request)

	assert.Error(t, err)
}

func TestUserService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repo.NewMockUserRepository(ctrl)
	logger := zap.NewNop()
	service := NewUserService(mockRepo, logger)

	testUsers := []repo.UserModel{
		{Email: "user1@test.com"},
		{Email: "user2@test.com"},
	}

	mockRepo.EXPECT().
		List(gomock.Any()).
		Return(testUsers, nil).
		Times(1)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	resp, err := service.List(ctx, &pb.ListRequest{})

	assert.NoError(t, err)
	assert.Equal(t, len(testUsers), len(resp.Users))
	assert.Equal(t, testUsers[0].Email, resp.Users[0].Email)

	mockRepo.EXPECT().
		List(gomock.Any()).
		Return(nil, assert.AnError).
		Times(1)

	_, err = service.List(context.Background(), &pb.ListRequest{})
	assert.Error(t, err)
}
