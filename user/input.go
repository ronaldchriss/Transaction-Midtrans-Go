package user

type RegisterUserInput struct {
	Name         string `json:"name" binding:"required"`
	Occupation   string `json:"occupation" binding:"required"`
	PasswordHash string `json:"password" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
}

type LoginInput struct {
	Email    string `json: "email" binding: "required,email"`
	Password string `json: "password" binding:"required"`
}

type CheckEmail struct {
	Email string `json: "email" binding: "required,email"`
}
