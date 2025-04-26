package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"yuyuid.id/internal/model"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (dt *UserRepositoryImpl) Count(ctx context.Context) (int64, error) {
	var total int64
	err := dt.DB.WithContext(ctx).Count(&total).Error
	if err != nil {
		return total, err
	}
	return total, nil
}

func (dt *UserRepositoryImpl) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	if err := dt.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (dt *UserRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (model.User, error) {
	var user model.User
	if err := dt.DB.WithContext(ctx).Where("uuid = ? ", uuid).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (dt *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := dt.DB.WithContext(ctx).Where("email = ? AND deleted_at IS NULL ", email).First(&user).Error
	return user, err
}

func (dt *UserRepositoryImpl) FindById(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	if err := dt.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (dt *UserRepositoryImpl) Save(ctx context.Context, user *model.User) error {
	if err := dt.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (dt *UserRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	if err := dt.DB.WithContext(ctx).Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (dt *UserRepositoryImpl) Delete(ctx context.Context, id uint) error {
	if err := dt.DB.WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (dt *UserRepositoryImpl) IsExist(ctx context.Context, email string) error {
	var count int64
	// Memeriksa apakah ada user dengan email tersebut tanpa menarik data lainnya
	err := dt.DB.WithContext(ctx).Model(&model.User{}).Where("email = ? AND deleted_at IS NULL", email).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("email already registered")
	}
	return nil
}
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}
