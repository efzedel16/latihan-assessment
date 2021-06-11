package services

import (
	"book_list/entities"
	"book_list/inputs"
	"book_list/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUpUser(input inputs.UserInputSignUp) (entities.User, error)
	SignInUser(input inputs.UserInputSignIn) (entities.User, error)
	UpdateUserById(id inputs.InputUserId, input inputs.UpdateUserInput) (entities.User, error)
	// GetAllUsers() ([]formatters.UserDataFormatter, error)
	GetUserById(id inputs.InputUserId) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) SignUpUser(input inputs.UserInputSignUp) (entities.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user := entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(password),
	}

	if err != nil {
		return user, err
	}

	newUser, err := s.userRepository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) SignInUser(input inputs.UserInputSignIn) (entities.User, error) {
	user := entities.User{
		Email:    input.Email,
		Password: input.Password,
	}

	user, err := s.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user
	}
}
