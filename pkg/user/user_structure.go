package user_pkg

type IUserUseCase interface {
	NewUser(*INewUserParams) (*INewUserResponse, error)
}

type INewUserParams struct {
	Nome            string `json:"nome"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type INewUserResponse struct {
	Id    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
