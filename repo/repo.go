package repo

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	List(context.Context) ([]UserModel, error)
	GetByEmail(context.Context, string) (UserModel, error)
	Create(context.Context, UserModel) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) List(ctx context.Context) ([]UserModel, error) {
	users := make([]UserModel, 0)
	result := r.db.WithContext(ctx).Find(&users)
	return users, result.Error
}
func (r *userRepository) GetByEmail(ctx context.Context, email string) (UserModel, error) {
	model := UserModel{}
	result := r.db.WithContext(ctx).Where("email = ?", email).Find(&model)
	return model, result.Error
}
func (r *userRepository) Create(ctx context.Context, model UserModel) error {
	result := r.db.WithContext(ctx).Create(&model)
	fmt.Println(result.Error)
	return result.Error
}
