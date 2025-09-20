package auth

type SignUpRequest struct {
	IpAddress    string `json:"-"`
	UserAgent    string `json:"-"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}

type SignUpResponse struct {
	Token string
}

type SignInRequest struct {
	IpAddress    string `json:"-"`
	UserAgent    string `json:"-"`
	EmailAddress string `json:"emailAddress" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

type SignInResponse struct {
	Token string
}

type SignOutRequest struct {
	Token string
}

type MeResponse struct {
	User struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
	} `json:"user"`
}
