package services

import (
	"book_list/entities"
	"book_list/inputs"
	"book_list/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUpUser(inputSignUp inputs.UserInputSignUp) (entities.User, error)
	SignInUser(inputSignIn inputs.UserInputSignIn) (entities.User, error)
	UpdateUserById(inputId inputs.InputUserId, inputData inputs.UpdateUserInput) (entities.User, error)
	GetAllUsers() ([]formatters.UserDataFormatter, error)
	GetUserById(userId inputs.InputUserId) (entities.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) SignUpUser(inputSignUp inputs.UserInputSignUp) (entities.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputSignUp.Password), bcrypt.MinCost)
	userData := entities.User{
		Name: inputSignUp.Name,
		Email: inputSignUp.Email,
		PasswordHash: string(passwordHash),
	}

	if err != nil {
		return userData, err
	}

	newUser, ere
}