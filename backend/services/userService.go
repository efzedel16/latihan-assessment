package services

import (
	"book_list/entities"
	"book_list/inputs"
	"book_list/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUpUser(input inputs.UserInputSignUp) (entities.User, error)
	SignInUser(input inputs.UserInputSignIn) (entities.User, error)
	UpdateUserById(inputId inputs.InputUserId, input inputs.UpdateUserInput) (entities.User, error)
	// GetAllUsers() ([]formatters.UserDataFormatter, error)
	GetUserById(inputId inputs.InputUserId) (entities.User, error)
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
	userEmail := input.Email
	userPassword := input.Password
	user, err := s.userRepository.FindUserByEmail(userEmail)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found with this email")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPassword)); err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) UpdateUserById(inputId inputs.InputUserId, input inputs.UpdateUserInput) (entities.User, error) {
	user, err := s.userRepository.FindUserById(inputId.Id)
	if err != nil {
		return user, nil
	}

	user.Name = input.Name
	user.Email = input.Email
	if user.Id == 0 {
		return user, errors.New("not the user")
	}

	updateUser, err := s.userRepository.UpdateUser(user)
	if err != nil {
		return updateUser, nil
	}

	return user, nil
}
