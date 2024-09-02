package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)

	GetAllUsers() ([]User, error)

	UpdateUserByID(id uint, user User) (User, error)

	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {

	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
func (r *userRepository) GetAllUsers() ([]User, error) {

	var users []User
	err := r.db.Find(&users).Error
	return users, err
}
func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	user.ID = id //TODO: nuuuu hz need refactoring????
	result := r.db.Save(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
func (r *userRepository) DeleteUserByID(id uint) error {
	var user User
	user.ID = id
	result := r.db.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
