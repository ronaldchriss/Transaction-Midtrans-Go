package user

type RegisterUserInput struct {
	Name         string
	Occupation   string
	PasswordHash string
	Email        string
}
