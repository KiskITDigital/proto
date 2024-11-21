package store

type UserCreateParams struct {
	Email        string
	Phone        string
	PasswordHash string
	TOTPSalt     string
	FirstName    string
	LastName     string
	MiddleName   string
	AvatarURL    string
}

type UserGetParams struct {
	Email string
	ID    int
}
