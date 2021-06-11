package services

import "book_list/entities"

type UserService interface {
	SignUpUser(inputSignUp inputs.UserInputSignUp) (entities.User, error)
	SignInUser(inputSignIn inputs.UserInputSignIn) entities.User
}
