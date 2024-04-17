package repository

import (
	"gorm-terminal/internal/domain/model"

	"gorm.io/gorm"
)



type UserRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) *UserRepository {//NewUserはUserRepositoryのポインタを返す
	return &UserRepository{db}
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


func (r *UserRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(user *model.User,id uint) error {
	if err := r.db.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id uint) error {
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
