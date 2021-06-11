package inputs

type UserInputSignUp struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInputSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputUserId struct {
	Id int `uri:"id" binding:"required"`
}

type UpdateUserInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}
