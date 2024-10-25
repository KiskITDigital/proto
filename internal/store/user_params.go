package store

type UserCreateParams struct {
	OrganizationID int
	Email          string
	Phone          string
	PasswordHash   string
	TOTPSalt       string
	FirstName      string
	LastName       string
	MiddleName     string
	AvatarURL      string
	EmailVerified  bool
	Role           int
}

type UserGetParams struct {
	Email string
	ID    int
}
