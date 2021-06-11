package repositories

import (
	"book_list/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(userData entities.User) (entities.User, error)
	UpdateUser(userData entities.User) (entities.User, error)
	FindAllUsers() ([]entities.User, error)
	FindUserById(userId int) (entities.User, error)
	FindUserByEmail(userEmail string) (entities.User, error)
}

type userRepository struct {
	userDb *gorm.DB
}

func NewUserRepository(userDb *gorm.DB) *userRepository {
	return &userRepository{userDb}
}

func (r *userRepository) InsertUser(userData entities.User) (entities.User, error) {
	if err := r.userDb.Create(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}

func (r *userRepository) UpdateUser(userData entities.User) (entities.User, error) {

}
